package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var (
	k8sInitAppName string
	k8sInitImage   string
)

var k8sInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Kubernetes deployment and service manifests",
	Run: func(cmd *cobra.Command, args []string) {
		opts := k8s.Options{
			AppName: k8sInitAppName,
			Image:   k8sInitImage,
		}

		if err := k8s.InitK8sManifests(opts); err != nil {
			fmt.Println("info:", err.Error())
			return
		}

		fmt.Println("k8s manifests created at k8s/app")
	},
}

func init() {
	k8sInitCmd.Flags().StringVar(&k8sInitAppName, "app", "", "Application name for manifests")
	k8sInitCmd.Flags().StringVar(&k8sInitImage, "image", "", "Container image for manifests")
	k8sCmd.AddCommand(k8sInitCmd)
}
