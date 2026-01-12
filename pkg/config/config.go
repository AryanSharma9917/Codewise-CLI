package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	ConfigDirName  = ".codewise"
	ConfigFileName = "config.yaml"
)

type Config struct {
	Version string `yaml:"version"`
	User    struct {
		Name string `yaml:"name"`
	} `yaml:"user"`
	Defaults struct {
		AppName   string `yaml:"app_name"`
		Image     string `yaml:"image"`
		RepoURL   string `yaml:"repo_url"`
		Namespace string `yaml:"namespace"`
	} `yaml:"defaults"`
}

var DefaultConfig = []byte(`version: v1
user:
  name: aryan
defaults:
  app_name: myapp
  image: codewise:latest
  repo_url: https://github.com/example/repo
  namespace: default
`)

func InitConfig() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ConfigDirName)
	configPath := filepath.Join(configDir, ConfigFileName)

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	if _, err := os.Stat(configPath); err == nil {
		return configPath, fmt.Errorf("config already exists")
	}

	if err := os.WriteFile(configPath, DefaultConfig, 0644); err != nil {
		return "", err
	}

	return configPath, nil
}

func ReadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(home, ConfigDirName, ConfigFileName)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config file not found")
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
