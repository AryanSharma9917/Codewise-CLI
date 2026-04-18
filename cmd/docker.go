package cmd

import (

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
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := docker.InitDockerfile(); err != nil {
			return LogError(err.Error())
		}
		LogSuccess("Dockerfile created")
		return nil
	},
}

var dockerValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Dockerfile best practices",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := docker.ValidateDockerfile(); err != nil {
			return LogError(err.Error())
		}
		return nil
	},
}

var dockerBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Docker image",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := docker.BuildDockerImage(imageTag); err != nil {
			return LogError("Docker build failed: %v", err)
		}
		LogSuccess("Docker image built successfully")
		return nil
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
