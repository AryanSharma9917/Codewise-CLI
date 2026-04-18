package cmd

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/aryansharma9917/codewise-cli/pkg/env"
	"github.com/spf13/cobra"
)

var yes bool

var envDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if !yes {
			var confirm bool
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("Delete environment %q?", name),
			}
			if err := survey.AskOne(prompt, &confirm); err != nil {
				return LogError("error: %v", err)
			}
			if !confirm {
				return LogError("aborted")
			}
		}

		if err := env.DeleteEnv(name, env.DeleteOptions{Force: yes}); err != nil {
			return LogError("error: %v", err)
		}

		LogSuccess("environment %s deleted", name)
		return nil
	},
}

func init() {
	envDeleteCmd.Flags().BoolVar(&yes, "yes", false, "Skip confirmation")
	envCmd.AddCommand(envDeleteCmd)
}
