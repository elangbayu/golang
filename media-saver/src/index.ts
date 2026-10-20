import { Container, getContainer } from "@cloudflare/containers";
import { env } from "cloudflare:workers";

// Runs the Go server (Dockerfile) on demand. It sleeps after 10 idle minutes,
// which is longer than the server's 5-minute per-download timeout, so a
// download that outlives the Shortcut's wait still finishes.
export class MediaSaver extends Container<Env> {
  defaultPort = 8080;
  sleepAfter = "10m";
  envVars = {
    API_TOKEN: env.API_TOKEN,
    RCLONE_CONF: env.RCLONE_CONF,
    RCLONE_REMOTE: env.RCLONE_REMOTE,
    COOKIES: env.COOKIES ?? "",
  };
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const { pathname } = new URL(request.url);
    if (request.method !== "POST" || pathname !== "/save") {
      return new Response("not found", { status: 404 });
    }
    // Check the token here so stray requests never wake (and bill) the container.
    if (!(await tokenMatches(request.headers.get("Authorization") ?? "", `Bearer ${env.API_TOKEN}`))) {
      return new Response("unauthorized", { status: 401 });
    }
    // A single instance: downloads are serialised inside the container anyway.
    return getContainer(env.MEDIA_SAVER, "default").fetch(request);
  },
};

async function tokenMatches(got: string, want: string): Promise<boolean> {
  const enc = new TextEncoder();
  const [a, b] = await Promise.all([
    crypto.subtle.digest("SHA-256", enc.encode(got)),
    crypto.subtle.digest("SHA-256", enc.encode(want)),
  ]);
  return crypto.subtle.timingSafeEqual(a, b);
}
