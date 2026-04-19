package cmd

import (
	"github.com/aryansharma9917/codewise-cli/pkg/k8s"
	"github.com/spf13/cobra"
)

var (
	k8sDeleteNamespace string
	k8sDeleteContext   string
	k8sDeleteDryRun    bool
)

var k8sDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Kubernetes resources from the current cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !k8sDeleteDryRun {
			if err := k8s.CheckKubectl(); err != nil {
				return LogError(err.Error())
			}
			if err := k8s.CheckCluster(); err != nil {
				return LogError(err.Error())
			}
		}
		return k8s.DeleteManifests(k8sDeleteNamespace, k8sDeleteContext, k8sDeleteDryRun)
	},
}

func init() {
	k8sDeleteCmd.Flags().StringVar(&k8sDeleteNamespace, "namespace", "", "Kubernetes namespace for deletion")
	k8sDeleteCmd.Flags().StringVar(&k8sDeleteContext, "context", "", "Kubernetes context for deletion")
	k8sDeleteCmd.Flags().BoolVar(&k8sDeleteDryRun, "dry-run", false, "Preview deletion without applying")
	k8sCmd.AddCommand(k8sDeleteCmd)
}
