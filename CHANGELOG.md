# Changelog

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