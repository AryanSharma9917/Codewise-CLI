package deploy

import (
	"fmt"
	"os/exec"
)

func Rollback(envName string, revision int) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context
	release := environment.Helm.Release

	fmt.Println("Starting rollback...")
	fmt.Println("Environment:", envName)
	fmt.Println("Release:", release)
	fmt.Println("Revision:", revision)
	fmt.Println()

	args := []string{
		"rollback",
		release,
		fmt.Sprintf("%d", revision),
		"-n",
		ns,
	}

	if ctx != "" {
		args = append(args, "--kube-context", ctx)
	}

	cmd := exec.Command("helm", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("rollback failed")
	}

	fmt.Println("Rollback executed successfully.")
	fmt.Println("Verifying rollout...")

	return MonitorRollout(environment)
}
