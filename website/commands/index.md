# Command reference

```bash
codewise <command> [subcommand] [flags]
```

| Command | Purpose |
| --- | --- |
| `config` | Initialize and view global defaults |
| `deploy` | Plan, run, inspect, and roll back deployments |
| `docker` | Generate, validate, build, and push images |
| `encode` | Convert structured files and Base64 data |
| `env` | Create, list, and delete environment profiles |
| `helm` | Generate Helm charts |
| `init` | Scaffold a DevOps-ready project |
| `k8s` | Generate, apply, and delete Kubernetes manifests |
| `template` | Generate GitHub Actions and Argo CD templates |
| `version` | Print or check the CLI version |
| `completion` | Generate shell completion scripts |
| `help` | Show help for Codewise or a specific command |

Run help at any level:

```bash
codewise --help
codewise deploy --help
codewise deploy rollback --help
```

## Global flags

| Flag | Description |
| --- | --- |
| `-h`, `--help` | Show context-sensitive help |
| `-v`, `--version` | Print the CLI version |

## Version and completion

```bash
codewise version
codewise version --latest
codewise completion bash
codewise completion zsh
codewise completion fish
codewise completion powershell
```

| Command | Flags |
| --- | --- |
| `version` | `-l`, `--latest` |
| `completion bash` | `--no-descriptions` |
| `completion fish` | `--no-descriptions` |
| `completion powershell` | `--no-descriptions` |
| `completion zsh` | `--no-descriptions` |

Disable completion descriptions when a shell setup needs smaller generated output:

```bash
codewise completion bash --no-descriptions
```

Ask for help directly or through Cobra's generated help command:

```bash
codewise --help
codewise help deploy
codewise help deploy rollback
```

## Complete command inventory

This table is the exhaustive callable command surface. `--help` is available on every command and is omitted from the per-command flag column to keep the table readable.

| Command path | Arguments and command-specific flags |
| --- | --- |
| `codewise` | `-v`, `--version` |
| `codewise help [command]` | Command path to explain |
| `codewise version` | `-l`, `--latest` |
| `codewise completion bash` | `--no-descriptions` |
| `codewise completion fish` | `--no-descriptions` |
| `codewise completion powershell` | `--no-descriptions` |
| `codewise completion zsh` | `--no-descriptions` |
| `codewise config init` | None |
| `codewise config view` | None |
| `codewise env create <name>` | `-i`, `--interactive` |
| `codewise env list` | None |
| `codewise env delete <name>` | `--yes` |
| `codewise docker init` | None |
| `codewise docker validate` | None |
| `codewise docker build` | `-t`, `--tag <image>` |
| `codewise docker push` | `-t`, `--tag <image>` (required) |
| `codewise k8s init` | `--app <name>`, `--image <image>` |
| `codewise k8s apply` | `--namespace <name>`, `--context <name>`, `--dry-run` |
| `codewise k8s delete` | `--namespace <name>`, `--context <name>`, `--dry-run` |
| `codewise helm init` | `--app <name>`, `--image <image>` |
| `codewise deploy plan` | `--env <name>` (required) |
| `codewise deploy run` | `--env <name>` (required), `--dry-run` |
| `codewise deploy preview` | `--pr <number>` or `--image <image>` (at least one required); `--keep` |
| `codewise deploy status` | `--env <name>` (required) |
| `codewise deploy logs` | `--env <name>` (required), `--follow` |
| `codewise deploy history` | `--env <name>` (required) |
| `codewise deploy rollback` | `--env <name>` (required), `--revision <number>` |
| `codewise encode` | `-i`, `--input <file>` (required); `-o`, `--output <file>` (required); `--json-to-yaml`; `--env-to-json`; `--base64`; `--decode`; `-f`, `--force` |
| `codewise init` | `-p`, `--project <name>`; `--with-docker`; `--with-deployment` |
| `codewise template github-action` | `-o`, `--output <file>` (required); `--app-name <name>`; `--repo <url>` |
| `codewise template argo-app` | `-o`, `--output <file>` (required); `--app-name <name>`; `--repo <url>` |

The parent commands—`config`, `deploy`, `docker`, `env`, `helm`, `k8s`, `template`, and `completion`—display their available subcommands when invoked with `--help`.
