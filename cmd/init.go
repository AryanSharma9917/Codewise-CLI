package cmd

import (
	"github.com/spf13/cobra"
	"github.com/aryansharma9917/codewise-cli/pkg/generator"
)

var (
	projectName    string
	withDocker     bool
	withDeployment bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold a new DevOps-ready project",
	Run: func(cmd *cobra.Command, args []string) {
		generator.ScaffoldProject(projectName, withDocker, withDeployment)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	initCmd.Flags().BoolVar(&withDocker, "with-docker", false, "Include Dockerfile")
	initCmd.Flags().BoolVar(&withDeployment, "with-deployment", false, "Include Kubernetes deployment YAML")
}
