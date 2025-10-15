# Build Stage
FROM golang:1.25.2 AS build
COPY . /app
WORKDIR /app
RUN go env -w GOFLAGS=-buildvcs=false
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build

# Runtime Stage
FROM alpine:latest
COPY --from=build /app /app
WORKDIR /app
ENTRYPOINT ["/app/linkedin-agent-app"]