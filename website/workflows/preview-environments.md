# Preview environments

Preview environments create a temporary Codewise environment for a pull request or image.

```bash
codewise deploy preview --pr 42
codewise deploy preview --image ghcr.io/example/myapp:pr-42
```

Codewise derives a unique name, configures a matching Kubernetes namespace, deploys it, and removes the local environment profile afterward.

Preserve the profile for investigation:

```bash
codewise deploy preview --pr 42 --keep
```

## CI considerations

- Build and publish the preview image before running the command.
- Use a dedicated non-production Kubernetes context.
- Give CI only the namespace-scoped permissions it needs.
- Add an independent cleanup job for abandoned preview namespaces.
- Never expose registry or kubeconfig credentials in command output.
