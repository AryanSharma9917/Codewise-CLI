# Helm deployment

## Create the chart

```bash
codewise helm init --app myapp --image ghcr.io/example/myapp:v1
```

Inspect and customize:

```text
helm/chart/
├── Chart.yaml
├── values.yaml
└── templates/
    ├── deployment.yaml
    └── service.yaml
```

## Create an environment

```bash
codewise env create staging --interactive
codewise deploy plan --env staging
```

Codewise selects Helm automatically when `helm/chart/` exists and GitOps is not configured.

## Preview and deploy

```bash
codewise deploy run --env staging --dry-run
kubectl config current-context
codewise deploy run --env staging
```

## Operate the release

```bash
codewise deploy status --env staging
codewise deploy history --env staging
codewise deploy rollback --env staging --revision 2
```
