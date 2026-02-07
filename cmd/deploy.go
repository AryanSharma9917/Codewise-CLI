package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/spf13/cobra"
)

var (
	deployEnv string
	dryRun    bool
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an application using the configured environment",
	Run: func(cmd *cobra.Command, args []string) {

		if deployEnv == "" {
			fmt.Println("please provide --env")
			return
		}

		if err := deploy.Run(deployEnv, dryRun); err != nil {
			fmt.Println("deploy error:", err)
		}
	},
}

func init() {

	deployCmd.Flags().StringVar(&deployEnv, "env", "", "Environment to deploy")
	deployCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview deployment")

	rootCmd.AddCommand(deployCmd)
}
