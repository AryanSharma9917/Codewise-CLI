# Codewise CLI Project Overview

## What This Project Does

Codewise CLI is a DevOps-focused command-line tool that unifies common workflows used in application delivery.

It provides one command surface for:
- Dockerfile scaffolding and image build operations
- Kubernetes manifest generation and apply/delete operations
- Helm chart scaffolding
- Deployment status/log/history/rollback flows
- Environment profile management
- Data transformation utilities (YAML, JSON, TOML, XML, ENV, Base64)
- Template generation (for example GitHub Actions and Argo app templates)

The goal is to reduce command sprawl and boilerplate when moving from code to deployable workloads.

## How It Is Built

The project is implemented in Go and organized as a Cobra-based CLI.

High-level build path:
- `main.go` is the executable entry point.
- `cmd/` contains CLI command definitions and flag wiring.
- `pkg/` contains implementation logic used by commands.

Build commands:

```bash
go build -o codewise-cli main.go
# or
make build
```

Test commands:

```bash
make test
go test ./... -v
```

## Runtime Connectivity And Control Flow

The internal execution flow is:

1. User runs a CLI command (for example `codewise k8s apply`).
2. Cobra command in `cmd/` parses args and flags.
3. Command handler calls a domain package in `pkg/`.
4. Domain package reads files/config and performs requested action.
5. For deployment workflows, external CLIs/services are invoked where needed.
6. Result is printed to terminal and process exits with success or error code.

Conceptual connectivity map:

```text
Terminal User
   -> codewise (binary)
      -> cmd/* (Cobra command layer)
         -> pkg/* (business logic)
            -> Local Filesystem (templates, manifests, config)
            -> External Tools (docker, kubectl, helm)
            -> Kubernetes API (through kubectl/helm operations)
```

## Project Structure And Responsibilities

- `cmd/`: Command definitions, flags, argument validation, command wiring.
- `pkg/config/`: Configuration load/save defaults.
- `pkg/docker/`: Dockerfile init, validation, image build helpers.
- `pkg/k8s/`: Manifest generation, apply/delete orchestration.
- `pkg/helm/`: Helm chart scaffold generation.
- `pkg/deploy/`: Deploy planning, execution, logs, status, history, rollback.
- `pkg/env/`: Environment create/list/delete behaviors.
- `pkg/encoder/`: Format conversion and Base64 utilities.
- `pkg/generator/`: Project and template generation logic.
- `templates/`: Template source files used by template commands.
- `k8s/` and `helm/`: Generated scaffold outputs.
- `tests/` and `testdata/`: Test cases and fixture inputs/outputs.

## Tech Stack

Primary:
- Go 1.20
- Cobra CLI framework (`github.com/spf13/cobra`)
- YAML support (`gopkg.in/yaml.v3`)
- TOML support (`github.com/BurntSushi/toml`)
- XML mapping support (`github.com/clbanning/mxj/v2`)
- Interactive prompts (`github.com/AlecAivazis/survey/v2`)

Tooling and ecosystem dependencies:
- Docker CLI and daemon for image-related workflows
- kubectl and Kubernetes cluster context for K8s apply/delete workflows
- Helm CLI for chart-related workflows

## How Commands Connect To Domains

- `config` commands map to config domain for defaults and view/init operations.
- `docker` commands map to Docker domain for init/validate/build.
- `k8s` commands map to K8s domain for scaffold/apply/delete.
- `helm` commands map to Helm domain for scaffold init.
- `deploy` commands map to deployment domain for run/plan/status/logs/history/rollback.
- `env` commands map to environment domain for profile lifecycle.
- `encode` maps to transformation domain for file format conversions.
- `template` maps to generator/template domain for predefined scaffold files.

## Development And Contribution Model

Typical change flow:

1. Add or update command wiring under `cmd/`.
2. Implement core logic in `pkg/`.
3. Add or update tests under `tests/` with fixtures in `testdata/`.
4. Validate with `go test ./... -v`.
5. Build with `make build` and run manual smoke commands.

This layering keeps CLI surface concerns separate from business logic and improves testability.
