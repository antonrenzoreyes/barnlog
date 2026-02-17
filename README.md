# Barn Log

Barn Log is a Go-first monorepo for a local-first, event-driven farm logging application.

Current repository status is bootstrap-stage: a minimal Go app exists, with project stack and standards defined in `.aiassistant/rules/`.

## Tech Stack

- Backend: Go 1.26.x, `net/http` + `chi`, SQLite, `sqlc`, `slog`
- Frontend: Svelte + TypeScript, Vite, Tailwind CSS, IndexedDB (Dexie)
- Architecture: pragmatic event-driven model (append-only events and derived state)

See full stack decisions in `.aiassistant/rules/STACK.md`.

## Standards

- Go architecture and coding rules: `.aiassistant/rules/GO_STANDARDS.md`
- MCP usage and security standards: `.aiassistant/rules/MCP_STANDARDS.md`

## Prerequisites

- Go 1.26+

## Run

```bash
go run ./backend/cmd/server
```

## Repository Layout (Current)

```text
barnlog/
├─ go.mod
├─ backend/
│  ├─ cmd/
│  │  └─ server/
│  └─ internal/
│     ├─ domain/
│     ├─ application/
│     ├─ ports/
│     ├─ adapters/
│     └─ infrastructure/
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
