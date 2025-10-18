package cmd

import (
	"github.com/spf13/cobra"
	"github.com/aryansharma9917/codewise-cli/pkg/generator"
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
	Run: func(cmd *cobra.Command, args []string) {
		data := generator.TemplateData{
			AppName: appName,
			Repo:    repoURL,
		}
		generator.RenderTemplate("github-action", outputPath, data)
	},
}

var argoAppCmd = &cobra.Command{
	Use:   "argo-app",
	Short: "Generate an ArgoCD application manifest",
	Run: func(cmd *cobra.Command, args []string) {
		data := generator.TemplateData{
			AppName: appName,
			Repo:    repoURL,
		}
		generator.RenderTemplate("argo-app", outputPath, data)
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
		cmd.MarkFlagRequired("output")
	}
}
