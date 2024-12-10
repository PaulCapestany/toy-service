# Toy Service

This is a minimal toy microservice written in Go to demonstrate an incremental approach to building a well-structured, cloud-native application. It will eventually showcase:

- OpenAPI-defined endpoints
- Semantic versioning integration
- Automated CI/CD with Tekton or GitHub Actions
- GitOps deployment via Argo CD
- Environmental configuration and mocking strategies
- Documentation and contributor guidelines

## Project Structure

The following structure separates the main application entry point (cmd/server/main.go) from internal logic (internal/handlers), making it easier to scale and test. Tests reside in a tests directory. As the service grows, we can add more packages and directories.

```
toy-service/
├─ cmd/
│  └─ server/
│     └─ main.go          # main entry point
├─ internal/
│  └─ handlers/
│     └─ healthz.go       # handler for /healthz endpoint
│     └─ healthz_test.go  # unit tests for healthz handler
├─ go.mod
├─ go.sum
├─ Dockerfile
├─ Makefile
└─ README.md
```

## Current Status

Currently, the service only has a `/healthz` endpoint for verifying basic functionality. Future steps will add `/echo`, `/info`, and the full pipeline.

## Running Locally

```bash
make build
make run