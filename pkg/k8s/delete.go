package k8s

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func DeleteManifests(namespace, context string, dryRun bool) error {
	path := filepath.Join("k8s", "app")

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("no manifests found at %s", path)
	}

	namespace = resolveNamespace(namespace)
	context = resolveContext(context)

	args := []string{"delete", "-f", path}

	if namespace != "" {
		args = append(args, "-n", namespace)
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	if dryRun {
		fmt.Println("dry run:", "kubectl", args)
		return nil
	}

	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("running:", "kubectl", args)
	return cmd.Run()
}
