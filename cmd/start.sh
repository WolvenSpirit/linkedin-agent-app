#!/bin/sh

echo "Export your env vars if you have already not done so, also read the README.md"

docker build . -t linkedin-agent-app:latest
docker compose up -d