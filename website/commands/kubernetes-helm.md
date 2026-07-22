# Kubernetes and Helm commands

## `k8s init`

```bash
codewise k8s init --app myapp --image myapp:latest
```

| Flag | Description |
| --- | --- |
| `--app` | Application name used in manifests |
| `--image` | Container image used by the Deployment |

Files are generated under `k8s/app/`.

## `k8s apply`

```bash
codewise k8s apply --namespace dev --context minikube
codewise k8s apply --namespace dev --dry-run
```

| Flag | Description |
| --- | --- |
| `--namespace` | Target namespace |
| `--context` | Kubernetes context |
| `--dry-run` | Print the operation without applying resources |

## `k8s delete`

```bash
codewise k8s delete --namespace dev --context minikube
codewise k8s delete --namespace dev --dry-run
```

The flags match `k8s apply`.

## `helm init`

```bash
codewise helm init --app myapp --image myapp:latest
```

The generated chart is stored under `helm/chart/`.
