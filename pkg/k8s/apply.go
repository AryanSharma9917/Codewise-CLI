package k8s

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aryansharma9917/codewise-cli/pkg/config"
)

func CheckKubectl() error {
	_, err := exec.LookPath("kubectl")
	if err != nil {
		return errors.New("kubectl not found in PATH")
	}
	return nil
}

func CheckCluster() error {
	cmd := exec.Command("kubectl", "version", "--short")
	if err := cmd.Run(); err != nil {
		return errors.New("cluster unreachable or misconfigured")
	}
	return nil
}

func resolveNamespace(flag string) string {
	if flag != "" {
		return flag
	}

	cfg, err := config.ReadConfig()
	if err == nil && cfg.Defaults.Namespace != "" {
		return cfg.Defaults.Namespace
	}

	return "default"
}

func resolveContext(flag string) string {
	if flag != "" {
		return flag
	}

	cfg, err := config.ReadConfig()
	if err == nil && cfg.Defaults.Context != "" {
		return cfg.Defaults.Context
	}

	return ""
}

func ApplyManifests(namespace, context string) error {
	path := filepath.Join("k8s", "app")

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("no manifests found at %s", path)
	}

	namespace = resolveNamespace(namespace)
	context = resolveContext(context)

	args := []string{"apply", "-f", path}

	if namespace != "" {
		args = append(args, "-n", namespace)
	}

	if context != "" {
		args = append(args, "--context", context)
	}

	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("running: kubectl", args)
	return cmd.Run()
}
