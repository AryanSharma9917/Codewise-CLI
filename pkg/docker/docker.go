package docker

import (
	"fmt"
	"os"
)

const dockerfileName = "Dockerfile"

var defaultDockerfile = []byte(`
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
`)

// InitDockerfile creates a Dockerfile if it doesn't exist
func InitDockerfile() error {
	if _, err := os.Stat(dockerfileName); err == nil {
		return fmt.Errorf("Dockerfile already exists")
	}

	return os.WriteFile(dockerfileName, defaultDockerfile, 0644)
}
