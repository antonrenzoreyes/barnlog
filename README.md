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
go run .
```

## Repository Layout (Current)

```text
barnlog/
├─ main.go
├─ go.mod
├─ backend/
│  └─ db/
│     └─ dev.sqlite3
├─ logs/
├─ .ai/
│  └─ mcp/
└─ .aiassistant/
   └─ rules/
```

## Near-Term Direction

- Replace bootstrap `main.go` with backend service entrypoint under `backend/cmd/server`
- Add HTTP routes and domain/application layers per Go standards
- Introduce migrations, queries, and event-store schema
- Add frontend app scaffold and API integration
