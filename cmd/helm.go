package cmd

import "github.com/spf13/cobra"

var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Helm chart tooling",
}

func init() {
	rootCmd.AddCommand(helmCmd)
}
