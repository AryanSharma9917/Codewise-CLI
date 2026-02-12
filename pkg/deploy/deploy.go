package deploy

import "fmt"

func Run(envName string, dryRun bool) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	////////////////////////////////////////////////////
	// PREFLIGHT FIRST
	////////////////////////////////////////////////////

	if err := Preflight(environment); err != nil {
		return err
	}

	command, _, err := BuildCommand(environment)
	if err != nil {
		return err
	}

	executor := Executor{
		DryRun: dryRun,
	}

	fmt.Println("Starting deployment...")

	return executor.Run(command.Name, command.Args...)
}
