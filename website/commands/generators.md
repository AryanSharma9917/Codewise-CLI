# Project and template generators

## `init`

Create a new project directory and initialize its Git repository:

```bash
codewise init --project myapp
codewise init --project myapp --with-docker --with-deployment
```

| Flag | Short | Description |
| --- | --- | --- |
| `--project` | `-p` | Project directory name |
| `--with-docker` | | Add a Dockerfile |
| `--with-deployment` | | Add a Kubernetes Deployment |

## `template github-action`

```bash
codewise template github-action \
  --output .github/workflows/deploy.yaml \
  --app-name myapp \
  --repo https://github.com/example/myapp
```

## `template argo-app`

```bash
codewise template argo-app \
  --output k8s/argo-app.yaml \
  --app-name myapp \
  --repo https://github.com/example/myapp
```

Both template commands require `--output` and accept `--app-name` and `--repo`.
