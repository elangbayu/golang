# media-saver

Share a post from X, TikTok or Instagram on iOS, pick **Save to Drive**, and the video/photos land in a Google Drive folder.

```
iOS Share Sheet ─► Shortcut ─HTTPS─► Cloudflare Worker (checks token)
                                        └─► Cloudflare Container (wakes on demand)
                                              yt-dlp / gallery-dl ─► rclone ─► Google Drive
```

Files go to `Social Saves/instagram`, `/tiktok` and `/x`. Change the folder with `RCLONE_REMOTE` in `wrangler.jsonc`.

The container only runs (and bills) while it's handling a download, plus 10 idle minutes. Cost per download is a fraction of a cent on a `basic` instance. Uploading to Drive counts as container egress. See [Containers pricing](https://developers.cloudflare.com/containers/platform/pricing/). Requires the Workers Paid plan.

## 1. Connect Google Drive (rclone)

On your computer, with [rclone](https://rclone.org/install/) installed:

```sh
rclone config     # n → name: gdrive → storage: drive → scope: 1 → defaults → log in via browser
rclone lsd gdrive:
cat ~/.config/rclone/rclone.conf    # you'll paste this into a secret in step 3
```

If you use your own Google OAuth client ID (recommended for reliability), set the Google Cloud app's publishing status to **In production**. In "Testing" mode, refresh tokens expire after 7 days.

## 2. (Recommended) Instagram cookies

Instagram blocks most anonymous downloads, and Cloudflare's IPs are cloud addresses, so expect it to need a login. Export cookies for instagram.com only, in Netscape format, with a browser extension such as "Get cookies.txt LOCALLY". Use a secondary account. Exporting only instagram.com keeps the file small. Worker secrets have a size limit, so don't export every site.

## 3. Deploy

Needs Node.js and Docker running on the machine you deploy from. Wrangler builds the image locally and pushes it to Cloudflare.

```sh
npm install
npx wrangler login
npx wrangler secret put API_TOKEN          # paste a long random string, e.g. `openssl rand -hex 32`
npx wrangler secret put RCLONE_CONF < ~/.config/rclone/rclone.conf
npx wrangler secret put COOKIES < cookies.txt   # optional
npm run deploy
```

The deploy prints your URL, `https://media-saver.<your-subdomain>.workers.dev`. You can also attach a custom domain in the dashboard. The first container rollout can take a few minutes.

Test it:

```sh
curl -X POST https://media-saver.<sub>.workers.dev/save -H "Authorization: Bearer $API_TOKEN" \
     -H "Content-Type: application/json" -d '{"url":"https://x.com/user/status/123"}'
```

The first request after the container has been asleep takes a few extra seconds while it starts.

## 4. iOS Shortcut

Shortcuts app → **+** → name it **Save to Drive**:

1. Tap the **ⓘ** (Details) → enable **Show in Share Sheet**. Set *Share Sheet Types* to **URLs** and **Text**. Under "If there's no input", choose **Get Clipboard**.
2. Add **Get Contents of URL**:
   - URL: `https://media-saver.<sub>.workers.dev/save`
   - Method: **POST**
   - Headers: `Authorization` = `Bearer <your API_TOKEN>`
   - Request Body: **JSON**, with key `url` (Text) = **Shortcut Input**
3. Add **Show Notification** with the body set to **Contents of URL**.

Usage: in TikTok, Instagram or X tap **Share**, then pick **Save to Drive**. In TikTok and Instagram it may be under **More…**. If it isn't there, tap **Copy link** and run the shortcut from a widget or Back Tap. It reads the clipboard.

You'll see "Saved N file(s) from tiktok". Downloads longer than about 50 s reply "Still downloading…" and finish in the background.

## How it works

- `src/index.ts` is the Worker. It rejects requests without the right token, so random traffic never wakes the container. Valid requests go to a single container instance.
- `main.go` is the server inside the container. It writes `RCLONE_CONF` and `COOKIES` to files at startup. It takes the first link in the request and removes tracking params.
  - Instagram: gallery-dl first, then yt-dlp.
  - TikTok and X: yt-dlp first, then gallery-dl.
  - It downloads to a temp dir, then runs `rclone move` to Drive. Files are named by post ID, so re-sharing a post doesn't duplicate it.
- Logs: `npx wrangler tail`, or Workers → media-saver → Logs / Containers in the dashboard. Each link is logged as `OK` or `FAIL`.

## Maintenance

yt-dlp and gallery-dl need frequent updates as the sites change. When downloads start failing, run `npm run deploy` again. The image rebuild installs the latest versions. If the image is cached, run `docker builder prune` first.

## Local testing (optional)

`docker compose up --build` runs the same container on `localhost:8080`. It reads `./config/rclone.conf` and `./config/cookies.txt` and `API_TOKEN` from `.env`.
