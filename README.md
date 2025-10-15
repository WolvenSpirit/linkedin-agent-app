### Build the dockerfile

```sh
docker build . -t linkedin-agent-app:latest
```

Make sure you (temporarily, for the purpose of this demo) export all your env vars
- unipile_dsn
- unipile_access_token
- NGROK_AUTHTOKEN

```sh
docker compose up
```