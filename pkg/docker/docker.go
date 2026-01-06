package docker

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const dockerfileName = "Dockerfile"

// InitDockerfile creates a Dockerfile if it doesn't exist
func InitDockerfile() error {
	if _, err := os.Stat(dockerfileName); err == nil {
		return fmt.Errorf("Dockerfile already exists")
	}

	defaultDockerfile := []byte(`
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
	return os.WriteFile(dockerfileName, defaultDockerfile, 0644)
}

// ValidateDockerfile inspects Dockerfile best practices
func ValidateDockerfile() error {
	file, err := os.Open(dockerfileName)
	if err != nil {
		return fmt.Errorf("Dockerfile not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hasMultiStage := false
	hasNonRoot := false
	baseImage := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "FROM") {
			if baseImage == "" {
				baseImage = line
			} else {
				hasMultiStage = true
			}
		}

		if strings.HasPrefix(line, "USER") {
			hasNonRoot = true
		}
	}

	fmt.Println("Dockerfile validation:")
	fmt.Println("----------------------")
	fmt.Println("Base image:", baseImage)

	if hasMultiStage {
		fmt.Println("✔ Multi-stage build detected")
	} else {
		fmt.Println("⚠ Single-stage build")
	}

	if hasNonRoot {
		fmt.Println("✔ Non-root user configured")
	} else {
		fmt.Println("⚠ Running as root user")
	}

	return nil
}

// BuildDockerImage runs docker build
func BuildDockerImage(tag string) error {
	if tag == "" {
		tag = "codewise:latest"
	}

	cmd := exec.Command("docker", "build", "-t", tag, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running:", strings.Join(cmd.Args, " "))
	return cmd.Run()
}
