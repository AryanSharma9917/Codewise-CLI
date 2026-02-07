package deploy

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executor struct {
	DryRun bool
}

func (e *Executor) Run(name string, args ...string) error {

	cmdStr := name + " " + strings.Join(args, " ")

	if e.DryRun {
		fmt.Println("[dry-run]", cmdStr)
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running:", cmdStr)

	return cmd.Run()
}
