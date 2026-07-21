package deploy

import (
	"fmt"
	"strings"
)

func Plan(envName string) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	command, strategy, err := BuildCommand(environment)
	if err != nil {
		return err
	}

	fmt.Println("\nDeployment Plan")
	fmt.Println("---------------")

	fmt.Println("Environment:", envName)
	fmt.Println("Strategy:", strategy)

	// Display GitOps specific information
	if strategy == StrategyGitOps {
		info, err := GetGitOpsDeploymentInfo(environment)
		if err != nil {
			return err
		}

		fmt.Println("\nGitOps Configuration:")
		fmt.Println("  Repository:", info["repo"])
		fmt.Println("  Path:", info["path"])
		fmt.Println("  Branch:", info["branch"])
		fmt.Println("  Target Namespace:", info["namespace"])
		fmt.Println("  Sync Policy:", info["syncPolicy"])

		fmt.Println("\nArgoCD Application Manifest:")
		fmt.Println("---")
		fmt.Println(info["manifest"])
		fmt.Println("---")
	} else {
		fmt.Println("\nCommand:")
		fmt.Printf("%s %s\n\n",
			command.Name,
			strings.Join(command.Args, " "),
		)
	}

	return nil
}
