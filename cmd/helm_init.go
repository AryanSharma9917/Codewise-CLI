package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/config"
	"github.com/aryansharma9917/codewise-cli/pkg/helm"
	"github.com/spf13/cobra"
)

var (
	helmAppName string
	helmImage   string
)

var helmInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Helm chart from Codewise configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadConfig()
		if err != nil {
			fmt.Println("info: config not found, run `codewise config init` first")
			return
		}

		app := helmAppName
		if app == "" {
			app = cfg.Defaults.AppName
		}

		image := helmImage
		if image == "" {
			image = cfg.Defaults.Image
		}

		if err := helm.InitChart(app, image); err != nil {
			fmt.Println("info:", err.Error())
			return
		}

		fmt.Printf("helm chart created at helm/chart\n")
	},
}

func init() {
	helmInitCmd.Flags().StringVar(&helmAppName, "app", "", "Application name for chart")
	helmInitCmd.Flags().StringVar(&helmImage, "image", "", "Container image for chart")
	helmCmd.AddCommand(helmInitCmd)
}
