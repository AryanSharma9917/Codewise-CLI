package deploy

import (
	"fmt"
	"strings"
)

func Run(envName string, dryRun bool) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	// A dry run must be usable without cluster access and must not create
	// resources such as namespaces.
	if !dryRun {
		if err := Preflight(environment); err != nil {
			return err
		}

		if err := EnsureNamespace(environment); err != nil {
			return err
		}
	}

	command, strategy, err := BuildCommand(environment)
	if err != nil {
		return err
	}

	executor := Executor{
		DryRun: dryRun,
	}

	fmt.Println("Starting deployment...")

	// Handle GitOps deployments specially
	if strategy == StrategyGitOps {
		app, err := BuildGitOpsApp(environment)
		if err != nil {
			return err
		}

		manifest, err := app.RenderManifest()
		if err != nil {
			return err
		}

		fmt.Println("GitOps deployment configuration:")
		fmt.Println("  Repository:", app.RepoURL)
		fmt.Println("  Path:", app.Path)
		fmt.Println("  Branch:", app.Branch)
		fmt.Println("  Target Namespace:", app.DestNS)
		fmt.Println()

		if dryRun {
			fmt.Println("[dry-run] Would deploy ArgoCD Application:")
			fmt.Println(manifest)
			return nil
		}

		executor.Stdin = strings.NewReader(manifest)
	}

	if err := executor.Run(command.Name, command.Args...); err != nil {
		return err
	}

	if dryRun {
		return nil
	}

	// MONITOR ROLLOUT
	if err := MonitorRollout(environment); err != nil {

		// fetch diagnostics on failure
		FetchDiagnostics(environment)

		return err
	}

	return nil
}
