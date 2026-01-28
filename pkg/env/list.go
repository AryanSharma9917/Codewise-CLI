package env

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ListEnvs() ([]Env, error) {
	base, err := baseEnvPath()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	var envs []Env
	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()
			e, err := LoadEnv(name)
			if err != nil {
				continue
			}
			envs = append(envs, *e)
		}
	}

	return envs, nil
}

func LoadEnv(name string) (*Env, error) {
	dir, err := envDir(name)
	if err != nil {
		return nil, err
	}

	k8s := K8sConfig{}
	helm := HelmConfig{}
	gitops := GitOpsConfig{}
	values := ValuesConfig{}

	// read each file if exists
	_ = readYAML(filepath.Join(dir, "k8s.yaml"), &k8s)
	_ = readYAML(filepath.Join(dir, "helm.yaml"), &helm)
	_ = readYAML(filepath.Join(dir, "gitops.yaml"), &gitops)
	_ = readYAML(filepath.Join(dir, "values.yaml"), &values)

	return &Env{
		Name:   name,
		K8s:    k8s,
		Helm:   helm,
		GitOps: gitops,
		Values: values,
	}, nil
}

func readYAML(path string, out interface{}) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, out)
}
