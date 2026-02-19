package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/spf13/cobra"
)

var statusEnv string

var deployStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show deployment status for an environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		if statusEnv == "" {
			return fmt.Errorf("please provide --env")
		}

		return deploy.Status(statusEnv)
	},
}

func init() {
	deployCmd.AddCommand(deployStatusCmd)
	deployStatusCmd.Flags().StringVar(&statusEnv, "env", "", "Environment name")
}
