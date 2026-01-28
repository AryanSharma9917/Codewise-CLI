package cmd

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"github.com/aryansharma9917/codewise-cli/pkg/env"
	"github.com/spf13/cobra"
)

var interactive bool

var envCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new environment",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if interactive {
			if err := createEnvInteractive(name); err != nil {
				fmt.Println("error:", err)
				return
			}
			fmt.Println("environment", name, "created")
			return
		}

		if err := env.CreateEnv(name, env.CreateOptions{}); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("environment", name, "created")
	},
}

func init() {
	envCreateCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Enable interactive mode")
	envCmd.AddCommand(envCreateCmd)
}

func createEnvInteractive(name string) error {
	cfg, _ := config.ReadConfig()

	defaultNs := fallback(cfg.Defaults.Namespace, name)
	defaultCtx := fallback(cfg.Defaults.Context, "")
	defaultRepo := fallback(cfg.Defaults.RepoURL, "")
	defaultBranch := fallback(cfg.Defaults.Branch, "main")
	defaultImage := fallback(cfg.Defaults.Image, "codewise")
	defaultTag := fallback(cfg.Defaults.ImageTag, "latest")

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

	k8s := env.K8sConfig{
		Namespace: fallback(answers.Namespace, defaultNs),
		Context:   fallback(answers.Context, defaultCtx),
	}

	helm := env.HelmConfig{
		Release: name,
		Chart:   "./helm/chart",
		Values:  "./values.yaml",
	}

	gitops := env.GitOpsConfig{
		Repo:   fallback(answers.Repo, defaultRepo),
		Path:   "",
		Branch: fallback(answers.Branch, defaultBranch),
	}

	values := env.ValuesConfig{}
	values.Image.Repository = fallback(answers.Image, defaultImage)
	values.Image.Tag = fallback(answers.Tag, defaultTag)

	return env.CreateEnvFromParts(name, k8s, helm, gitops, values)
}

func fallback(input, def string) string {
	if input != "" {
		return input
	}
	return def
}
