package container

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	runningContainer = []string{}
)

// containerCmd contains the container command
var ContainerCmd = &cobra.Command{
	Use:   "container [command] [flag]",
	Short: "container related commands",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}
