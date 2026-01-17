package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var (
	k8sNamespace string
	k8sContext   string
	k8sDryRun    bool
)

var k8sApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply Kubernetes manifests to the current cluster",
	Run: func(cmd *cobra.Command, args []string) {
		// If not a dry run, check cluster connectivity
		if !k8sDryRun {
			if err := k8s.CheckKubectl(); err != nil {
				fmt.Println("info:", err.Error())
				return
			}
			if err := k8s.CheckCluster(); err != nil {
				fmt.Println("info:", err.Error())
				return
			}
		}

		if err := k8s.ApplyManifests(k8sNamespace, k8sContext, k8sDryRun); err != nil {
			fmt.Println("info:", err.Error())
			return
		}
	},
}

func init() {
	k8sApplyCmd.Flags().StringVar(&k8sNamespace, "namespace", "", "Kubernetes namespace for deployment")
	k8sApplyCmd.Flags().StringVar(&k8sContext, "context", "", "Kubernetes context for deployment")
	k8sApplyCmd.Flags().BoolVar(&k8sDryRun, "dry-run", false, "Preview changes without applying them")
	k8sCmd.AddCommand(k8sApplyCmd)
}
