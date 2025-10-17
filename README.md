### Build the dockerfile

Bare bones demo for connecting to LinkedIn via Unipile:

- Checkpoints supported: `OTP`, `IN_APP_VALIDATION`.
- Supports passing `li_at` LinkedIn cookie value in order to login, where, due to a validated session already existing, we expect no checkpoint response.
- Updates `status` for account record based on the webhook events received from Unipile.
- The account and status can be then queried from the front page (root path).

Ngrok token is necessary, it's also necessary that on your public domain address you register the path to the webhook endpoint via Unipile dashboard.

```sh
docker build . -t linkedin-agent-app:latest
```

Make sure you (temporarily, for the purpose of this demo) export all your env vars
- unipile_dsn
- unipile_access_token
- NGROK_AUTHTOKEN
- DB_ROOT_PASSWORD
- db_user
- db_password
- db_name

```sh
docker compose up
```