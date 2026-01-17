package helm

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitChart(appName, image string) error {
	base := filepath.Join("helm", "chart")

	// Ensure base path
	if _, err := os.Stat(base); err == nil {
		return fmt.Errorf("helm chart already exists at %s", base)
	}

	if err := os.MkdirAll(filepath.Join(base, "templates"), 0755); err != nil {
		return err
	}

	// Chart.yaml
	chart := []byte(fmt.Sprintf(`apiVersion: v2
name: %s
description: A Helm chart for deploying %s
version: 0.1.0
appVersion: "1.0.0"
`, appName, appName))

	if err := os.WriteFile(filepath.Join(base, "Chart.yaml"), chart, 0644); err != nil {
		return err
	}

	// values.yaml
	values := []byte(fmt.Sprintf(`image:
  repository: %s
  tag: latest
  pullPolicy: IfNotPresent

namespace: default
`, image))

	if err := os.WriteFile(filepath.Join(base, "values.yaml"), values, 0644); err != nil {
		return err
	}

	// templates
	deployment := []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Release.Name }}
  namespace: {{ .Values.namespace }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app: {{ .Release.Name }}
    spec:
      containers:
        - name: app
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
`)

	service := []byte(`apiVersion: v1
kind: Service
metadata:
  name: {{ .Release.Name }}
  namespace: {{ .Values.namespace }}
spec:
  ports:
    - port: 80
      targetPort: 80
  selector:
    app: {{ .Release.Name }}
`)

	if err := os.WriteFile(filepath.Join(base, "templates", "deployment.yaml"), deployment, 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(base, "templates", "service.yaml"), service, 0644); err != nil {
		return err
	}

	return nil
}
