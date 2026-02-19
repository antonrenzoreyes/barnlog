---
apply: by file patterns
patterns: *.go
---

# GO_STANDARDS.md
## Go Architecture & Coding Standards

These standards define how Go services should be structured, written, and evolved.
They prioritize long-term maintainability, clarity, and business-focused design.

---

## 🎯 Core Principles

1. Clarity over cleverness
2. Business logic lives in the domain
3. Infrastructure is replaceable
4. Compile-time safety preferred
5. Design for change, not convenience

---

## 📁 Project Structure (Clean Architecture)

/cmd/service-name/main.go

/internal/domain/
/internal/application/
/internal/ports/
/internal/adapters/
/internal/infrastructure/

/pkg (optional shared libraries)

### Dependency Rules

domain → nothing outside domain  
application → domain + ports  
ports → domain/application contracts only  
adapters/infrastructure → application + ports (implements ports)

Infrastructure must never be imported by domain or application.

---

## 🧠 Domain Layer (DDD Lite)

Domain entities must:
- Use private fields
- Enforce validity via constructors
- Expose behavior through methods
- Prevent invalid states by design

Business rules belong in domain methods — never in handlers.

---

## 📦 Repository Pattern

Interfaces live in domain or application.
Implementations live in infrastructure.

Domain must not contain database models or queries.

---

## 🔀 CQRS

Commands mutate state.
Queries read state.

Use when logic becomes non-trivial.

---

## 🌐 APIs

Public: HTTP + OpenAPI  
Internal: package boundaries and clear interfaces (no gRPC by default in this monolith)

---

## 🔐 Authentication

Never custom auth.
Use JWT/OAuth providers.

---

## 🧪 Testing

Domain → unit tests  
Application → use-case/service tests  
Transport (HTTP handlers) → handler tests  
Infrastructure → integration tests

Avoid mocking databases.

---

## 🧾 DRY With Boundaries

Do not share models across layers.
Always map explicitly.

---

## 📏 Go Conventions

Prefer clarity.
Wrap errors with context.
Pass context.Context first.
Use errors.Is / errors.As for error handling.
Keep interfaces small and define them where consumed.
Return concrete types; accept interfaces when useful.
Enforce formatting and static checks: gofmt, go vet, staticcheck.
Use clear package names (short, lowercase, no stutter).
Keep functions small and focused; avoid hidden side effects.
Use `net.JoinHostPort` for host:port construction (never string formatting).
Use `os.Root` / `os.OpenInRoot` when combining trusted base paths with untrusted file names.
Prefer `runtime.AddCleanup` over `runtime.SetFinalizer` in new code.
Prefer `sync.WaitGroup.Go` over manual `Add`/`Done` boilerplate.
Use `go fix ./...` during Go version upgrades and review resulting diffs.
Use `tool` directives in `go.mod` (Go 1.24+) instead of `tools.go` blank imports.
Use `go doc` (or `go doc -http`) instead of deprecated `go tool doc`.

---

## 🚀 Design Preferences

Small focused types  
Explicit flows  
Replaceable infrastructure

Avoid god services and shared mega-models.

---

## 🤖 AI Generation Rules

Default to Clean Architecture.
Put logic in domain first.
Separate models per layer.
Prefer clarity over shortcuts.
