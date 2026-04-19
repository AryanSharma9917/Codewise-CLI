# Codewise CLI Testing Runbook

This runbook is designed for direct copy-paste testing from project root.

## 1) Prerequisites

Use these checks to confirm required tooling is available before running any workflow.

```bash
go version
docker --version
kubectl version --client
helm version
```

## 2) Build And Baseline Validation

Build the binary, then verify command wiring and version output.

```bash
make clean
make build
./codewise-cli --help
./codewise-cli version
go run . --help
go run . version
```

## 3) Automated Tests

Run the existing test suites first, then optional race and coverage checks.

```bash
make test
go test ./... -v
go test ./... -race
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## 4) Top-Level Command Surface Check

These commands confirm every top-level command is registered and callable.

```bash
./codewise-cli config --help
./codewise-cli deploy --help
./codewise-cli docker --help
./codewise-cli encode --help
./codewise-cli env --help
./codewise-cli helm --help
./codewise-cli init --help
./codewise-cli k8s --help
./codewise-cli template --help
```

## 5) Config Workflow

This validates local configuration initialization and read-back.

```bash
./codewise-cli config init
./codewise-cli config view
cat ~/.codewise/config.yaml
```

## 6) Encode/Convert Workflow (Fixture-Based)

These commands validate common encode paths using files in `testdata/`.

```bash
./codewise-cli encode -i testdata/sample.yaml -o /tmp/cw_sample_1.json
./codewise-cli encode -i testdata/sample.json -o /tmp/cw_sample_2.yaml --json-to-yaml
./codewise-cli encode -i testdata/sample.env -o /tmp/cw_sample_3.json --env-to-json
./codewise-cli encode -i testdata/sample.txt -o /tmp/cw_sample_4.b64 --base64
./codewise-cli encode -i /tmp/cw_sample_4.b64 -o /tmp/cw_sample_5.txt --base64 --decode
```

Quick output validation for generated artifacts.

```bash
ls -lh /tmp/cw_sample_1.json /tmp/cw_sample_2.yaml /tmp/cw_sample_3.json /tmp/cw_sample_4.b64 /tmp/cw_sample_5.txt
diff -u testdata/sample.txt /tmp/cw_sample_5.txt
```

## 7) Scaffold Generation (Local Safe)

Use this to validate project scaffold generation for Kubernetes and Helm.

```bash
rm -rf helm/chart k8s/app
./codewise-cli helm init
./codewise-cli k8s init
find helm -maxdepth 4 -type f | sort
find k8s -maxdepth 4 -type f | sort
```

## 8) Template Commands

These validate template generation commands.

```bash
./codewise-cli template github-action
./codewise-cli template argo-app
```

## 9) Docker Workflow (Requires Docker Daemon)

Run these if Docker is available and running on your machine.

```bash
./codewise-cli docker init
./codewise-cli docker validate
./codewise-cli docker build
docker images | grep -i codewise
```

## 10) Kubernetes Workflow (Dry-Run First)

Use dry-run for safe validation without applying to the cluster.

```bash
./codewise-cli k8s apply --dry-run
./codewise-cli k8s delete --dry-run
```

## 11) Kubernetes Real Apply (Only On Test Cluster)

Run this section only when your kube context points to a non-production cluster.

```bash
kubectl config current-context
./codewise-cli k8s apply --namespace dev
./codewise-cli deploy status --namespace dev
./codewise-cli deploy logs --namespace dev
./codewise-cli deploy history --namespace dev
./codewise-cli deploy rollback --namespace dev
./codewise-cli k8s delete --namespace dev
```

## 12) Environment Command Validation

Use help checks and simple actions to validate env command behavior.

```bash
./codewise-cli env --help
./codewise-cli env list
./codewise-cli env create dev
./codewise-cli env list
./codewise-cli env delete dev
./codewise-cli env list
```

## 13) Negative Testing

These commands validate error handling and message quality for invalid inputs.

```bash
./codewise-cli encode -i does-not-exist.yaml -o /tmp/out.json
./codewise-cli k8s apply --context definitely-not-real --dry-run
./codewise-cli deploy status --namespace does-not-exist
```

## 14) One-Line Regression Pack

Use this compact command for quick regression checks before pushing changes.

```bash
make build && go test ./... -v && ./codewise-cli encode -i testdata/sample.yaml -o /tmp/reg.json && ./codewise-cli helm init && ./codewise-cli k8s init && ./codewise-cli k8s apply --dry-run
```

## Expected Success Signals

- Build exits with code 0.
- Test commands pass with code 0.
- Encode commands generate output files and decode roundtrip matches input.
- Helm and Kubernetes scaffold files are generated without runtime errors.
- Dry-run commands do not mutate cluster resources and still validate command flow.
