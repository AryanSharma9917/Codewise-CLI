package deploy

import (
	"fmt"
	"os/exec"
)

func checkDependency(name string) error {

	_, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("%s not found in PATH. please install it to continue", name)
	}

	return nil
}

func Run(envName string, dryRun bool) error {

	environment, err := LoadEnvironment(envName)
	if err != nil {
		return err
	}

	strategy := ResolveStrategy(environment)

	executor := Executor{
		DryRun: dryRun,
	}

	switch strategy {

	case StrategyHelm:

		if err := checkDependency("helm"); err != nil {
			return err
		}

		args := []string{
			"upgrade",
			"--install",
			environment.Helm.Release,
			environment.Helm.Chart,
			"--namespace",
			environment.K8s.Namespace,
		}

		// inject kube-context if provided
		if environment.K8s.Context != "" {
			args = append(args, "--kube-context", environment.K8s.Context)
		}

		return executor.Run("helm", args...)

	case StrategyKubectl:

		if err := checkDependency("kubectl"); err != nil {
			return err
		}

		args := []string{
			"apply",
			"-f",
			"k8s/",
			"-n",
			environment.K8s.Namespace,
		}

		if environment.K8s.Context != "" {
			args = append(args, "--context", environment.K8s.Context)
		}

		return executor.Run("kubectl", args...)

	default:
		return fmt.Errorf("unknown deployment strategy")
	}
}
