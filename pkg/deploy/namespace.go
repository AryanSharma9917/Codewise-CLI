package deploy

import (
	"fmt"
	"os/exec"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func namespaceExists(namespace string, context string) bool {

	args := []string{"get", "ns", namespace}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func createNamespace(namespace string, context string) error {

	args := []string{"create", "ns", namespace}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create namespace: %s", string(output))
	}

	return nil
}

func EnsureNamespace(environment *env.Env) error {

	ns := environment.K8s.Namespace
	ctx := environment.K8s.Context

	fmt.Printf("Checking namespace \"%s\"...\n", ns)

	if namespaceExists(ns, ctx) {
		fmt.Println("Namespace exists")
		return nil
	}

	fmt.Println("Namespace not found. Creating namespace...")

	if err := createNamespace(ns, ctx); err != nil {
		return err
	}

	fmt.Println("Namespace created")
	return nil
}
