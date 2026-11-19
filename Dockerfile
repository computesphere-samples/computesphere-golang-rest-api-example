# Go 1.26.11 carries the 2026-11 Go security release (Go 1.22 is no longer
# supported). Pinned by digest so a stale cached tag cannot bring back an
# unfixed build.
FROM golang:1.26.11@sha256:4165ee6bf3b7811fd7b40373290405c83d0167dcb87ee5f29c427cdcef45ff8f

ENV GIN_MODE=release

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build
RUN go build -o /computesphere-golang-rest-api-example

EXPOSE 8000
CMD [ "/computesphere-golang-rest-api-example" ]