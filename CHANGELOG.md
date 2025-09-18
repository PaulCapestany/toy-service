# Changelog

## v0.3.3 - 2025-09-18

### chore: add gofmt automation

Introduced a `make fmt` target so contributors can format Go sources consistently. Documentation, defaults, and the OpenAPI spec now reflect version `v0.3.3`.

## v0.3.2 - 2025-09-18

### chore: ignore local build artifacts

Added a `.gitignore` to drop compiled binaries, macOS metadata, and log files from version control so the repo stays clean for contributors.

## v0.3.1 - 2025-09-18

### feat: Add `/version` endpoint for lightweight build checks

Introduced a dedicated `GET /version` endpoint that returns the service name, semantic version, and git commit hash. The OpenAPI specification, router, and tests were extended so automation can verify deployments without parsing the broader `/info` payload. Defaults now point to `v0.3.1`.

## v0.3.0 - 2025-09-17

### feat: Add make clean target for removing build artifacts

Introduce a `make clean` goal that deletes the `./bin` directory, making it easier to reset the workspace between builds.

## v0.2.22 - 2024-12-19

### fix: Add CORS support so that `toy-web` can communicate with `toy-service` in local development

Now the server sets `Access-Control-Allow-Origin: *` and related headers for all endpoints, enabling the frontend to send requests without CORS errors.

## v0.2.1 - 2024-12-11

### fix: Enforce non-empty `message` constraint in the `/echo` endpoint. 
  
The server now returns a `400 Bad Request` if the `message` field is missing or empty, ensuring compliance with the OpenAPI specification. (Previously, it returned `200` even when `message` was invalid.)

## v0.2.0 - 2024-12-10

- Added `/echo` endpoint for echoing and modifying messages
- Added `/info` endpoint for returning environment-based metadata
- Introduced environment variables (SERVICE_ENV, LOG_VERBOSITY, FAKE_SECRET, VERSION, GIT_COMMIT) for runtime configuration
- Added corresponding tests and updated documentation examples
- Create OpenAPI Specification within openapi.yaml
- Improved documentation in README.md

## v0.1.0 - 2024-12-10
- Initial release of toy-service with:
  - /healthz endpoint
  - Basic test coverage
    - Dockerfile, Makefile, and README
