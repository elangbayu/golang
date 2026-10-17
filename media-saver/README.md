# media-saver

Share a post from X, TikTok or Instagram on iOS, pick **Save to Drive**, and the video/photos land in a Google Drive folder.

```
iOS Share Sheet ──► Shortcut ──POST link──► media-saver (yt-dlp / gallery-dl) ──rclone──► Google Drive
```

Files are stored as `<RCLONE_REMOTE>/instagram`, `/tiktok` and `/x`. The default remote is `gdrive:Social Saves`.

## 1. Connect Google Drive (rclone)

On any computer with [rclone](https://rclone.org/install/) installed:

```sh
rclone config            # n → name: gdrive → storage: drive → scope: 1 (full) → defaults → log in via browser
rclone lsd gdrive:       # sanity check
mkdir -p config && cp ~/.config/rclone/rclone.conf config/rclone.conf
```

## 2. (Recommended) Cookies for Instagram

Instagram blocks most anonymous downloads. TikTok and X sometimes do the same. Export cookies from a logged-in browser in Netscape format, for example with the "Get cookies.txt LOCALLY" extension. Save them as `config/cookies.txt`. Use a secondary account if you can, because heavy automated use can get an account flagged.

## 3. Run the server

```sh
cp .env.example .env     # set API_TOKEN to a long random string (e.g. `openssl rand -hex 24`)
docker compose up -d --build
curl localhost:8080/healthz
```

**Where to host it:** the best option is a machine at home (a Raspberry Pi, NAS or always-on PC) plus [Tailscale](https://tailscale.com) on both the machine and the iPhone. That keeps it private and gives it a residential IP. Instagram and TikTok often block datacenter IPs, so a cheap VPS can work, but expect more failures. If you expose it publicly, put it behind HTTPS, for example with Caddy or a Cloudflare Tunnel.

Test it:

```sh
curl -X POST http://<host>:8080/save -H "Authorization: Bearer $API_TOKEN" \
     -H "Content-Type: application/json" -d '{"url":"https://x.com/user/status/123"}'
```

## 4. iOS Shortcut

Shortcuts app → **+** → name it **Save to Drive**:

1. Tap the **ⓘ** (Details) → enable **Show in Share Sheet**. Set *Share Sheet Types* to **URLs** and **Text**. Some apps share text that has the link in it.
2. Add the **Receive URLs and Text from Share Sheet** input (added automatically in step 1). Under "If there's no input", choose **Get Clipboard**. That lets you also run it from a copied link.
3. Add **Get Contents of URL**:
   - URL: `http://<host>:8080/save` (your Tailscale name/IP or public HTTPS URL)
   - Method: **POST**
   - Headers: `Authorization` = `Bearer <your API_TOKEN>`
   - Request Body: **JSON**, with key `url` (Text) = **Shortcut Input**
4. Add **Show Notification** with the body set to **Contents of URL**.

Usage: in TikTok, Instagram or X tap **Share**. Then pick **Save to Drive**. In TikTok and Instagram it may sit under **More…** or in the row of app icons. If you don't see it, use **Copy link** and run the shortcut from the widget or Back Tap. It reads the clipboard.

You'll get "Saved N file(s) from tiktok". Long downloads return "Still downloading…" after about 50 s and still finish in the background.

## How it works

- `POST /save` takes JSON `{"url": ...}`, a form field `url`, or plain text. The first link in the body is used, and tracking query params are removed.
- Instagram tries gallery-dl first, then yt-dlp. TikTok and X try yt-dlp first, then gallery-dl. gallery-dl covers photo posts, carousels and TikTok slideshows.
- Downloads go to a temp dir, then `rclone move` sends them to Drive. Files are named by post ID, so sharing the same post twice doesn't create duplicates.
- Only one download runs at a time.

## Maintenance

The sites change often, and yt-dlp and gallery-dl ship fixes quickly. When downloads start failing, rebuild to update them:

```sh
docker compose build --no-cache && docker compose up -d
docker compose logs -f   # shows OK/FAIL per link
```
