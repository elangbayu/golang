// Secrets aren't in wrangler.jsonc, so `wrangler types` doesn't see them.
declare namespace Cloudflare {
  interface Env {
    API_TOKEN: string;
    RCLONE_CONF: string;
    COOKIES?: string;
  }
}
interface Env extends Cloudflare.Env {}
