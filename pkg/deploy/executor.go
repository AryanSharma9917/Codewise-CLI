package deploy

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Executor struct {
	DryRun bool
	Stdin  io.Reader
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

	if e.Stdin != nil {
		cmd.Stdin = e.Stdin
	}

	fmt.Println("Running:", cmdStr)

	return cmd.Run()
}
