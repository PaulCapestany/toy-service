# Toy Service

Welcome to `toy-service`, a minimal microservice in Go designed to demonstrate best practices in open-source software development, including OpenAPI-driven design, Semantic Versioning, `[GitOps](https://github.com/paulcapestany/gitops)` deployments, and more. As we evolve this service, we’ll illustrate how to incrementally adopt industry-standard tools and workflows to ensure a clear API contract, smooth contributor experience, and reliable CI/CD.

## Key Features and Goals

- **OpenAPI-defined endpoints:**
  The API contract is defined via an [OpenAPI Specification](https://www.openapis.org/). This ensures clarity and consistency for contributors and integrators, allowing automatic client code generation, documentation, and validation.

- **Semantic Versioning (SemVer) Integration:**
  Using [Semantic Versioning](https://semver.org/) makes it clear when changes are backward-compatible, introduce new features, or are breaking. This is crucial for multi-service ecosystems where external consumers depend on stable API contracts.

- **Automated CI/CD with Tekton:**
  Future steps will introduce continuous integration and delivery pipelines, ensuring every commit and pull request is tested, built, and (optionally) deployed automatically.

- **GitOps Deployment with Argo CD:**
  By using Git as the single source of truth for environment configurations and deployments via our `[gitops](https://github.com/paulcapestany/gitops)` repo, we’ll enable near-instant feedback loops for contributors and ensure reliable, versioned rollouts.

- **Environment Configuration & Mocking Strategies:**
  Contributors can easily adjust environment variables to test different configurations locally or in dev environments, without needing complex local setups.

- **Contributor Guidelines and Documentation:**
  We’ll provide clear instructions for setting up development environments, writing tests, following versioning policies, and making good pull requests.

## Project Structure

The code is structured to separate concerns and maintain clarity:

```
toy-service/
├── CHANGELOG.md
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go # Entry point for the service
├── internal/
│   └── handlers/  # Handlers for each endpoint
│       ├── echo.go
│       ├── echo_test.go
│       ├── env.go
│       ├── healthz.go
│       ├── healthz_test.go
│       ├── info.go
│       └── info_test.go
└── spec/
    └── openapi.yaml # OpenAPI specification for this service
```

## Running Locally

Before running, ensure you have Go 1.20+ installed and Docker (optional for containerization).

```shell
make build
make run
```

This will start the service on http://localhost:8080. Current endpoints:

- `GET /healthz`: Basic health check, returns `{ "status": "ok" }`.
- `POST /echo`: Accepts `{ "message": "Hello" }` and returns `{ "message": "Hello [modified]", "version": "...", "commit": "...", "env": "..." }`.
- `GET /info`: Returns metadata about the service, including environment (env), version (version), log verbosity (logVerbosity), fake secret (fakeSecret), and current commit hash (commit).

## Environment Variables

You can influence runtime behavior and metadata via environment variables:

- `SERVICE_ENV` (e.g., dev, prod)
- `LOG_VERBOSITY` (e.g., info, debug)
- `FAKE_SECRET` (e.g., redacted, topsecret)
- `VERSION` (e.g., v0.2.1)
- `GIT_COMMIT` (e.g., abc1234)

For example:

```shell
export SERVICE_ENV=prod
export LOG_VERBOSITY=debug
export FAKE_SECRET=topsecret
export VERSION=v0.2.1
export GIT_COMMIT=abc1234

make run
```

## Example Usage

```shell
curl http://localhost:8080/healthz
# {"status":"ok"}

curl -X POST -H "Content-Type: application/json" -d '{"message":"Hello"}' http://localhost:8080/echo
# {"message":"Hello [modified]","version":"v0.2.1","commit":"unknown","env":"dev"}

curl http://localhost:8080/info
# {"name":"toy-service","version":"v0.2.1","env":"dev","logVerbosity":"info","fakeSecret":"redacted","commit":"unknown"}
```

## Testing

We validate all responses against the OpenAPI specification and run unit/integration tests with:

```shell
make test
```

If a change breaks the API contract (for example, by returning fields not defined in the spec or missing required properties), these tests will fail. Contributors should ensure that any API change is reflected in the OpenAPI spec and that all tests pass before opening a PR.

## Development Workflow

### Branching Strategy

#### Main Branch:
The main branch is the trunk of development. It should always be in a working state, tested, and ready for release.

#### Feature Branches:
Create a new branch (`feature/your-feature-name`) for each new feature, bug fix, or documentation improvement. Work on your changes there, then open a Pull Request (PR) against `main`.

### Pull Requests & Reviews:
When you open a PR, automated tests (and later CI/CD pipelines) run. Once everything passes and at least one reviewer approves, we merge it into `main`.

## Semantic Versioning (SemVer)

We follow SemVer rules:

- **MAJOR (X.0.0):** Incompatible API changes.
- **MINOR (0.X.0):** Backwards-compatible feature additions.
- **PATCH (0.0.X):** Backwards-compatible bug fixes.

### Workflow for SemVer & Releases:

#### Commit Messages & Conventional Commits (Optional but Recommended):
Use commit prefixes like `feat:`, `fix:`, `docs:`, `chore:`, etc. This helps automated tools infer whether to bump MAJOR, MINOR, or PATCH versions.

Example:

- `feat: add /echo endpoint` → MINOR bump
- `fix: correct error handling in /info` → PATCH bump
- `feat!: remove old endpoint (with !)` → MAJOR bump (breaking change)

Check out [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for more details.

#### Tagging Releases:
Once changes are merged to `main` and you’re ready to release, create a Git tag like `v0.2.0`. This tag represents a snapshot of the code at release time. CI/CD will build and push corresponding Docker images and update GitOps manifests.

For early development, you might manually assign versions. Later, you can automate versioning using tools like semantic-release.

#### Updating the CHANGELOG:
Every time you cut a release, update `CHANGELOG.md` to summarize what changed. The CHANGELOG entry is user-facing: just the highlights of what changed and what’s new in the release. For example:

```md
## v0.2.0 - YYYY-MM-DD
- Added /echo endpoint
- Added /info endpoint
- Introduced environment variables for configuration
```

This helps contributors and users know what’s new, what’s fixed, and when breaking changes occur.

## Integrating OpenAPI Changes

When you add or modify an endpoint:

- Update the `spec/openapi.yaml` file so the specification remains the single source of truth for API behavior.
- Confirm that your changes align with SemVer. Adding a new endpoint in a backward-compatible manner is a MINOR release. Removing or changing an existing endpoint’s contract in a backward-incompatible way requires a MAJOR release bump.

## Future Enhancements

- **CI/CD Pipelines:**
  We will introduce Tekton or GitHub Actions workflows that run tests, verify code coverage, lint code, and build Docker images automatically on every PR and push.

- **GitOps with Argo CD:**
  In a future iteration, you’ll see a separate Git repository managing Kubernetes manifests for toy-service. Argo CD will continuously reconcile the cluster state with the declared manifests, automatically deploying new versions as they appear.

- **Pre-Release Tags:**
  For testing in dev environments, you may use pre-release tags (`v0.2.0-alpha.1`) to indicate that this is a work-in-progress or experimental build. Contributors can test against these pre-releases before stable versions are cut.

## Contributing

1. Fork and Clone this repository.
2. Create a Feature Branch:

```shell
git checkout -b feature/my-new-endpoint
```

3. Implement and Test:
   Add or modify handlers, update `openapi.yaml`, run `make test`.

4. Open a Pull Request:
   Once done, push your branch and open a PR against `main`.

5. Code Review & Merge:
   After tests pass and review is complete, changes are merged. If necessary, we tag and release a new version.

## License

This project is free and open source :)

