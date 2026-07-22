# GitOps with Argo CD

Codewise can generate and apply an Argo CD Application that points at a Git repository. Argo CD then reconciles the desired state.

## Requirements

- Argo CD installed in the cluster
- Permission to create Applications in the `argocd` namespace
- A repository containing Kubernetes manifests, a Kustomize overlay, or a Helm chart
- Repository credentials configured in Argo CD when the repository is private

## Configure an environment

Create an environment and update its `gitops.yaml`:

```yaml
repo: https://github.com/your-org/deployments
path: environments/staging
branch: main
```

Set the target namespace and context in `k8s.yaml`:

```yaml
namespace: staging
context: test-cluster
```

## Plan first

```bash
codewise deploy plan --env staging
```

GitOps takes priority over Helm and raw-manifest strategies when `repo` is configured.

## Safe preview

```bash
codewise deploy run --env staging --dry-run
```

This renders the Argo CD Application without contacting the cluster or creating a namespace.

## Deploy

```bash
kubectl config current-context
codewise deploy run --env staging
```

The generated Application enables automated synchronization, pruning, self-healing, and namespace creation.

::: warning Rollback behavior
For GitOps deployments, prefer reverting the desired-state commit in Git. Helm revision rollback commands are intended for Helm-managed releases.
:::
