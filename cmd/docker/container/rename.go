package container

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short : "Rename a container"
	Run : func(cmd *cobra.Command, args []string) {
		renameContainer()
	},
}

func renameContainer() {
	ctx, cli := dockerClient()
	containerList := runningContaienrList(cli, clx)

	if len(containerList) == 0 {
		fmt.Println("No container running")
		return
	}
	for _, container := 
}