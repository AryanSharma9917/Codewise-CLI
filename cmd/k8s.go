package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes helpers",
}

var k8sInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate Kubernetes manifests",
	Run: func(cmd *cobra.Command, args []string) {
		if err := k8s.InitK8sManifests(); err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println("✅ Kubernetes manifests created in ./k8s")
	},
}

func init() {
	k8sCmd.AddCommand(k8sInitCmd)
	rootCmd.AddCommand(k8sCmd)
}
