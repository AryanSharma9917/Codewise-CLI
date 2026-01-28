package env

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	codewiseHomeEnv = "CODEWISE_HOME"
)

func baseEnvPath() (string, error) {
	// Check override via env var
	if home := os.Getenv(codewiseHomeEnv); home != "" {
		return filepath.Join(home, "envs"), nil
	}

	// Fallback to ~/.codewise/envs
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to resolve user home: %w", err)
	}

	return filepath.Join(home, ".codewise", "envs"), nil
}

func envDir(name string) (string, error) {
	base, err := baseEnvPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, name), nil
}

func envExists(name string) bool {
	dir, err := envDir(name)
	if err != nil {
		return false
	}
	info, err := os.Stat(dir)
	return err == nil && info.IsDir()
}

func ensureBaseDir() error {
	base, err := baseEnvPath()
	if err != nil {
		return err
	}
	return os.MkdirAll(base, 0755)
}
