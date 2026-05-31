package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/aryansharma9917/codewise-cli/pkg/env"
	"github.com/spf13/cobra"
)

var (
	previewPR    int
	previewImage string
	previewKeep  bool
)

var deployPreviewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Create a temporary preview environment for a PR",
	RunE: func(cmd *cobra.Command, args []string) error {
		if previewPR == 0 && previewImage == "" {
			return ExitError("please provide --pr or --image to create a preview")
		}

		name := fmt.Sprintf("preview-pr-%d-%d", previewPR, time.Now().Unix())

		// create minimal env parts
		k8s := env.K8sConfig{Namespace: name}
		helmCfg := env.HelmConfig{Release: name, Chart: "./helm/chart", Values: "./values.yaml"}
		gitops := env.GitOpsConfig{Repo: "", Path: "", Branch: "main"}
		values := env.ValuesConfig{}
		if previewImage != "" {
			// allow image specified as repo:tag or tag
			if strings.Contains(previewImage, ":") {
				parts := strings.SplitN(previewImage, ":", 2)
				values.Image.Repository = parts[0]
				values.Image.Tag = parts[1]
			} else {
				values.Image.Repository = "codewise"
				values.Image.Tag = previewImage
			}
		} else {
			values.Image.Repository = "codewise"
			values.Image.Tag = "latest"
		}

		if err := env.CreateEnvFromParts(name, k8s, helmCfg, gitops, values); err != nil {
			return LogErrorf("failed to create preview env: %v", err)
		}

		defer func() {
			if previewKeep {
				LogSuccess("preview environment kept: " + name)
				return
			}
			// best-effort cleanup
			if err := env.RemoveEnv(name); err != nil {
				fmt.Fprintln(os.Stderr, "warning: failed to remove preview env:", err)
			}
		}()

		fmt.Println("Created preview environment:", name)

		if err := deploy.Run(name, false); err != nil {
			return LogErrorf("preview deploy failed: %v", err)
		}

		LogSuccess("preview deployment complete: " + name)
		return nil
	},
}

func init() {
	deployPreviewCmd.Flags().IntVar(&previewPR, "pr", 0, "PR number to label the preview")
	deployPreviewCmd.Flags().StringVar(&previewImage, "image", "", "Image repository[:tag] to use for preview")
	deployPreviewCmd.Flags().BoolVar(&previewKeep, "keep", false, "Keep the preview environment after deployment")

	deployCmd.AddCommand(deployPreviewCmd)
}
