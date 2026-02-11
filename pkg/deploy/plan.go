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

	fmt.Println("\nCommand:")
	fmt.Printf("%s %s\n\n",
		command.Name,
		strings.Join(command.Args, " "),
	)

	return nil
}
