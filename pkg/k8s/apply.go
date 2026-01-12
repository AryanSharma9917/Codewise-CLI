package k8s

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func ApplyManifests() error {
	path := filepath.Join("k8s", "app")

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("no manifests found at %s", path)
	}

	cmd := exec.Command("kubectl", "apply", "-f", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("running: kubectl apply -f", path)
	return cmd.Run()
}
