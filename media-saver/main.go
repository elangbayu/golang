// media-saver receives a link shared from iOS (via a Shortcut), downloads the
// media with yt-dlp / gallery-dl, and moves it to Google Drive with rclone.
package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type config struct {
	addr        string
	token       string
	remote      string // rclone destination, e.g. "gdrive:Social Saves"
	cookiesFile string // optional Netscape cookies.txt (needed for most Instagram content)
	workDir     string
	jobTimeout  time.Duration
}

// platform describes which downloaders to try, in order, for a site.
type platform struct {
	name  string
	hosts []string
	tools []string
}

var platforms = []platform{
	// gallery-dl handles posts, reels and carousels (photos + videos) on Instagram.
	{"instagram", []string{"instagram.com"}, []string{"gallery-dl", "yt-dlp"}},
	// yt-dlp for videos; gallery-dl as fallback for photo slideshows / image tweets.
	{"tiktok", []string{"tiktok.com"}, []string{"yt-dlp", "gallery-dl"}},
	{"x", []string{"x.com", "twitter.com", "fxtwitter.com", "vxtwitter.com"}, []string{"yt-dlp", "gallery-dl"}},
}

var urlRe = regexp.MustCompile(`https?://[^\s<>"]+`)

func main() {
	cfg := config{
		addr:        envOr("ADDR", ":8080"),
		token:       os.Getenv("API_TOKEN"),
		remote:      envOr("RCLONE_REMOTE", "gdrive:Social Saves"),
		cookiesFile: os.Getenv("COOKIES_FILE"),
		workDir:     envOr("WORK_DIR", os.TempDir()),
		jobTimeout:  5 * time.Minute,
	}
	if cfg.token == "" {
		log.Fatal("API_TOKEN must be set")
	}

	var mu sync.Mutex // one download at a time keeps a small box (or a Pi) happy
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "ok") })
	mux.HandleFunc("POST /save", func(w http.ResponseWriter, r *http.Request) {
		if !authorized(r, cfg.token) {
			reply(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		raw, err := readURL(r)
		if err != nil {
			reply(w, http.StatusBadRequest, err.Error())
			return
		}
		p, link, err := classify(raw)
		if err != nil {
			reply(w, http.StatusBadRequest, err.Error())
			return
		}

		// Detach from the request so the save still finishes if the Shortcut times out.
		done := make(chan result, 1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			ctx, cancel := context.WithTimeout(context.Background(), cfg.jobTimeout)
			defer cancel()
			n, err := save(ctx, cfg, p, link)
			if err != nil {
				log.Printf("FAIL %s %s: %v", p.name, link, err)
			} else {
				log.Printf("OK   %s %s: %d file(s)", p.name, link, n)
			}
			done <- result{n, err}
		}()

		select {
		case res := <-done:
			if res.err != nil {
				reply(w, http.StatusBadGateway, "Failed: "+res.err.Error())
				return
			}
			reply(w, http.StatusOK, fmt.Sprintf("Saved %d file(s) from %s", res.n, p.name))
		case <-time.After(50 * time.Second):
			reply(w, http.StatusAccepted, "Still downloading, will land in Drive shortly")
		}
	})

	log.Printf("listening on %s, saving to %q", cfg.addr, cfg.remote)
	log.Fatal(http.ListenAndServe(cfg.addr, mux))
}

type result struct {
	n   int
	err error
}

// save downloads link into a fresh temp dir and moves the files to Drive.
func save(ctx context.Context, cfg config, p platform, link string) (int, error) {
	dir, err := os.MkdirTemp(cfg.workDir, "dl-")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(dir)

	var errs []error
	for _, tool := range p.tools {
		if err := download(ctx, cfg, tool, link, dir); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", tool, err))
			continue
		}
		files, _ := mediaFiles(dir)
		if len(files) == 0 {
			errs = append(errs, fmt.Errorf("%s: no files downloaded", tool))
			continue
		}
		dest := strings.TrimRight(cfg.remote, "/") + "/" + p.name
		if out, err := exec.CommandContext(ctx, "rclone", "move", dir, dest).CombinedOutput(); err != nil {
			return 0, fmt.Errorf("rclone: %v: %s", err, tail(out))
		}
		return len(files), nil
	}
	return 0, errors.Join(errs...)
}

func download(ctx context.Context, cfg config, tool, link, dir string) error {
	var args []string
	switch tool {
	case "yt-dlp":
		args = []string{
			"--no-playlist", "--no-progress",
			"-f", "bv*[ext=mp4]+ba[ext=m4a]/b[ext=mp4]/bv*+ba/b",
			"--merge-output-format", "mp4",
			"-o", filepath.Join(dir, "%(uploader_id,uploader)s_%(id)s.%(ext)s"),
		}
	case "gallery-dl":
		args = []string{"-D", dir}
	}
	if cfg.cookiesFile != "" {
		args = append(args, "--cookies", cfg.cookiesFile)
	}
	args = append(args, link)

	out, err := exec.CommandContext(ctx, tool, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, tail(out))
	}
	return nil
}

func mediaFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() && !strings.HasSuffix(path, ".part") {
			files = append(files, path)
		}
		return err
	})
	return files, err
}

// readURL accepts JSON {"url": "..."}, a form field "url", or a raw text body.
// iOS often shares "Check this out! https://..." so the first URL is extracted.
func readURL(r *http.Request) (string, error) {
	r.Body = http.MaxBytesReader(nil, r.Body, 64<<10)
	var text string
	switch ct := r.Header.Get("Content-Type"); {
	case strings.HasPrefix(ct, "application/json"):
		var body struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return "", fmt.Errorf("bad json: %w", err)
		}
		text = body.URL
	case strings.HasPrefix(ct, "application/x-www-form-urlencoded"), strings.HasPrefix(ct, "multipart/form-data"):
		text = r.FormValue("url")
	default:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return "", err
		}
		text = string(b)
	}
	m := urlRe.FindString(text)
	if m == "" {
		return "", errors.New("no URL found in request")
	}
	return m, nil
}

func classify(raw string) (platform, string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return platform{}, "", err
	}
	host := strings.ToLower(u.Hostname())
	for _, p := range platforms {
		for _, h := range p.hosts {
			if host == h || strings.HasSuffix(host, "."+h) {
				if p.name == "x" && host != "x.com" && host != "twitter.com" {
					u.Host = "x.com" // normalise fx/vx mirrors
				}
				u.RawQuery = "" // strip share tracking (?igsh=, ?s=20, ...)
				return p, u.String(), nil
			}
		}
	}
	return platform{}, "", fmt.Errorf("unsupported site: %s", host)
}

func authorized(r *http.Request, token string) bool {
	got := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	return subtle.ConstantTimeCompare([]byte(got), []byte(token)) == 1
}

func reply(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

func tail(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > 400 {
		s = "..." + s[len(s)-400:]
	}
	return s
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
