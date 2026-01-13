package cmd

import "github.com/spf13/cobra"

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes helpers",
}

func init() {
	rootCmd.AddCommand(k8sCmd)
}
