# Quick start

This walkthrough creates local deployment assets and previews Kubernetes operations without changing a cluster.

## 1. Build Codewise

```bash
make build
./codewise-cli version
```

## 2. Initialize configuration

```bash
./codewise-cli config init
./codewise-cli config view
```

The configuration is stored at `~/.codewise/config.yaml`.

## 3. Generate application assets

From your application directory:

```bash
codewise docker init
codewise k8s init --app myapp --image ghcr.io/example/myapp:latest
codewise helm init --app myapp --image ghcr.io/example/myapp:latest
```

Review the generated `Dockerfile`, `k8s/app/`, and `helm/chart/` files before deployment.

## 4. Create an environment

```bash
codewise env create dev
codewise env list
```

Environment files live under `~/.codewise/envs/dev/` by default.

## 5. Preview operations

```bash
codewise k8s apply --namespace dev --dry-run
codewise deploy plan --env dev
codewise deploy run --env dev --dry-run
```

Dry-run deployment does not contact the cluster or create namespaces.

## 6. Deploy when ready

Confirm your Kubernetes context first:

```bash
kubectl config current-context
codewise deploy run --env dev
```

::: warning
Use a non-production cluster until you have reviewed the generated assets and deployment plan.
:::
