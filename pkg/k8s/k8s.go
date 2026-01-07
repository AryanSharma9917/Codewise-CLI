package k8s

import (
	"fmt"
	"os"
	"path/filepath"
)

const k8sDir = "k8s/app"

var deploymentYAML = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: codewise-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: codewise
  template:
    metadata:
      labels:
        app: codewise
    spec:
      containers:
        - name: codewise
          image: codewise:latest
          ports:
            - containerPort: 8080
`)

var serviceYAML = []byte(`apiVersion: v1
kind: Service
metadata:
  name: codewise-service
spec:
  type: ClusterIP
  selector:
    app: codewise
  ports:
    - port: 80
      targetPort: 8080
`)

// InitK8sManifests creates basic Kubernetes manifests
func InitK8sManifests() error {
	if _, err := os.Stat(k8sDir); err == nil {
		return fmt.Errorf("k8s directory already exists")
	}

	if err := os.MkdirAll(k8sDir, 0755); err != nil {
		return err
	}

	deployPath := filepath.Join(k8sDir, "deployment.yaml")
	svcPath := filepath.Join(k8sDir, "service.yaml")

	if err := os.WriteFile(deployPath, deploymentYAML, 0644); err != nil {
		return err
	}

	if err := os.WriteFile(svcPath, serviceYAML, 0644); err != nil {
		return err
	}

	return nil
}
