package cmd

import (
	"github.com/spf13/cobra"
	"github.com/aryansharma9917/Codewise-CLI/pkg/generator"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold a new DevOps-ready project",
	Run: func(cmd *cobra.Command, args []string) {
		generator.ScaffoldProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
