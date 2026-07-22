# Configuration

Initialize and inspect the global configuration:

```bash
codewise config init
codewise config view
```

The file is stored at `~/.codewise/config.yaml`.

```yaml
version: v1
user:
  name: aryan
defaults:
  app_name: myapp
  image: codewise
  image_tag: latest
  repo_url: https://github.com/example/repo
  namespace: default
  context: ""
  branch: main
```

## Environment profiles

An environment separates Kubernetes, Helm, GitOps, and image settings:

```text
~/.codewise/envs/dev/
├── gitops.yaml
├── helm.yaml
├── k8s.yaml
└── values.yaml
```

Create one non-interactively or answer guided prompts:

```bash
codewise env create dev
codewise env create staging --interactive
```

For automation and tests, set `CODEWISE_HOME` to store environments outside your normal Codewise directory:

```bash
export CODEWISE_HOME=/tmp/codewise-ci
codewise env create test
```

Command flags take precedence over configured defaults when both are available.
