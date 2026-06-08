package env

import (
	"fmt"
	"os"
	"path/filepath"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"gopkg.in/yaml.v3"
)

type CreateOptions struct {
	Interactive bool
}

func CreateEnv(name string, opts CreateOptions) error {
	if opts.Interactive {
		cfg, _ := config.ReadConfig()

		defaultNs := inferOrDefault(cfg.Defaults.Namespace, name)
		defaultCtx := inferOrDefault(cfg.Defaults.Context, "")
		defaultRepo := inferOrDefault(cfg.Defaults.RepoURL, "")
		defaultBranch := inferOrDefault(cfg.Defaults.Branch, "main")
		defaultImage := inferOrDefault(cfg.Defaults.Image, "codewise")
		defaultTag := inferOrDefault(cfg.Defaults.ImageTag, "latest")

		answers := struct {
			Namespace string
			Context   string
			Repo      string
			Branch    string
			Image     string
			Tag       string
		}{}

		qs := []*survey.Question{
			{Name: "Namespace", Prompt: &survey.Input{Message: fmt.Sprintf("Namespace (default: %s)", defaultNs)}},
			{Name: "Context", Prompt: &survey.Input{Message: fmt.Sprintf("Kubernetes context (default: %s)", defaultCtx)}},
			{Name: "Repo", Prompt: &survey.Input{Message: fmt.Sprintf("GitOps repo (default: %s)", defaultRepo)}},
			{Name: "Branch", Prompt: &survey.Input{Message: fmt.Sprintf("GitOps branch (default: %s)", defaultBranch)}},
			{Name: "Image", Prompt: &survey.Input{Message: fmt.Sprintf("Image repository (default: %s)", defaultImage)}},
			{Name: "Tag", Prompt: &survey.Input{Message: fmt.Sprintf("Image tag (default: %s)", defaultTag)}},
		}

		if err := survey.Ask(qs, &answers); err != nil {
			return err
		}

		k8s := K8sConfig{
			Namespace: firstNonEmpty(answers.Namespace, defaultNs),
			Context:   firstNonEmpty(answers.Context, defaultCtx),
		}

		helm := HelmConfig{
			Release: name,
			Chart:   "./helm/chart",
			Values:  "./values.yaml",
		}

		gitops := GitOpsConfig{
			Repo:   firstNonEmpty(answers.Repo, defaultRepo),
			Path:   "",
			Branch: firstNonEmpty(answers.Branch, defaultBranch),
		}

		values := ValuesConfig{}
		values.Image.Repository = firstNonEmpty(answers.Image, defaultImage)
		values.Image.Tag = firstNonEmpty(answers.Tag, defaultTag)

		return CreateEnvFromParts(name, k8s, helm, gitops, values)
	}

	return createSilent(name)
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
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
