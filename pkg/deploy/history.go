package deploy

import (
	"fmt"
	"os/exec"
)

func History(envName string) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context
	release := environment.Helm.Release

	fmt.Println("Release History")
	fmt.Println("----------------")
	fmt.Println("Environment:", envName)
	fmt.Println("Namespace:", ns)
	fmt.Println("Release:", release)
	fmt.Println()

	args := []string{
		"history",
		release,
		"-n",
		ns,
	}

	if ctx != "" {
		args = append(args, "--kube-context", ctx)
	}

	cmd := exec.Command("helm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to fetch history")
	}

	fmt.Println(string(out))
	return nil
}
