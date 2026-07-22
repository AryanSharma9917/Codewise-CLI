# Docker commands

## `docker init`

Generate a multi-stage Go `Dockerfile` in the current directory:

```bash
codewise docker init
```

Codewise will not overwrite an existing Dockerfile.

## `docker validate`

Inspect the Dockerfile for its base image, multi-stage builds, and non-root user configuration:

```bash
codewise docker validate
```

## `docker build`

```bash
codewise docker build
codewise docker build --tag ghcr.io/example/myapp:v1
```

| Flag | Short | Description |
| --- | --- | --- |
| `--tag` | `-t` | Image tag; defaults to `codewise:latest` |

## `docker push`

```bash
codewise docker push --tag ghcr.io/example/myapp:v1
```

Pushing requires registry authentication and an explicit tag.
