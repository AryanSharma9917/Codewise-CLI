package deploy

import (
	"fmt"
	"os/exec"
)

func checkDependency(name string) error {

	_, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("%s not found in PATH. please install it to continue", name)
	}

	return nil
}

func Run(envName string, dryRun bool) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	command, _, err := BuildCommand(environment)
	if err != nil {
		return err
	}

	if err := checkDependency(command.Name); err != nil {
		return err
	}

	executor := Executor{
		DryRun: dryRun,
	}

	return executor.Run(command.Name, command.Args...)
}
