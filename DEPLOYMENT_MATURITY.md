# Deployment Maturity Ladder

This guide defines a practical path for deploying applications with Codewise CLI.

The goal is to move safely through 3 stages:
- Local (developer machine)
- Staging (shared non-production cluster)
- Production (high-confidence rollout)

## Stage 1: Local Deployment Ready

Use this stage to prove command flow and basic behavior.

### Exit Criteria
- The project builds successfully.
- Automated tests pass.
- Kubernetes and Helm scaffolds generate correctly.
- Dry-run commands work without errors.
- At least one environment profile exists and loads correctly.

### Suggested Checks
Run from project root:

1. make build
2. go test ./... -v
3. ./codewise-cli helm init
4. ./codewise-cli k8s init
5. ./codewise-cli k8s apply --dry-run
6. ./codewise-cli deploy plan --env dev

### Notes
- This stage should not mutate any production resources.
- Keep context pointed to local clusters only (kind, minikube, k3d, etc.).

## Stage 2: Staging Deployment Ready

Use this stage to validate behavior in a real shared cluster.

### Exit Criteria
- Environment-specific deployment works end-to-end.
- Rollout monitoring succeeds for all target deployments.
- Status, logs, and history commands are usable for troubleshooting.
- Rollback path is tested at least once on staging.
- Team can reproduce deployment steps from docs without tribal knowledge.

### Suggested Checks
1. ./codewise-cli env create staging
2. ./codewise-cli deploy plan --env staging
3. ./codewise-cli deploy run --env staging
4. ./codewise-cli deploy status --env staging
5. ./codewise-cli deploy logs --env staging --follow
6. ./codewise-cli deploy history --env staging
7. ./codewise-cli deploy rollback --env staging --revision <known-good-revision>

### Notes
- Use a dedicated staging namespace and cluster context.
- Test rollback before any production attempt.

## Stage 3: Production Deployment Ready

Use this stage only when reliability and observability are strong.

### Exit Criteria
- Deploy package has automated tests for deploy logic (strategy resolution, command building, error paths).
- Preflight checks fail fast with actionable error messages.
- CI blocks merges on test failures.
- Rollback runbook is documented and time-to-recovery is known.
- A production promotion checklist exists and is followed.

### Minimum Promotion Checklist
1. Confirm current kube context and namespace are correct.
2. Run deployment plan for target environment.
3. Verify image tag and chart/manifest inputs.
4. Deploy with a human-in-the-loop approval.
5. Confirm rollout and health signals.
6. Confirm logs and status after rollout.
7. Keep rollback command prepared with known-good revision.

## Should We Deploy Now?

- Yes for Local and Staging progression.
- No for full Production usage until deployment test coverage and runbook consistency are improved.

## What To Do Next In This Repo

1. Add deployment-focused automated tests under tests/deploy.
2. Standardize docs and examples to use env-based deploy flags consistently.
3. Improve deploy command output where subprocess stderr/stdout is currently suppressed.
4. Add a release workflow for shipping the CLI binary (separate from app deployment).

This sequence gives fast progress without taking unnecessary production risk.