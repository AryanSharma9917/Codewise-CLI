package deploy

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func getPods(namespace string, context string) ([]string, error) {

	args := []string{
		"get",
		"pods",
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
		return nil, fmt.Errorf("failed to list pods")
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	var pods []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			pods = append(pods, line)
		}
	}

	return pods, nil
}

func describePod(pod string, namespace string, context string) {

	fmt.Printf("\nDiagnostics for %s:\n", pod)

	args := []string{
		"describe",
		pod,
		"-n",
		namespace,
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("failed to describe pod")
		return
	}

	fmt.Println(string(out))
}

func FetchDiagnostics(environment *env.Env) {

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context

	fmt.Println("\nFetching failure diagnostics...")

	pods, err := getPods(ns, ctx)
	if err != nil || len(pods) == 0 {
		fmt.Println("No pods found for diagnostics.")
		return
	}

	for _, pod := range pods {
		describePod(pod, ns, ctx)
	}
}
