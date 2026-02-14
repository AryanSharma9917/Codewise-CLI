package deploy

import "fmt"

func Run(envName string, dryRun bool) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	////////////////////////////////////
	// PREFLIGHT
	////////////////////////////////////

	if err := Preflight(environment); err != nil {
		return err
	}

	////////////////////////////////////
	// ENSURE NAMESPACE
	////////////////////////////////////

	if err := EnsureNamespace(environment); err != nil {
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

	if err := executor.Run(command.Name, command.Args...); err != nil {
		return err
	}

	////////////////////////////////////
	// SKIP ROLLOUT FOR DRY RUN
	////////////////////////////////////

	if dryRun {
		return nil
	}

	////////////////////////////////////
	// MONITOR ROLLOUT
	////////////////////////////////////

	return MonitorRollout(environment)
}
