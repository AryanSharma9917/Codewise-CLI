# GitOps Deployment Strategy

## Overview

Codewise-CLI now supports GitOps-based deployments through integration with ArgoCD. This allows you to deploy applications by managing configuration in a Git repository, with automatic synchronization and reconciliation handled by ArgoCD.

## Configuration

To enable GitOps deployments, configure the `gitops` section in your environment file:

```yaml
name: staging
k8s:
  namespace: staging
  context: prod-cluster
gitops:
  repo: https://github.com/your-org/your-deployment-repo
  path: environments/staging
  branch: main
```

### GitOps Configuration Fields

| Field | Required | Description |
|-------|----------|-------------|
| `repo` | Yes | Git repository URL containing Kubernetes manifests |
| `path` | No | Path within the repo (default: `.`) |
| `branch` | No | Git branch to deploy from (default: `main`) |

## Usage

### View Deployment Plan

Before deploying, review what will be deployed:

```bash
codewise deploy plan --env staging
```

This shows the ArgoCD Application manifest that will be created:

```
Deployment Plan
---------------
Environment: staging
Strategy: gitops

GitOps Configuration:
  Repository: https://github.com/your-org/your-deployment-repo
  Path: environments/staging
  Branch: main
  Target Namespace: staging
  Sync Policy: Automated with auto-prune and self-heal

ArgoCD Application Manifest:
---
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: staging
  namespace: argocd
spec:
  destination:
    namespace: staging
    name: prod-cluster
  project: default
  source:
    repoURL: https://github.com/your-org/your-deployment-repo
    targetRevision: main
    path: environments/staging
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
    - CreateNamespace=true
---
```

### Deploy with GitOps

```bash
codewise deploy run --env staging
```

This will:
1. Validate cluster connectivity and tools
2. Create the target namespace if needed
3. Create an ArgoCD Application that points to your Git repository
4. Monitor the rollout and verify deployment success

### Dry-run Mode

Test deployment without making changes:

```bash
codewise deploy run --env staging --dry-run
```

## How It Works

### Strategy Resolution

Codewise-CLI uses the following resolution order to determine the deployment strategy:

1. **GitOps** - If `gitops.repo` is configured
2. **Helm** - If `helm/chart/` directory exists
3. **Kubectl** - Default, uses raw Kubernetes manifests in `k8s/`

### GitOps Deployment Flow

```
┌─────────────────────────────────────┐
│ Preflight Checks                    │
│ • helm availability                 │
│ • kubectl availability              │
│ • cluster connectivity              │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│ Ensure Namespace                    │
│ • Create namespace if needed        │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│ Build ArgoCD Application Manifest   │
│ • Template with environment config  │
│ • Set sync policies                 │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│ Apply Manifest with kubectl         │
│ • Create/update Application         │
│ • Stream via stdin                  │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│ Monitor Rollout                     │
│ • Watch Helm release status         │
│ • Verify pod readiness              │
│ • Timeout: 120 seconds              │
└─────────────────────────────────────┘
```

## ArgoCD Application Manifest

The generated manifest includes:

- **Destination**: Points to your cluster with target namespace
- **Source**: Git repository, path, and branch configuration
- **Sync Policy**: 
  - Automated syncing enabled
  - Auto-prune removes resources deleted from Git
  - Self-heal enabled for drift correction
  - Creates namespace if it doesn't exist

## Repository Structure

Your Git repository should contain manifests in the configured path:

```
your-deployment-repo/
├── environments/
│   ├── staging/
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   └── configmap.yaml
│   └── production/
│       ├── deployment.yaml
│       ├── service.yaml
│       └── configmap.yaml
├── base/
│   └── kustomization.yaml
└── README.md
```

Supported formats:
- Plain YAML manifests
- Kustomize overlays
- Helm charts
- ArgoCD plugins

## Monitoring and Troubleshooting

### Check Application Status

```bash
codewise deploy status --env staging
```

Shows:
- Helm release status
- Pod status and readiness
- Recent events

### View Logs

```bash
codewise deploy logs --env staging --follow
```

### View Deployment History

```bash
codewise deploy history --env staging
```

### Rollback to Previous Revision

```bash
codewise deploy rollback --env staging --revision 2
```

## Requirements

- ArgoCD must be installed on your Kubernetes cluster
- Service account with permissions to create Applications in `argocd` namespace
- Git repository with read access (public or configured credentials)
- `helm` and `kubectl` CLI tools

## Advantages of GitOps Strategy

1. **Declarative**: All configuration in Git, single source of truth
2. **Automated**: ArgoCD continuously reconciles desired vs actual state
3. **Auditable**: Full Git history of all changes
4. **Rollback-friendly**: Easy git revert for rollbacks
5. **Multi-cluster**: Same Git repo can deploy to multiple clusters
6. **Self-healing**: ArgoCD automatically corrects drift

## Examples

### Example 1: Staging Environment

**Environment file**: `environments/staging.yaml`
```yaml
name: staging
k8s:
  namespace: app-staging
  context: eks-staging
gitops:
  repo: https://github.com/myorg/deployment-configs
  path: environments/staging
  branch: main
```

Deploy:
```bash
codewise deploy run --env staging
```

### Example 2: Production with Specific Revision

**Environment file**: `environments/prod.yaml`
```yaml
name: production
k8s:
  namespace: app-prod
  context: eks-prod
gitops:
  repo: https://github.com/myorg/deployment-configs
  path: environments/prod
  branch: release-v1.2
```

Deploy:
```bash
codewise deploy run --env production
```

## Integration with CI/CD

Combine GitOps deployments with your CI/CD pipeline:

```yaml
# Example GitHub Actions workflow
deploy:
  runs-on: ubuntu-latest
  steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v2
    - run: go build -o bin/codewise-cli .
    - run: ./bin/codewise-cli deploy run --env staging
```

## Sync Policies

The generated Application uses these sync policies:

- **Automated Sync**: Changes to Git are automatically applied
- **Prune**: Resources removed from Git are deleted from cluster
- **Self-Heal**: Cluster state is automatically corrected to match Git
- **Create Namespace**: Target namespace is created if missing

To customize sync policies, edit the ArgoCD Application directly or modify `pkg/deploy/gitops.go`.
