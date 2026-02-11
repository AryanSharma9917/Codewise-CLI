package env

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	codewiseHomeEnv = "CODEWISE_HOME"
)

/////////////////////////////////////////////////////////
// PATH HELPERS
/////////////////////////////////////////////////////////

func baseEnvPath() (string, error) {

	if home := os.Getenv(codewiseHomeEnv); home != "" {
		return filepath.Join(home, "envs"), nil
	}

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

/////////////////////////////////////////////////////////
// YAML READER
/////////////////////////////////////////////////////////

func readYAML(path string, out interface{}) error {

	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(raw, out)
}

/////////////////////////////////////////////////////////
// SINGLE SOURCE OF TRUTH
/////////////////////////////////////////////////////////

func LoadEnv(name string) (*Env, error) {

	if !envExists(name) {
		return nil, fmt.Errorf("environment %q does not exist", name)
	}

	dir, err := envDir(name)
	if err != nil {
		return nil, err
	}

	k8s := K8sConfig{}
	helm := HelmConfig{}
	gitops := GitOpsConfig{}
	values := ValuesConfig{}

	if err := readYAML(filepath.Join(dir, "k8s.yaml"), &k8s); err != nil {
		return nil, err
	}

	if err := readYAML(filepath.Join(dir, "helm.yaml"), &helm); err != nil {
		return nil, err
	}

	if err := readYAML(filepath.Join(dir, "gitops.yaml"), &gitops); err != nil {
		return nil, err
	}

	if err := readYAML(filepath.Join(dir, "values.yaml"), &values); err != nil {
		return nil, err
	}

	return &Env{
		Name:   name,
		K8s:    k8s,
		Helm:   helm,
		GitOps: gitops,
		Values: values,
	}, nil
}
