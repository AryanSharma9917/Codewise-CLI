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
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if !yes {
			var confirm bool
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("Delete environment %q?", name),
			}
			if err := survey.AskOne(prompt, &confirm); err != nil {
				fmt.Println("error:", err)
				return
			}
			if !confirm {
				fmt.Println("aborted")
				return
			}
		}

		if err := env.DeleteEnv(name, env.DeleteOptions{Force: yes}); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("environment", name, "deleted")
	},
}

func init() {
	envDeleteCmd.Flags().BoolVar(&yes, "yes", false, "Skip confirmation")
	envCmd.AddCommand(envDeleteCmd)
}
