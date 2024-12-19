# Contributing to the bitiq Ecosystem

Welcome to the bitiq ecosystem! We appreciate your interest in contributing. This guide applies to all bitiq subprojects, including `toy-service`, `toy-web`, and many others that may emerge over time.

## Table of Contents

- [Contributing to the bitiq Ecosystem](#contributing-to-the-bitiq-ecosystem)
  - [Table of Contents](#table-of-contents)
  - [Code of Conduct](#code-of-conduct)
  - [Project Structure and Philosophy](#project-structure-and-philosophy)
  - [Getting Started](#getting-started)
  - [Semantic Versioning \& Conventional Commits](#semantic-versioning--conventional-commits)
  - [Development Workflow](#development-workflow)
  - [Testing and Quality Assurance](#testing-and-quality-assurance)
  - [Documentation](#documentation)
  - [Submitting Changes](#submitting-changes)
  - [Release Process](#release-process)
  - [Support and Contact](#support-and-contact)

## Code of Conduct

All contributors are expected to uphold a respectful, inclusive environment. For full details, please see our [Code of Conduct](./CODE_OF_CONDUCT.md) (not yet implemented, but planned).

## Project Structure and Philosophy

bitiq is composed of multiple microservices and frontends, each defined with OpenAPI specs and focused on being cloud-agnostic. By having a consistent structure (README, TODOs, CHANGELOG, docs), we help newcomers ramp up quickly and ensure that all parts of the ecosystem feel familiar.

## Getting Started

1. **Clone the Repository:**  
   Start by forking and cloning the particular subproject (`toy-service`, `toy-web`, etc.).

2. **Set Up Your Environment:**  
   Refer to the project’s `README.md` for environment prerequisites, dependencies, and local run instructions.

## Semantic Versioning & Conventional Commits

We strictly follow [Semantic Versioning (SemVer)](https://semver.org/) and [Conventional Commits](https://www.conventionalcommits.org/) to communicate the impact of changes clearly and maintain clean release notes.

- **MAJOR:** Incompatible API changes.
- **MINOR:** Backward-compatible feature additions.
- **PATCH:** Backward-compatible bug fixes or documentation improvements.

Commit types include `feat`, `fix`, `docs`, `chore`, `refactor`, `perf`, `test`. Append `!` for breaking changes, e.g., `feat!: ...`.

## Development Workflow

1. **Create a Branch:** Branch off `main` for any changes.
2. **Code & Test:** Implement features or fixes and write tests.
3. **Pull Request (PR):** Open a PR for review. CI will run tests automatically.
4. **Feedback & Revisions:** Expect reviews and possibly requested changes.
5. **Merge & Release:** Once approved, maintainers merge and update `CHANGELOG.md` according to SemVer.

## Testing and Quality Assurance

Each subproject includes instructions on running tests (e.g., `make test` or similar). We aim for comprehensive unit testing, integration testing, and OpenAPI conformance tests where applicable.

For frontends (like `toy-web`), tests may be minimal at first, but we welcome contributions to introduce lightweight frameworks like Cypress or Jest.

## Documentation

- **OpenAPI Specs:** Serve as the primary source of API documentation.
- **`docs/` Directory:** Additional guides, architecture overviews, and contributor tutorials may be found here.

As bitiq grows, we may produce global documentation resources accessible from a central location.

## Submitting Changes

- Open a PR against `main`.
- Follow Conventional Commits for your commit messages.
- Update `CHANGELOG.md` upon completion of a TODO item.

## Release Process

Merges to `main` may trigger automated CI/CD and GitOps processes for test environments. Actual stable releases are cut with SemVer tags, and `CHANGELOG.md` entries are updated accordingly.

## Support and Contact

For questions or help, open an issue in the relevant repository. We encourage an open and collaborative environment. Feel free to propose improvements, features, or ask for clarification at any time.