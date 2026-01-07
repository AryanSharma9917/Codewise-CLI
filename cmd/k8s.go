package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var (
	k8sAppName string
	k8sImage   string
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes helpers",
}

var k8sInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate Kubernetes manifests",
	Run: func(cmd *cobra.Command, args []string) {
		err := k8s.InitK8sManifests(k8s.Options{
			AppName: k8sAppName,
			Image:   k8sImage,
		})
		if err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println("✅ Kubernetes manifests created in ./k8s/app")
	},
}

func init() {
	k8sInitCmd.Flags().StringVar(&k8sAppName, "app", "", "Application name")
	k8sInitCmd.Flags().StringVar(&k8sImage, "image", "", "Docker image name")

	k8sCmd.AddCommand(k8sInitCmd)
	rootCmd.AddCommand(k8sCmd)
}
