# How Codewise works

Codewise is a Go application built with Cobra. Commands parse user input, domain packages implement the behavior, and external tools perform platform operations.

```text
Terminal
  └─ codewise
      ├─ cmd/                 command definitions and flags
      └─ pkg/                 workflow logic
          ├─ local files     config, templates, manifests, charts
          └─ external CLIs   docker, kubectl, helm
                              └─ registry or Kubernetes API
```

## Repository map

| Path | Responsibility |
| --- | --- |
| `main.go` | Program entry point |
| `cmd/` | Cobra commands, flags, and input validation |
| `pkg/config/` | Global configuration |
| `pkg/docker/` | Dockerfile and image operations |
| `pkg/k8s/` | Manifest generation and Kubernetes execution |
| `pkg/helm/` | Helm chart generation |
| `pkg/deploy/` | Plans, strategies, execution, diagnostics, and rollback |
| `pkg/env/` | Environment profile lifecycle |
| `pkg/encoder/` | Data conversion utilities |
| `pkg/generator/` | Project and automation templates |

## Deployment strategy selection

Codewise resolves deployment strategies in this order:

1. GitOps when the environment has a Git repository configured.
2. Helm when `helm/chart/` exists.
3. Raw Kubernetes manifests otherwise.

Use `codewise deploy plan --env NAME` to see the selected strategy before execution.
