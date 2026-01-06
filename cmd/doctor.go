package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/docker"
	"github.com/spf13/cobra"
)

var imageTag string

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related helpers",
}

var dockerInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.InitDockerfile(); err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println("✅ Dockerfile created")
	},
}

var dockerValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Dockerfile best practices",
	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.ValidateDockerfile(); err != nil {
			fmt.Println("❌", err.Error())
		}
	},
}

var dockerBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		if err := docker.BuildDockerImage(imageTag); err != nil {
			fmt.Println("❌ Docker build failed")
		}
	},
}

func init() {
	dockerBuildCmd.Flags().StringVarP(
		&imageTag,
		"tag",
		"t",
		"",
		"Docker image tag (default: codewise:latest)",
	)

	dockerCmd.AddCommand(dockerInitCmd)
	dockerCmd.AddCommand(dockerValidateCmd)
	dockerCmd.AddCommand(dockerBuildCmd)

	rootCmd.AddCommand(dockerCmd)
}
