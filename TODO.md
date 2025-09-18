# TODOs for toy-service Repository

This document outlines planned enhancements and maintenance tasks for `toy-service`.

## Guidelines

- **Semantic Versioning (SemVer)** and **Conventional Commits** rules as described in the global [CONTRIBUTING.md](./CONTRIBUTING.md).
  
- **Upon Completion:**  
  When a task is completed:
  - Remove it from `TODO.md`.
  - Update `CHANGELOG.md` with the version increment and commit message.

---

## Upcoming Tasks

### MINOR Enhancements (Non-Breaking Feature Additions)

- **feat: provide a /metrics endpoint (Prometheus format)**  
  Introduce basic Prometheus-compatible metrics (e.g., request counts, latencies) to improve observability and performance monitoring.

### PATCH Improvements (Backward-Compatible Fixes, Docs, Chores)

- **docs: add a CONTRIBUTING.md with detailed contribution guidelines**  
  Create a `CONTRIBUTING.md` to clarify how to propose changes, run tests locally, follow code style guidelines, and name branches.

- **docs: enhance inline code comments**  
  Improve `godoc` comments in `internal/handlers/*`, making the codebase more maintainable and developer-friendly.

- **chore: enhance OpenAPI spec with more examples**  
  Add richer request/response examples to `spec/openapi.yaml` for better understanding by consumers and testers.

- **chore: integrate static analysis (golangci-lint)**  
  Add a `Makefile` target and GitHub Action to run `golangci-lint`, ensuring consistent code quality and reducing technical debt.

### Long-Term Goals (Potential Future MAJOR or MINOR)

- **feat!: restructure request/response models into a shared `pkg/models` package**  
  Standardize data models by placing them in `pkg/models`. This may break compatibility if external consumers rely on current response formats, warranting a MAJOR version bump if implemented.
