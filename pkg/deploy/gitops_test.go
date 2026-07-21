package deploy

import (
	"strings"
	"testing"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func TestBuildGitOpsApp(t *testing.T) {
	tests := []struct {
		name    string
		env     *env.Env
		wantErr bool
		check   func(*GitOpsApp) bool
	}{
		{
			name: "valid gitops config",
			env: &env.Env{
				Name: "staging",
				K8s: env.K8sConfig{
					Namespace: "staging-ns",
					Context:   "prod-cluster",
				},
				GitOps: env.GitOpsConfig{
					Repo:   "https://github.com/org/repo",
					Path:   "manifests/staging",
					Branch: "main",
				},
			},
			wantErr: false,
			check: func(app *GitOpsApp) bool {
				return app.Name == "staging" &&
					app.RepoURL == "https://github.com/org/repo" &&
					app.Path == "manifests/staging" &&
					app.Branch == "main" &&
					app.DestNS == "staging-ns"
			},
		},
		{
			name: "missing repo",
			env: &env.Env{
				Name: "staging",
				K8s: env.K8sConfig{
					Namespace: "staging-ns",
				},
				GitOps: env.GitOpsConfig{},
			},
			wantErr: true,
		},
		{
			name: "default path and branch",
			env: &env.Env{
				Name: "prod",
				K8s: env.K8sConfig{
					Namespace: "prod-ns",
				},
				GitOps: env.GitOpsConfig{
					Repo: "https://github.com/org/repo",
				},
			},
			wantErr: false,
			check: func(app *GitOpsApp) bool {
				return app.Path == "." && app.Branch == "main"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := BuildGitOpsApp(tt.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildGitOpsApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil && !tt.check(app) {
				t.Errorf("BuildGitOpsApp() validation failed")
			}
		})
	}
}

func TestRenderManifest(t *testing.T) {
	app := &GitOpsApp{
		Name:      "my-app",
		Namespace: "argocd",
		RepoURL:   "https://github.com/org/repo",
		Path:      "k8s",
		Branch:    "main",
		DestName:  "in-cluster",
		DestNS:    "default",
		Project:   "default",
	}

	manifest, err := app.RenderManifest()
	if err != nil {
		t.Fatalf("RenderManifest() error = %v", err)
	}

	// Verify manifest contains required fields
	checks := []string{
		"apiVersion: argoproj.io/v1alpha1",
		"kind: Application",
		"name: my-app",
		"namespace: argocd",
		"repoURL: https://github.com/org/repo",
		"path: k8s",
		"targetRevision: main",
		"automated:",
		"prune: true",
		"selfHeal: true",
	}

	for _, check := range checks {
		if !strings.Contains(manifest, check) {
			t.Errorf("RenderManifest() missing '%s' in manifest", check)
		}
	}
}

func TestValidateGitOpsConfig(t *testing.T) {
	tests := []struct {
		name    string
		env     *env.Env
		wantErr bool
	}{
		{
			name: "valid config",
			env: &env.Env{
				K8s: env.K8sConfig{
					Namespace: "default",
				},
				GitOps: env.GitOpsConfig{
					Repo: "https://github.com/org/repo",
				},
			},
			wantErr: false,
		},
		{
			name: "missing repo",
			env: &env.Env{
				K8s: env.K8sConfig{
					Namespace: "default",
				},
			},
			wantErr: true,
		},
		{
			name: "missing namespace",
			env: &env.Env{
				GitOps: env.GitOpsConfig{
					Repo: "https://github.com/org/repo",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGitOpsConfig(tt.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGitOpsConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResolveStrategyGitOps(t *testing.T) {
	// Test that GitOps strategy is chosen when repo is configured
	env := &env.Env{
		GitOps: env.GitOpsConfig{
			Repo: "https://github.com/org/repo",
		},
	}

	strategy := ResolveStrategy(env)
	if strategy != StrategyGitOps {
		t.Errorf("ResolveStrategy() = %v, want %v", strategy, StrategyGitOps)
	}
}

func TestBuildCommandGitOps(t *testing.T) {
	env := &env.Env{
		K8s: env.K8sConfig{
			Context: "my-cluster",
		},
		GitOps: env.GitOpsConfig{
			Repo: "https://github.com/org/repo",
		},
	}

	cmd, strategy, err := BuildCommand(env)
	if err != nil {
		t.Fatalf("BuildCommand() error = %v", err)
	}

	if strategy != StrategyGitOps {
		t.Errorf("BuildCommand() strategy = %v, want %v", strategy, StrategyGitOps)
	}

	if cmd.Name != "kubectl" {
		t.Errorf("BuildCommand() name = %v, want kubectl", cmd.Name)
	}

	// Should have context flag
	hasContext := false
	for i, arg := range cmd.Args {
		if arg == "--context" && i+1 < len(cmd.Args) && cmd.Args[i+1] == "my-cluster" {
			hasContext = true
			break
		}
	}

	if !hasContext {
		t.Errorf("BuildCommand() missing --context flag")
	}
}
