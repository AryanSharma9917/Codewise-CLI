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
	Short: "Deployment operations",
}

var deployRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute deployment",
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

var deployPlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "Preview deployment execution plan",
	Run: func(cmd *cobra.Command, args []string) {

		if deployEnv == "" {
			fmt.Println("please provide --env")
			return
		}

		if err := deploy.Plan(deployEnv); err != nil {
			fmt.Println("plan error:", err)
		}
	},
}

func init() {

	deployRunCmd.Flags().StringVar(&deployEnv, "env", "", "Environment to deploy")
	deployRunCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview deployment")

	deployPlanCmd.Flags().StringVar(&deployEnv, "env", "", "Environment to plan")

	deployCmd.AddCommand(deployRunCmd)
	deployCmd.AddCommand(deployPlanCmd)

	rootCmd.AddCommand(deployCmd)
}
