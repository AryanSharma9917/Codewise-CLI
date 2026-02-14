package deploy

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func MonitorRollout(environment *env.Env) error {

	release := environment.Helm.Release
	namespace := environment.K8s.Namespace
	context := environment.K8s.Context

	fmt.Println("Waiting for deployment rollout...")

	// We assume Helm creates a Deployment with the release name.
	// This is standard Helm behavior unless overridden.

	args := []string{
		"rollout",
		"status",
		"deployment/" + release,
		"-n",
		namespace,
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)

	// Stream output live
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("deployment failed to roll out")
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Deployment successfully rolled out.")
	return nil
}
