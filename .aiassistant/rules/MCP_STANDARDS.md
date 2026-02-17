---
apply: always
---

# MCP_STANDARDS.md
## MCP Usage Standards

These standards define the approved MCP servers, expected usage, and security rules for this repository.

---

## Approved MCPs

- filesystem
- git
- github
- sqlite
- openapi
- context7
- gopls

Do not add new MCP servers without updating this file.

---

## Usage Rules

### filesystem
- Use for local file reads/writes in this repository.
- Scope access to the repo root only.
- Do not use it for secrets outside the project directory.

### git
- Use for diffs, status, history, and review context.
- Prefer non-destructive workflows.
- Use Git MCP for Git operations by default.
- Use shell `git` only when Git MCP cannot perform the required operation or for quick local diagnostics.

### github
- Use for issues, pull requests, labels, and repository metadata.
- Use read-only mode by default unless a write action is explicitly needed.

### sqlite
- Use for schema inspection and local query validation.
- Target local project database paths only.

### openapi
- Use as an API contract source of truth.
- Prefer schema-based decisions to ad-hoc endpoint assumptions.

### context7
- Use for up-to-date package and framework docs.
- Prefer official docs returned by context7 over blog-style sources.

### gopls
- Use for Go code intelligence (diagnostics, references, symbols).
- Use for static guidance, not as a replacement for tests.

---

## Reliability Standards

- Prefer official/vendor MCP servers where available.
- For community MCP servers, pin versions when stability matters.
- If an MCP fails, continue with local tools and document the fallback in responses.
- Revalidate MCP configuration after major IDE or toolchain upgrades.

---

## Project Conventions

- Use `context7` for dependency and framework documentation checks.
- Use `gopls` for Go-aware edits and diagnostics.
- Use `openapi` for request/response and contract alignment.
- Use `github` for feature-ticket and PR workflow traceability.
- Use `filesystem` + `git` for local implementation and review.

---

## Maintenance

- Review MCP list and permissions monthly.
- Remove unused MCP servers.
- Keep `.ai/mcp/mcp.json` minimal and environment-specific.
