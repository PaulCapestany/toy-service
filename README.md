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

- **Semantic Versioning & Conventional Commits:**  
  We strictly follow [SemVer](https://semver.org/) and encourage [Conventional Commits](https://www.conventionalcommits.org/) to communicate change impact clearly.

- **Automated CI/CD & GitOps Integrations (Planned):**  
  Although the current repo provides only a baseline, it’s designed to integrate smoothly with CI/CD pipelines and GitOps workflows in the future.

- **Robust Testing & Validation:**  
  Comprehensive tests ensure the service meets its OpenAPI contract and that changes are safe and well-defined.

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
make build
make run

# Clear build artifacts if you need a fresh build
make clean
```

By default, the service runs at http://localhost:8080.

### Example Endpoints

- **GET /healthz:** Check if the service is running.
- **POST /echo:** Accepts a JSON `{"message":"..."}`, returns modified message plus version info.
- **GET /info:** Returns environment, version, commit hash, and more.
- **GET /version:** Lightweight health/version probe that returns only the service name, version, and commit hash.

### Environment Variables

Control runtime behavior via:
- `SERVICE_ENV` (e.g., dev, prod)
- `LOG_VERBOSITY` (e.g., info, debug)
- `FAKE_SECRET` (e.g., topsecret, redacted)
- `VERSION` (e.g., v0.3.1)
- `GIT_COMMIT` (e.g., abc1234)

**Example:**
```bash
export SERVICE_ENV=prod
export LOG_VERBOSITY=debug
export FAKE_SECRET=topsecret
export VERSION=v0.3.1
export GIT_COMMIT=abc1234

make run
```

### Testing & Validation

```bash
make test
```

Tests verify that handlers respond correctly, match the OpenAPI spec, and respect the contract defined in `spec/openapi.yaml`.

### Development Workflow

- **Branch from `main`** for new features/fixes.
- **Open a Pull Request:** CI runs tests automatically. Once approved, changes are merged.
- **Versioning & Releases:** Use semantic versioning and, optionally, conventional commit messages to guide release processes.
- **GitOps Integration:** Changes in `main` can be automatically deployed to dev/test environments via a separate GitOps repository (e.g., using Argo CD in the future).

## Related Projects

- **[toy-web](https://github.com/paulcapestany/toy-web)**: A planned minimalistic frontend companion, designed as a simple HTML/JS client that interacts with `toy-service`. Ideal starting point for designers and frontend-focused contributors.
- **[gitops](https://github.com/paulcapestany/gitops)**: Manages CI/CD (i.e. building and deployment) of the bitiq microservices using GitOps principles. Ideal starting point for SREs and infrastructure-focused contributors.
 
## Contributing

See `TODO.md` for upcoming tasks and `CHANGELOG.md` for past changes. Additional contributor details will be provided in `CONTRIBUTING.md` (as per planned TODOs).

Fork, clone, and start experimenting. We welcome new contributors!

## License

This project is free and open source.
