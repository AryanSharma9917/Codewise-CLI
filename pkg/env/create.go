package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"gopkg.in/yaml.v3"
)

type CreateOptions struct {
	Interactive bool
}

func CreateEnv(name string, opts CreateOptions) error {
	if opts.Interactive {
		// handled at CLI layer
		return fmt.Errorf("interactive mode not implemented in CreateEnv")
	}
	return createSilent(name)
}

func createSilent(name string) error {
	if err := ensureBaseDir(); err != nil {
		return err
	}

	if envExists(name) {
		return fmt.Errorf("environment %q already exists", name)
	}

	dir, err := envDir(name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// global config support
	cfg, _ := config.ReadConfig()

	k8s := K8sConfig{
		Namespace: inferOrDefault(cfg.Defaults.Namespace, name),
		Context:   inferOrDefault(cfg.Defaults.Context, ""),
	}

	helm := HelmConfig{
		Release: name,
		Chart:   "./helm/chart",
		Values:  "./values.yaml",
	}

	gitops := GitOpsConfig{
		Repo:   "",
		Path:   "",
		Branch: "main",
	}

	values := ValuesConfig{}
	values.Image.Repository = inferOrDefault(cfg.Defaults.Image, "codewise")
	values.Image.Tag = inferOrDefault(cfg.Defaults.ImageTag, "latest")

	if err := writeYAML(filepath.Join(dir, "k8s.yaml"), k8s); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "helm.yaml"), helm); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "gitops.yaml"), gitops); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "values.yaml"), values); err != nil {
		return err
	}

	return nil
}

// used by interactive path
func CreateEnvFromParts(name string, k8s K8sConfig, helm HelmConfig, gitops GitOpsConfig, values ValuesConfig) error {
	if err := ensureBaseDir(); err != nil {
		return err
	}

	if envExists(name) {
		return fmt.Errorf("environment %q already exists", name)
	}

	dir, err := envDir(name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := writeYAML(filepath.Join(dir, "k8s.yaml"), k8s); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "helm.yaml"), helm); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "gitops.yaml"), gitops); err != nil {
		return err
	}
	if err := writeYAML(filepath.Join(dir, "values.yaml"), values); err != nil {
		return err
	}

	return nil
}

func writeYAML(path string, data interface{}) error {
	out, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}

func inferOrDefault(cfgVal, fallback string) string {
	if cfgVal != "" {
		return cfgVal
	}
	return fallback
}
