package cmd

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/docker"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related helpers",
}

var dockerInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		err := docker.InitDockerfile()
		if err != nil {
			fmt.Println("ℹ️", err.Error())
			return
		}
		fmt.Println("✅ Dockerfile created")
	},
}

func init() {
	dockerCmd.AddCommand(dockerInitCmd)
	rootCmd.AddCommand(dockerCmd)
}
