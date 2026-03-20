# Barn Log

Barn Log is a Go-first monorepo for a local-first, event-driven farm logging application.

## Tech Stack

- Backend: Go 1.26.x, `net/http` + `chi`, SQLite, `sqlc`, `slog`
- Frontend: Bun, SvelteKit 2 (Svelte 5) + TypeScript, Vite 8, Tailwind CSS 4, IndexedDB (Dexie)
- Frontend testing: Vitest (unit/component) + Playwright (e2e)
- Architecture: pragmatic event-driven model (append-only events first, projections added as needed)

## Project Guides

- General repository workflow: `AGENTS.md`
- Frontend-specific guidance: `frontend/AGENTS.md`
- Backend-specific guidance: `backend/AGENTS.md`
- Events table semantics and usage: `backend/internal/domain/events/EVENTS_TABLE.md`

## AI Skills

This repository uses repo-owned, project-local skills.

- Canonical skill files (including custom skills) live in `.agents/skills/` and are committed to git.
- `skills-lock.json` tracks registry-managed skill sources and hashes.
- `bunx skills` is used to install/update registry skills; custom skills remain managed directly in this repository.

```bash
bunx skills list
bunx skills check
bunx skills update
bunx skills experimental_install
```

## Planning Workflow

Planning artifacts live alongside the codebase:

- Epic docs: `docs/epics/EPIC-XX.md`
- Story docs: `docs/user-stories/US-XX.md`
- GitHub issues track execution with native sub-issues:
  - Epic -> Story
  - Story -> Task

## Prerequisites

- Go 1.26+
- Bun (for frontend install, dev, lint, build, and tests)
- `migrate` CLI (`golang-migrate`) for manual migration commands and schema snapshot generation
- `sqlite3` CLI for schema snapshot generation

## Run

```bash
go run ./backend/cmd/server
```

Frontend app:

```bash
cd frontend
bun install
bun run dev
```

## Backend Runtime

Backend configuration is environment-driven. The server uses `slog`, `chi`, graceful shutdown, SQLite path configuration, and automatic migrations on startup by default.

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
