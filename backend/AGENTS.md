# Backend Agent Guide

## Scope

- This file applies to all work under `backend/`.

## Stack Baseline

- Language: Go 1.26.x.
- HTTP: `net/http` with `chi`.
- Persistence: SQLite with SQL migrations and SQLC-generated accessors.
- Logging: `slog`.
- API docs: spec-first OpenAPI in `backend/openapi/openapi.yaml` with generated artifacts in `backend/docs`.
- Architecture: pragmatic event-driven model with append-only events first and projections added as needed.

## Source Of Truth

- Match changes to the code that exists in this worktree, not to assumed upstream examples.
- Use `go.mod` and `go list -m` to confirm dependency versions before relying on external API docs.

## API Context

- Use `go.mod` and `go list -m` as the source of truth for dependency versions when API behavior may vary by version.
- Prefer `gopls` MCP tooling for Go-aware code navigation, symbol inspection, and refactoring when it is available.
- Prefer `go doc` for local package and symbol documentation.
- Use `pkg.go.dev` only when fuller public module documentation is needed, and make sure the viewed version matches this project.
- Use `backend/openapi/openapi.yaml` as the API contract source of truth.

## Dependency Direction

Use this direction only:

- `domain` -> no internal dependency
- `application` -> `domain` + `ports`
- `ports` -> contracts only
- `adapters` -> `application` + transport mapping
- `infrastructure` -> `ports` or `application` as implementations

Practical rule: `domain` and `application` must not import transport or database packages.

## Architecture Rules

- Keep business rules in `backend/internal/domain`.
- Put use-case orchestration in `backend/internal/application`.
- Keep contracts in `backend/internal/ports`.
- Keep HTTP transport concerns in `backend/internal/adapters`.
- Keep concrete infrastructure concerns in `backend/internal/infrastructure`.
- Wire dependencies in `backend/cmd/server`.

## Runtime Rules

- Keep runtime configuration environment-driven unless the task explicitly changes configuration strategy.
- Treat default paths, addresses, and startup behaviors as implementation details unless the task requires changing runtime config behavior.
- If a change affects startup, shutdown, migrations, or operational endpoints, verify the impact on existing runtime behavior before modifying it.

## Testing And Validation

- Prefer targeted Go tests for the packages you change.
- When changing HTTP behavior, test both success and failure paths where practical.
- When changing migrations, validate the schema snapshot and any generated artifacts affected by the change.
- Report the exact commands you ran and any checks you could not run.

## Security

- Default to least privilege for filesystem, processes, and external services.
- Do not introduce secrets into source, fixtures, or tests.
- Be explicit about file-path and database-path assumptions before destructive or stateful operations.

## Standard Commands

- Run backend: `go run ./backend/cmd/server`
- List project modules: `go list -m all`
- Inspect one module version: `go list -m <module>`
- Read package docs: `go doc <package>`
- Read symbol docs: `go doc <package>.<symbol>`
- Run tests: `go test ./...`

## Done Criteria

- Changes respect the backend layering rules.
- Generated artifacts are updated when their sources change.
- The smallest relevant validation has been run.
- Any unresolved risk, skipped validation, or assumption is called out clearly.
