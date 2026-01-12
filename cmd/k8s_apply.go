package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var k8sApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply Kubernetes manifests to the current cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if err := k8s.CheckKubectl(); err != nil {
			fmt.Println("info:", err.Error())
			return
		}

		if err := k8s.CheckCluster(); err != nil {
			fmt.Println("info:", err.Error())
			return
		}

		if err := k8s.ApplyManifests(); err != nil {
			fmt.Println("info:", err.Error())
			return
		}
	},
}

func init() {
	k8sCmd.AddCommand(k8sApplyCmd)
}
