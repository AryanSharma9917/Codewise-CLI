# Installation

## Requirements

- Go 1.25 or newer to build the current module
- Docker for image builds and pushes
- `kubectl` plus a configured cluster for Kubernetes operations
- Helm for Helm-based deployments

Only Go is required for local scaffolding, encoding, help, and version commands.

## Build from source

```bash
git clone https://github.com/aryansharma9917/codewise-cli.git
cd codewise-cli
make build
./codewise-cli version
```

You can also build directly:

```bash
go build -o codewise-cli main.go
```

## Install on your PATH

Move the compiled binary to a directory already included in your `PATH`:

```bash
sudo install -m 0755 codewise-cli /usr/local/bin/codewise
codewise --help
```

## Verify external tools

```bash
docker --version
kubectl version --client
helm version
```

Codewise reports a clear error when a command requires a missing external tool.
