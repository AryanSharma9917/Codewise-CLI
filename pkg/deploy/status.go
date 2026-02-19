package deploy

import (
	"fmt"
	"os/exec"
	"strings"
)

func Status(envName string) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context
	release := environment.Helm.Release

	fmt.Println("Deployment Status")
	fmt.Println("-----------------")
	fmt.Println("Environment:", envName)
	fmt.Println("Namespace:", ns)
	fmt.Println("Release:", release)
	fmt.Println()

	// Helm Status
	args := []string{
		"status",
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
		fmt.Println("Helm status not available")
	} else {
		fmt.Println(string(out))
	}

	// Pods
	fmt.Println("Pods:")
	podArgs := []string{
		"get",
		"pods",
		"-n",
		ns,
		"-o",
		"wide",
	}

	if ctx != "" {
		podArgs = append(podArgs, "--context", ctx)
	}

	podCmd := exec.Command("kubectl", podArgs...)
	podsOut, err := podCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Unable to fetch pods")
	} else {
		fmt.Println(strings.TrimSpace(string(podsOut)))
	}

	return nil
}
