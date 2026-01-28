package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
	"github.com/spf13/cobra"
)

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List environments",
	Run: func(cmd *cobra.Command, args []string) {
		envs, err := env.ListEnvs()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if len(envs) == 0 {
			fmt.Println("no environments found")
			return
		}

		for _, e := range envs {
			fmt.Printf("%-10s namespace=%s context=%s\n",
				e.Name, e.K8s.Namespace, e.K8s.Context)
		}
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
}
