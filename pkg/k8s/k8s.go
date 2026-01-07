package k8s

import (
	"fmt"
	"os"
	"path/filepath"
)

const k8sDir = "k8s/app"

type Options struct {
	AppName string
	Image   string
}

func deploymentYAML(opts Options) []byte {
	if opts.AppName == "" {
		opts.AppName = "codewise-app"
	}
	if opts.Image == "" {
		opts.Image = "codewise:latest"
	}

	return []byte(fmt.Sprintf(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
        - name: %s
          image: %s
          ports:
            - containerPort: 8080
`, opts.AppName, opts.AppName, opts.AppName, opts.AppName, opts.Image))
}

func serviceYAML(appName string) []byte {
	if appName == "" {
		appName = "codewise-app"
	}

	return []byte(fmt.Sprintf(`apiVersion: v1
kind: Service
metadata:
  name: %s-service
spec:
  type: ClusterIP
  selector:
    app: %s
  ports:
    - port: 80
      targetPort: 8080
`, appName, appName))
}

// InitK8sManifests creates Kubernetes manifests with optional values
func InitK8sManifests(opts Options) error {
	if _, err := os.Stat(k8sDir); err == nil {
		return fmt.Errorf("k8s/app directory already exists")
	}

	if err := os.MkdirAll(k8sDir, 0755); err != nil {
		return err
	}

	deployPath := filepath.Join(k8sDir, "deployment.yaml")
	svcPath := filepath.Join(k8sDir, "service.yaml")

	if err := os.WriteFile(deployPath, deploymentYAML(opts), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(svcPath, serviceYAML(opts.AppName), 0644); err != nil {
		return err
	}

	return nil
}
