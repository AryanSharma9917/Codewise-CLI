package deploy

import (
	"fmt"
	"os/exec"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func runCheck(name string, args ...string) error {

	cmd := exec.Command(name, args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Preflight(environment *env.Env) error {

	fmt.Println("Running preflight checks...")

	strategy := ResolveStrategy(environment)

	////////////////////////////////////////////////////
	// Check binary availability
	////////////////////////////////////////////////////

	switch strategy {

	case StrategyHelm:

		if err := runCheck("helm", "version"); err != nil {
			return fmt.Errorf("helm not available or not functioning")
		}

	case StrategyKubectl:

		if err := runCheck("kubectl", "version", "--client"); err != nil {
			return fmt.Errorf("kubectl not available or not functioning")
		}
	}

	////////////////////////////////////////////////////
	// Cluster connectivity check
	////////////////////////////////////////////////////

	args := []string{"cluster-info"}

	if environment.K8s.Context != "" {
		args = append(args, "--context", environment.K8s.Context)
	}

	if err := runCheck("kubectl", args...); err != nil {
		return fmt.Errorf("cannot reach kubernetes cluster. check kube-context")
	}

	fmt.Println("Cluster reachable")
	return nil
}
