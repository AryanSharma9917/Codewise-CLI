package docker

import (
	"github.com/aryansharma9917/Codewise-CLI/cmd/docker/container"
	"github.com/spf13/cobra"
)

// DockerCmd will be the root command for all docker related commands
var DockerCmd = &cobra.Command{
	Use:   "docker [command] [flags]",
	Short: "Docker related commands. Like generating a Dockerfile for a language.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		checkNilErr(err)
	},
}

func init() {
	DockerCmd.AddCommand(container.ContainerCmd)
	DockerCmd.AddCommand(dockerfileCmd)
}
