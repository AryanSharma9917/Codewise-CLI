package cmd

import "github.com/spf13/cobra"

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage deployment environments",
	Long:  "Create, list, and delete Codewise deployment environments.",
}

func init() {
	rootCmd.AddCommand(envCmd)
}
