# Docker to Kubernetes

This workflow packages an application and previews its Kubernetes deployment.

## Generate and validate the Dockerfile

```bash
codewise docker init
codewise docker validate
```

Review the generated file, especially its build command, exposed port, and runtime user.

## Build the image

```bash
codewise docker build --tag ghcr.io/example/myapp:v1
```

Run the image locally before publishing it:

```bash
docker run --rm -p 8080:8080 ghcr.io/example/myapp:v1
```

## Publish the image

```bash
docker login ghcr.io
codewise docker push --tag ghcr.io/example/myapp:v1
```

## Generate and preview manifests

```bash
codewise k8s init --app myapp --image ghcr.io/example/myapp:v1
codewise k8s apply --namespace dev --dry-run
```

## Apply to a test cluster

```bash
kubectl config current-context
codewise k8s apply --namespace dev
```

Clean up with `codewise k8s delete --namespace dev` when the test is complete.
