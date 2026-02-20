package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/spf13/cobra"
)

var historyEnv string

var deployHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Show Helm release history for an environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		if historyEnv == "" {
			return fmt.Errorf("please provide --env")
		}

		return deploy.History(historyEnv)
	},
}

func init() {
	deployCmd.AddCommand(deployHistoryCmd)
	deployHistoryCmd.Flags().StringVar(&historyEnv, "env", "", "Environment name")
}
