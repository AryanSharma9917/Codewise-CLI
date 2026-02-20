package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/spf13/cobra"
)

var logsEnv string
var followLogs bool

var deployLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Stream logs for deployment pods",
	RunE: func(cmd *cobra.Command, args []string) error {

		if logsEnv == "" {
			return fmt.Errorf("please provide --env")
		}

		return deploy.Logs(logsEnv, followLogs)
	},
}

func init() {
	deployCmd.AddCommand(deployLogsCmd)
	deployLogsCmd.Flags().StringVar(&logsEnv, "env", "", "Environment name")
	deployLogsCmd.Flags().BoolVar(&followLogs, "follow", false, "Stream logs")
}
