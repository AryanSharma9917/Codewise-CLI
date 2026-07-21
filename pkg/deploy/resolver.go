package deploy

import (
	"os"
	"path/filepath"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

type Strategy string

const (
	StrategyHelm    Strategy = "helm"
	StrategyKubectl Strategy = "kubectl"
	StrategyGitOps  Strategy = "gitops"
)

func ResolveStrategy(e *env.Env) Strategy {

	// check if gitops repo is configured
	if e.GitOps.Repo != "" {
		return StrategyGitOps
	}

	// check if helm chart exists
	if _, err := os.Stat(filepath.Join(".", "helm", "chart")); err == nil {
		return StrategyHelm
	}

	return StrategyKubectl
}
