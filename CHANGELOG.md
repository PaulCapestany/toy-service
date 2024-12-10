# Changelog

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