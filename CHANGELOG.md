# Changelog

## v0.3.23 - 2025-10-13

### fix: harden port parsing

Trim whitespace, validate the `PORT` env var before starting the HTTP server, and fall back to the default when invalid. Added tests covering whitespace and invalid values and refreshed default metadata to `v0.3.23`.

## v0.3.22 - 2025-10-10

### test: assert echo JSON error

Added a unit test ensuring `/echo` returns a JSON error and correct content type when the `message` field is empty. Refreshed default/version metadata to `v0.3.22` across docs and config.

## v0.3.21 - 2025-10-10

### fix: return JSON error for invalid echo input

Updated the `/echo` handler to return a proper JSON error payload with `application/json` content type on invalid input (400), aligning behavior with the OpenAPI spec. Refreshed default/version metadata to `v0.3.21` across code, docs, and spec.

## v0.3.20 - 2025-10-09

### feat: allow HEAD requests in CORS

Expanded the CORS middleware to accept HEAD requests and refreshed default metadata to `v0.3.20`.

## v0.3.19 - 2025-10-08

### chore: clarify clean target

Updated `make clean` help text to note it clears coverage files and refreshed default metadata to `v0.3.19`.

## v0.3.18 - 2025-10-08

### chore: clean coverage artifacts

Updated `make clean` to remove generated coverage outputs, documented the behavior, and refreshed default metadata to `v0.3.18`.

## v0.3.17 - 2025-10-08

### chore: add coverage html export

Added a `make coverage-html` helper that writes `coverage.html`, ignored it from Git, documented the workflow, and bumped defaults to `v0.3.17`.

## v0.3.16 - 2025-10-08

### chore: show coverage summary

Enhanced `make coverage` to print the total coverage line and documented how to preview the HTML report, while updating defaults to `v0.3.16`.

## v0.3.15 - 2025-10-08

### chore: add coverage make target

Added a `make coverage` helper that runs `go test` with coverage output, documented it in the README, and refreshed default version metadata to `v0.3.15`.

## v0.3.14 - 2025-10-08

### chore: add go vet lint target

Introduced a `make lint` helper that runs `go vet` for static analysis, documented it in the README, and updated default version metadata to `v0.3.14`.

## v0.3.11 - 2025-09-18

### feat: allow configuring listen port via PORT env

The server now honors the `PORT` environment variable (e.g., `PORT=9090 make run`). Tests and documentation were updated accordingly, and the version metadata now references `v0.3.11`.

## v0.3.10 - 2025-09-18

### docs: add docker usage examples

Documented how to build and run the service via Docker, including port overrides, and updated version references to `v0.3.10`.

## v0.3.9 - 2025-09-18

### docs: note jq dependency for curl helpers

Clarified that the README's curl samples rely on `jq` for formatting and bumped version references to `v0.3.9`.

## v0.3.8 - 2025-09-18

### docs: document configurable port

Added usage examples showing how to override the listen `PORT` in `make run`, plus updated defaults to `v0.3.8` across docs and specs.

## v0.3.7 - 2025-09-18

### docs: add CI smoke test guidance

Documented the `make deps`, `make fmt`, `make test` sequence for local pipeline checks and bumped version metadata to `v0.3.7`.

## v0.3.6 - 2025-09-18

### docs: add troubleshooting tips

Added a quick FAQ in the README covering missing Go tools and port conflicts, and updated version metadata to `v0.3.6`.

## v0.3.5 - 2025-09-18

### docs: add quick curl probes

Documented handy `curl` commands for `/healthz`, `/info`, and `/version` to speed up smoke tests, and refreshed version metadata to `v0.3.5` across docs and defaults.

## v0.3.4 - 2025-09-18

### chore: document targets with make help

Added a `make help` target that prints descriptions for the most common commands, and refreshed docs/config to reference version `v0.3.4`.

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
