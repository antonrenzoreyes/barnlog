# Barn Log

Barn Log is a Go-first monorepo for a local-first, event-driven farm logging application.

Current repository status is bootstrap-stage: a minimal Go app exists, with project stack and standards defined in `.aiassistant/rules/`.

## Tech Stack

- Backend: Go 1.26.x, `net/http` + `chi`, SQLite, `sqlc`, `slog`
- Frontend: Svelte + TypeScript, Vite, Tailwind CSS, IndexedDB (Dexie)
- Architecture: pragmatic event-driven model (append-only events first, projections added as needed)

See full stack decisions in `.aiassistant/rules/STACK.md`.

## Standards

- Go architecture and coding rules: `.aiassistant/rules/GO_STANDARDS.md`
- Go testing standards: `.aiassistant/rules/GO_TEST_STANDARDS.md`
- SQL migration standards: `.aiassistant/rules/SQL_MIGRATION_STANDARDS.md`
- Events context rules: `.aiassistant/rules/EVENTS_CONTEXT.md`
- MCP usage and security standards: `.aiassistant/rules/MCP_STANDARDS.md`

## Planning Workflow

Planning artifacts are managed with the custom `barnboard` skill.

- Epic docs: `docs/epics/EPIC-XX.md`
- Story docs: `docs/user-stories/US-XX.md`
- GitHub issues track execution with native sub-issues:
- Epic -> Story
- Story -> Task

## Data Model Notes

- Events table semantics and usage: `backend/internal/domain/events/EVENTS_TABLE.md`

## Prerequisites

- Go 1.26+
- `migrate` CLI (`golang-migrate`) for manual migration commands and schema snapshot generation
- `sqlite3` CLI for schema snapshot generation

## Run

```bash
go run ./backend/cmd/server
```

## Backend Runtime

Runtime wiring currently includes:

- Environment-driven configuration
- Structured logging via `slog`
- HTTP server with `chi` middleware
- Graceful shutdown on `SIGINT`/`SIGTERM`
- Basic operational endpoints: `/healthz`, `/readyz`
- SQLite database path configuration
- Automatic migration execution on startup

### Environment Variables

- `BARNLOG_ENV` (default: `dev`)
- `BARNLOG_HTTP_ADDR` (default: `:8080`)
- `BARNLOG_DB_PATH` (default: `backend/db/dev.sqlite3`)
- `BARNLOG_MIGRATIONS_PATH` (default: `backend/db/migrations`)
- `BARNLOG_AUTO_MIGRATE` (default: `true`)
- `BARNLOG_LOG_LEVEL` (default: `info`)
- `BARNLOG_SHUTDOWN_TIMEOUT` (default: `10s`)

## Migrations

Migrations are SQL-first and managed with `golang-migrate`.
Application startup runs pending migrations before serving traffic when `BARNLOG_AUTO_MIGRATE=true` (default).

File naming:
- `backend/db/migrations/000001_name.up.sql`
- `backend/db/migrations/000001_name.down.sql`
- Generated snapshot: `backend/db/schema.sql`

Install CLI locally:

```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.0
```

Run migrations manually:

```bash
migrate -path backend/db/migrations -database "sqlite3://backend/db/dev.sqlite3" up
```

Rollback one migration:

```bash
migrate -path backend/db/migrations -database "sqlite3://backend/db/dev.sqlite3" down 1
```

Generate schema snapshot:

```bash
make db-schema
```

Verify schema snapshot is current:

```bash
make db-schema-check
```

Install optional pre-commit hook (auto-regenerates/stages `backend/db/schema.sql` when migrations change):

```bash
make install-hooks
```

## SQLC

SQL queries for typed code generation live in:
- `backend/db/queries/`

SQLC configuration:
- `backend/sqlc.yaml`

Install SQLC CLI:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Generate Go code:

```bash
sqlc generate -f backend/sqlc.yaml
```

Generated package output:
- `backend/internal/infrastructure/sqlite/sqlc`

## Repository Layout (Current)

```text
barnlog/
├─ go.mod
├─ backend/
│  ├─ sqlc.yaml
│  ├─ cmd/
│  │  └─ server/
│  ├─ db/
│  │  ├─ migrations/
│  │  └─ schema.sql
│  └─ internal/
│     ├─ domain/
│     ├─ application/
│     ├─ ports/
│     ├─ adapters/
│     └─ infrastructure/
├─ scripts/
├─ .githooks/
├─ logs/
├─ .ai/
│  └─ mcp/
└─ .aiassistant/
   └─ rules/
```

## Backend Layering Guide

The backend is split to keep business rules stable while transport and storage remain replaceable.

- `backend/cmd/server`: composition root and process startup. Wire dependencies here.
- `backend/internal/domain`: core business concepts, invariants, and behavior.
- `backend/internal/application`: use cases that orchestrate domain behavior.
- `backend/internal/ports`: interfaces consumed by application/domain for external needs.
- `backend/internal/adapters`: transport-facing adapters (HTTP handlers, DTO mapping).
- `backend/internal/infrastructure`: concrete implementations (DB, logging integrations, file IO, external services).

### Dependency Direction

Use this direction only:

- `domain` -> no internal dependency
- `application` -> `domain` + `ports`
- `ports` -> contracts only (no concrete implementation)
- `adapters` -> `application` (+ mapping to transport models)
- `infrastructure` -> `ports`/`application` as implementations

Practical rule: domain and application must not import transport or database packages.

### Request Flow (Example)

1. HTTP request enters an adapter in `backend/internal/adapters`.
2. Adapter validates/parses input and calls an application use case.
3. Application coordinates domain logic and calls interfaces from `ports`.
4. Infrastructure implementations satisfy those ports (for example, SQLite repository).
5. Adapter maps use-case output to HTTP response.

This keeps the core business layer testable without HTTP or database setup.

## Near-Term Direction

- Add HTTP routes and domain/application layers per Go standards
- Introduce migrations, queries, and event-store schema
- Add frontend app scaffold and API integration
