package cmd

import (
	"github.com/aryansharma9917/codewise-cli/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	outputPath string
	appName    string
	repoURL    string
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate templates like GitHub Actions or ArgoCD apps",
}

var githubActionCmd = &cobra.Command{
	Use:   "github-action",
	Short: "Generate a GitHub Actions CI workflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := generator.TemplateData{
			AppName: appName,
			Repo:    repoURL,
		}
		if err := generator.RenderTemplate("github-action", outputPath, data); err != nil {
			return ExitError(err.Error())
		}
		return nil
	},
}

var argoAppCmd = &cobra.Command{
	Use:   "argo-app",
	Short: "Generate an ArgoCD application manifest",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := generator.TemplateData{
			AppName: appName,
			Repo:    repoURL,
		}
		if err := generator.RenderTemplate("argo-app", outputPath, data); err != nil {
			return ExitError(err.Error())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	// Add subcommands to the main `template` command
	templateCmd.AddCommand(githubActionCmd)
	templateCmd.AddCommand(argoAppCmd)

	// Common flags for both subcommands
	for _, cmd := range []*cobra.Command{githubActionCmd, argoAppCmd} {
		cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path (required)")
		cmd.Flags().StringVar(&appName, "app-name", "myapp", "Application name for template")
		cmd.Flags().StringVar(&repoURL, "repo", "https://github.com/example/repo", "Repository URL for template")
		if err := cmd.MarkFlagRequired("output"); err != nil {
			panic(err)
		}
	}
}
