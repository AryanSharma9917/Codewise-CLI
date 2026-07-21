package deploy

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

// GitOpsApp represents an ArgoCD Application manifest
type GitOpsApp struct {
	Name      string
	Namespace string
	RepoURL   string
	Path      string
	Branch    string
	DestName  string
	DestNS    string
	Project   string
}

// ArgoCDTemplate is the template for ArgoCD Application manifest
const ArgoCDTemplate = `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: {{ .Name }}
  namespace: {{ .Namespace }}
spec:
  destination:
    namespace: {{ .DestNS }}
    name: {{ .DestName }}
  project: {{ .Project }}
  source:
    repoURL: {{ .RepoURL }}
    targetRevision: {{ .Branch }}
    path: {{ .Path }}
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
    - CreateNamespace=true
`

// BuildGitOpsApp creates a GitOpsApp from an environment configuration
func BuildGitOpsApp(environment *env.Env) (*GitOpsApp, error) {
	if environment.GitOps.Repo == "" {
		return nil, fmt.Errorf("gitops repo not configured in environment")
	}

	path := environment.GitOps.Path
	if path == "" {
		path = "."
	}

	branch := environment.GitOps.Branch
	if branch == "" {
		branch = "main"
	}

	destName := environment.K8s.Context
	if destName == "" {
		destName = "in-cluster"
	}

	return &GitOpsApp{
		Name:      environment.Name,
		Namespace: "argocd",
		RepoURL:   environment.GitOps.Repo,
		Path:      path,
		Branch:    branch,
		DestName:  destName,
		DestNS:    environment.K8s.Namespace,
		Project:   "default",
	}, nil
}

// RenderManifest generates the ArgoCD Application manifest YAML
func (ga *GitOpsApp) RenderManifest() (string, error) {
	tmpl, err := template.New("argocd-app").Parse(ArgoCDTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse ArgoCD template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ga); err != nil {
		return "", fmt.Errorf("failed to render ArgoCD manifest: %w", err)
	}

	return buf.String(), nil
}

// ValidateGitOpsConfig checks if GitOps configuration is valid
func ValidateGitOpsConfig(e *env.Env) error {
	if e.GitOps.Repo == "" {
		return fmt.Errorf("gitops.repo is required for GitOps deployments")
	}

	if e.K8s.Namespace == "" {
		return fmt.Errorf("k8s.namespace is required for GitOps deployments")
	}

	return nil
}

// GetGitOpsDeploymentInfo returns deployment information for GitOps strategy
func GetGitOpsDeploymentInfo(e *env.Env) (map[string]interface{}, error) {
	if err := ValidateGitOpsConfig(e); err != nil {
		return nil, err
	}

	app, err := BuildGitOpsApp(e)
	if err != nil {
		return nil, err
	}

	manifest, err := app.RenderManifest()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"strategy":   "gitops",
		"name":       app.Name,
		"repo":       app.RepoURL,
		"path":       app.Path,
		"branch":     app.Branch,
		"namespace":  app.DestNS,
		"manifest":   manifest,
		"syncPolicy": "Automated with auto-prune and self-heal",
	}, nil
}
