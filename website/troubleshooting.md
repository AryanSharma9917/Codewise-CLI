# Troubleshooting

## Start with diagnostics

```bash
codewise version
docker --version
kubectl version --client
helm version
kubectl config current-context
```

## `Dockerfile already exists`

`docker init` refuses to overwrite an existing file. Review or rename the existing Dockerfile before generating another.

## Docker daemon permission denied

Confirm Docker is running and that your user can access its daemon:

```bash
docker info
```

Follow your operating system's Docker post-installation guidance instead of running Codewise permanently as root.

## No Kubernetes manifests found

Run from the application root and generate manifests first:

```bash
codewise k8s init --app myapp --image myapp:latest
codewise k8s apply --dry-run
```

## Cluster unreachable or misconfigured

```bash
kubectl config get-contexts
kubectl config current-context
kubectl cluster-info
```

Pass a specific context with `--context` where supported, or update the environment's `k8s.yaml`.

## Environment not found

```bash
codewise env list
codewise env create dev
```

If `CODEWISE_HOME` is set, Codewise reads environments from that location instead of `~/.codewise`.

## Output file already exists

Use a different output path or opt in to replacement:

```bash
codewise encode -i app.yaml -o app.json --force
```

## Validate the project itself

From the Codewise repository:

```bash
make build
go test ./... -v
go test ./... -race
```
