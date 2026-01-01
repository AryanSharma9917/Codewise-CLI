package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDirName  = ".codewise"
	ConfigFileName = "config.yaml"
)

// DefaultConfig is the initial config content
var DefaultConfig = []byte(`version: v1
user:
  name: aryan
defaults:
  app_name: myapp
  repo_url: https://github.com/example/repo
`)

// InitConfig creates the config directory and file
func InitConfig() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ConfigDirName)
	configPath := filepath.Join(configDir, ConfigFileName)

	// Create directory if not exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	// If config already exists, do nothing
	if _, err := os.Stat(configPath); err == nil {
		return configPath, fmt.Errorf("config already exists")
	}

	// Write default config
	if err := os.WriteFile(configPath, DefaultConfig, 0644); err != nil {
		return "", err
	}

	return configPath, nil
}
