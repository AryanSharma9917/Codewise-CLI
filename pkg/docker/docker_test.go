package docker

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withTempWorkingDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	oldwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir to temp dir: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Chdir(oldwd)
	})

	return dir
}

func TestInitDockerfileCreatesUpdatedGoBaseImage(t *testing.T) {
	withTempWorkingDir(t)

	if err := InitDockerfile(); err != nil {
		t.Fatalf("InitDockerfile() error = %v", err)
	}

	content, err := os.ReadFile(filepath.Join(".", dockerfileName))
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	text := string(content)
	if !strings.Contains(text, "FROM golang:1.25-alpine AS builder") {
		t.Fatalf("generated Dockerfile did not use Go 1.25: %s", text)
	}
}

func TestValidateDockerfileAcceptsGeneratedDockerfile(t *testing.T) {
	withTempWorkingDir(t)

	if err := InitDockerfile(); err != nil {
		t.Fatalf("InitDockerfile() error = %v", err)
	}

	if err := ValidateDockerfile(); err != nil {
		t.Fatalf("ValidateDockerfile() error = %v", err)
	}
}
