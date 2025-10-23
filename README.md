# Toy Service

Welcome to **`toy-service`**, a minimal Go microservice that demonstrates best practices for open-source software development within the bitiq ecosystem.

## Purpose & Context

- **Serves as a backend reference implementation:**  
  Shows how to define and implement RESTful APIs using OpenAPI, practice Semantic Versioning, and maintain clear contributor workflows.

- **Stable, Documented APIs:**  
  By using OpenAPI specifications and automated tests, `toy-service` helps contributors understand how to keep the API contract stable and well-documented.

This project pairs well with [toy-web](https://github.com/paulcapestany/toy-web) (a frontend demo) to give newcomers a complete, end-to-end example.

## Key Features

- **OpenAPI-defined endpoints:**  
  The API contract is clearly defined in `spec/openapi.yaml`, aiding clarity and validation.

- **Live reload of secrets (opt‑in):**  
  When running on Kubernetes, the service can re‑read a mounted Secret at runtime via a minimal reload webhook (`POST /-/reload`). This pairs with a tiny sidecar (e.g., `configmap-reload`) that watches the mounted secret directory and POSTs the endpoint on change. See Live Secret Reload below.

- **Semantic Versioning & Conventional Commits:**  
  We strictly follow [SemVer](https://semver.org/) and encourage [Conventional Commits](https://www.conventionalcommits.org/) to communicate change impact clearly.

- **Automated CI/CD & GitOps Integrations (Planned):**  
  Although the current repo provides only a baseline, it’s designed to integrate smoothly with CI/CD pipelines and GitOps workflows in the future.

- **Robust Testing & Validation:**  
  Comprehensive tests ensure the service meets its OpenAPI contract and that changes are safe and well-defined.

## Deployment via GitOps

This repository intentionally excludes Kubernetes manifests. Runtime deployment configuration lives in [bitiq-io/gitops](https://github.com/bitiq-io/gitops) under `charts/toy-service/`. Update that repository when you need to change image tags, env overlays, or pause Argo CD Image Updater.

When issues or pull requests here depend on deployment changes, include links to the relevant GitOps chart or values files so reviewers can follow the rollout. Keeping manifests in the GitOps repo ensures Argo CD remains the single source of truth—please do not add cluster YAML directly to `toy-service`.

## Project Structure

```text
toy-service/
├── CHANGELOG.md
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go          // Entry point for the service
├── internal/
│   └── handlers/            // HTTP handlers for each endpoint
│       ├── echo.go
│       ├── info.go
│       ├── healthz.go
│       └── ..._test.go
└── spec/
    └── openapi.yaml         // OpenAPI definition of the service's API
```

## Usage

**Prerequisites:**
- Go 1.20+
- Docker (optional for containerization)

**Steps:**
```bash
make help
make deps      # download Go modules before building
make build
make fmt       # format Go source files
make lint      # run go vet static checks
make run

# Clear build artifacts and coverage reports if you need a fresh build
make clean

# Run the test suite
make test

# Generate coverage profile (coverage.out) and summary
make coverage
# Export HTML coverage report (coverage.html)
make coverage-html

# Override the listen port (defaults to 8080)
PORT=9090 make run

# Need a refresher on available commands?
make help
```

By default, the service runs at http://localhost:8080.

### Example Endpoints

- **GET /healthz:** Check if the service is running.
- **POST /echo:** Accepts a JSON `{"message":"..."}`, returns modified message plus version info.
- **GET /info:** Returns environment, version, commit hash, and more.
- **GET /version:** Lightweight health/version probe that returns only the service name, version, and commit hash.
 - **GET /internal/config:** Internal-only helper that reports whether `FAKE_SECRET` is present (and its length), without exposing the value.
 - **POST /-/reload:** Reloads secrets from a mounted directory into process env (see Live Secret Reload).

#### Quick API Checks

Run these curl commands after `make run` (or when the service is deployed) to confirm the API is responding:

```bash
# Basic health probe
curl -s http://localhost:8080/healthz | jq

# Metadata dump
curl -s http://localhost:8080/info | jq

# Lightweight version check
curl -s http://localhost:8080/version | jq

# Secret presence (internal)
curl -s http://localhost:8080/internal/config | jq

# When the service runs inside Docker, use host.docker.internal instead of localhost
curl -s http://host.docker.internal:8080/healthz | jq

# Fail fast when endpoints are unavailable
curl -sf http://localhost:8080/healthz

## Kubernetes Quick Tip

After the GitOps repo deploys `toy-service`, you can port-forward the ClusterIP Service to exercise the API locally without exposing an ingress:

```bash
kubectl -n default port-forward svc/toy-service 8080:8080
```
```

> These examples use `jq` for pretty-printing; install it or drop the pipe if unavailable.

### Docker Usage

Build and run the containerized service with:

```bash
docker build -t toy-service:latest .
docker run -p 8080:8080 --rm toy-service:latest

# Override the listen port exposed from the container
PORT=9090 docker run -e PORT=9090 -p 9090:9090 --rm toy-service:latest

# Or use the Makefile helpers
make docker-build
make docker-run

# Stream container logs while debugging
docker logs -f $(docker ps -q --filter ancestor=toy-service:latest)
```

### Environment Variables

Control runtime behavior via:
- `SERVICE_ENV` (e.g., dev, prod)
- `LOG_VERBOSITY` (e.g., info, debug)
- `FAKE_SECRET` (e.g., topsecret, redacted)
  - When using file‑based reloads, this is set dynamically by `/-/reload` and does not need to be provided at process start.
- `VERSION` (e.g., v0.3.31)
- `PORT` (e.g., 8080)
- `GIT_COMMIT` (e.g., abc1234)

`LOG_VERBOSITY` defaults to `info`, so set it to `debug` (or higher) when you need extra detail.
Valid values include `debug`, `info`, `warn`, and `error`.
`SERVICE_ENV` defaults to `dev`, so override it when targeting staging or production.
`PORT` defaults to `8080`; change it when running multiple services locally.
`FAKE_SECRET` defaults to `redacted`, so provide a real value for integration tests that rely on it.
`GIT_COMMIT` defaults to `unknown` when running from source without CI metadata.

**Example:**
```bash
export SERVICE_ENV=prod
export LOG_VERBOSITY=debug
export FAKE_SECRET=topsecret
export VERSION=v0.3.31
export GIT_COMMIT=abc1234
export PORT=9090

make run
```

### Logging

The server emits structured JSON logs via [`zerolog`](https://github.com/rs/zerolog); tail `stdout` to inspect runtime events.
Pipe through `jq -c` for compact pretty-printing during local debugging.
Each entry includes `level`, `time`, and `message` fields so you can filter quickly.
For example, `jq 'select(.level=="error")'` will highlight only failures during tests.
Keep production runs at `LOG_VERBOSITY=info` to avoid excessive noise.
Timestamps default to Unix seconds because `zerolog` is configured with `zerolog.TimeFormatUnix` in `cmd/server/main.go`.

### Testing & Validation

```bash
make lint
make test

# Optional: quick coverage check with summary
make coverage
# Export annotated HTML report to coverage.html
make coverage-html
# View the report in your browser (macOS example; use xdg-open on Linux)
open coverage.html

# Run a targeted test for quicker feedback
go test ./cmd/server -run TestResolveAddr -v

# Lint the makefile itself (GNU make 4.4+)
make -n help
```

Tests verify that handlers respond correctly, match the OpenAPI spec, and respect the contract defined in `spec/openapi.yaml`.

### Live Secret Reload (Kubernetes)

In Kubernetes, prefer mounting Secrets as files (not env vars) so rotations can be applied without pod restarts. `toy-service` includes an opt‑in webhook (`POST /-/reload`) that re‑reads the `FAKE_SECRET` file from a mounted directory and updates the process environment so subsequent handler calls observe the new value via `os.Getenv`. The endpoint returns JSON confirming the reload and exposes the resulting `fakeSecretLen` for quick verification (never the secret value).

Defaults:

- Secret mount directory: `/etc/backend-secret`
- Secret key: `FAKE_SECRET` (file path `/etc/backend-secret/FAKE_SECRET`)
- Override base directory via `SECRET_FILE_DIR` (optional).

Example curl (local run):

```bash
# Set an initial value, run the server, then change the file and POST reload
export FAKE_SECRET=one
make run &
sleep 1
curl -s http://localhost:8080/internal/config | jq  # shows presence and length

# Simulate a file-mounted secret (for local only)
mkdir -p /tmp/secret && echo -n two >/tmp/secret/FAKE_SECRET
SECRET_FILE_DIR=/tmp/secret curl -s -X POST http://localhost:8080/-/reload
# => {"status":"ok","fakeSecretLen":3}
curl -s http://localhost:8080/internal/config | jq  # reflects new length
```

In Kubernetes, pair the endpoint with a sidecar such as `configmap-reload`:

```yaml
# container excerpt
- name: configmap-reload
  image: ghcr.io/jimmidyson/configmap-reload:latest
  args:
    - --volume-dir=/etc/backend-secret
    - --webhook-url=http://127.0.0.1:8080/-/reload
    - --webhook-method=POST
  volumeMounts:
    - name: backend-secret
      mountPath: /etc/backend-secret
      readOnly: true
```

Security notes:
- Never log secret values; `toy-service` only exposes presence/length via `/internal/config`.
- For services that cannot reload safely (e.g., DB drivers that read once), prefer orchestrated rolling restarts. If you use HashiCorp VSO, set `spec.rolloutRestartTargets` on the `VaultStaticSecret` to trigger a targeted restart only when the secret changes.

### Troubleshooting

- **`go: command not found`** – Install Go 1.20+ and ensure it’s on your `PATH`, then rerun `make deps`.
- **`gofmt: command not found`** – Go’s toolchain bundles `gofmt`; once Go is installed the `make fmt` target works.
- **Ports already in use** – Another process might occupy `8080`; set `PORT` and update `cmd/server/main.go` or stop the conflicting service.
- **Need more verbose logs?** – Set `LOG_VERBOSITY=debug` before `make run` to see request traces while troubleshooting.
- **`jq: command not found`** – Drop the `| jq` suffix from the curl examples or install it via your package manager (e.g., `brew install jq`).

### Development Workflow

- **Branch from `main`** for new features/fixes.
- **Open a Pull Request:** CI runs tests automatically. Once approved, changes are merged.
- **Versioning & Releases:** Use semantic versioning and, optionally, conventional commit messages to guide release processes.
- **GitOps Integration:** Changes in `main` can be automatically deployed to dev/test environments via a separate GitOps repository (e.g., using Argo CD in the future).

### CI/CD Smoke Test

When validating pipelines manually, mimic the GitHub Actions jobs locally:

```bash
make deps
make fmt
make test
```

These commands align with the default CI workflow and help catch issues before pushing.

## Related Projects

- **[toy-web](https://github.com/paulcapestany/toy-web)**: A planned minimalistic frontend companion, designed as a simple HTML/JS client that interacts with `toy-service`. Ideal starting point for designers and frontend-focused contributors.

 
- **[gitops](https://github.com/paulcapestany/gitops)**: Manages CI/CD (i.e. building and deployment) of the bitiq microservices using GitOps principles. Ideal starting point for SREs and infrastructure-focused contributors.
 
## Contributing

See `CONTRIBUTING.md` for workflow, commit conventions, and release guidance. Also check `TODO.md` for upcoming tasks and `CHANGELOG.md` for past changes.

Fork, clone, and start experimenting. We welcome new contributors!

## License

This project is free and open source.
