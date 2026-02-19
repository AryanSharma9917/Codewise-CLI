package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/deploy"
	"github.com/spf13/cobra"
)

var rollbackEnv string
var rollbackRevision int

var deployRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback a deployment to a previous Helm revision",
	RunE: func(cmd *cobra.Command, args []string) error {

		if rollbackEnv == "" {
			return fmt.Errorf("please provide --env")
		}

		if rollbackRevision <= 0 {
			return fmt.Errorf("please provide a valid --revision")
		}

		return deploy.Rollback(rollbackEnv, rollbackRevision)
	},
}

func init() {
	deployCmd.AddCommand(deployRollbackCmd)
	deployRollbackCmd.Flags().StringVar(&rollbackEnv, "env", "", "Environment name")
	deployRollbackCmd.Flags().IntVar(&rollbackRevision, "revision", 0, "Helm revision to rollback to")
}
