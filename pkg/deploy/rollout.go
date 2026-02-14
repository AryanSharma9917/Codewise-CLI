package deploy

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func getDeployments(namespace string, context string) ([]string, error) {

	args := []string{
		"get",
		"deployments",
		"-n",
		namespace,
		"-o",
		"name",
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments")
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	var deployments []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			deployments = append(deployments, line)
		}
	}

	return deployments, nil
}

func waitForDeployment(deployment string, namespace string, context string) error {

	args := []string{
		"rollout",
		"status",
		deployment,
		"-n",
		namespace,
		"--timeout=120s",
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}

func MonitorRollout(environment *env.Env) error {

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context

	fmt.Println("Waiting for deployment rollout...")

	deployments, err := getDeployments(ns, ctx)
	if err != nil {
		return err
	}

	if len(deployments) == 0 {
		fmt.Println("No deployments found. Skipping rollout monitoring.")
		return nil
	}

	for _, deploy := range deployments {
		fmt.Printf("Monitoring %s...\n", deploy)

		if err := waitForDeployment(deploy, ns, ctx); err != nil {
			return fmt.Errorf("deployment %s failed to roll out", deploy)
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("All deployments successfully rolled out.")
	return nil
}
