---
apply: by file patterns
patterns: *.sql
---

# SQL_MIGRATION_STANDARDS.md
## SQL Migration Standards

These standards define how database schema changes must be authored, validated, and reviewed.

---

## 🎯 Core Principles

1. Migrations are the source of truth for schema evolution.
2. Every schema change must be versioned and reviewable.
3. Keep migrations deterministic and reversible.
4. Prefer small, focused changes over large multipurpose migrations.

---

## 📁 Location and Tooling

- Migration files live in `backend/db/migrations/`.
- Migration engine: `golang-migrate`.
- Migration pair format:
  - `NNNNNN_name.up.sql`
  - `NNNNNN_name.down.sql`
- Current filename convention:
  - 6-digit version prefix + snake_case name (example: `000001_init`).

---

## 🧱 Authoring Rules

- One logical change per migration pair.
- Always create both `up` and `down` files.
- Use explicit SQL only; no hidden generation.
- Keep statements idempotency-friendly where practical (`IF EXISTS`, `IF NOT EXISTS`).
- Avoid data-destructive operations unless required and documented in the PR.

---

## ⬆️ Up Migration Rules

- Create/alter schema required for forward progress.
- Add indexes required by new query paths.
- Keep lock-sensitive operations small.
- If a migration is risky, split it into multiple sequential migrations.

---

## ⬇️ Down Migration Rules

- Reverse `up` steps in safe order.
- Drop dependent objects before referenced objects.
- Ensure a rollback path is executable in CI.

---

## 🧾 Schema Snapshot

- Generated schema snapshot lives at `backend/db/schema.sql`.
- Do not hand-edit `backend/db/schema.sql`.
- Regenerate snapshot after migration changes:
  - `make db-schema`
- Verify the snapshot is current:
  - `make db-schema-check`

---

## ✅ Validation and CI

Migration CI must enforce:
- valid migration pair presence (`.up.sql` + `.down.sql`)
- filename convention compliance
- SQLite up/down smoke execution
- schema snapshot freshness (`backend/db/schema.sql`)

---

## 🔎 Review Checklist

For every migration PR, confirm:
- migration names follow convention
- up/down pair exists and is coherent
- rollback is valid
- indexes are present for new access patterns
- `backend/db/schema.sql` was regenerated

---

## 🚫 Anti-Patterns

- Editing old committed migration files after merge
- Combining unrelated schema concerns in one migration
- Omitting `down` migration
- Treating `schema.sql` as the source of truth
