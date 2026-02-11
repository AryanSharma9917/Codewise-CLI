package deploy

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

type Command struct {
	Name string
	Args []string
}

func BuildCommand(environment *env.Env) (*Command, Strategy, error) {

	strategy := ResolveStrategy(environment)

	switch strategy {

	case StrategyHelm:

		args := []string{
			"upgrade",
			"--install",
			environment.Helm.Release,
			environment.Helm.Chart,
			"--namespace",
			environment.K8s.Namespace,
		}

		if environment.K8s.Context != "" {
			args = append(args, "--kube-context", environment.K8s.Context)
		}

		return &Command{
			Name: "helm",
			Args: args,
		}, strategy, nil

	case StrategyKubectl:

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

		return &Command{
			Name: "kubectl",
			Args: args,
		}, strategy, nil

	default:
		return nil, "", fmt.Errorf("unknown deployment strategy")
	}
}
