---
apply: by file patterns
patterns: backend/db/migrations/*.sql,backend/db/schema.sql,backend/internal/domain/events/*.md,backend/internal/**/*.go
---

# EVENTS_CONTEXT.md
## Events Context Rule

When changing event write/read logic or schema-related backend code, consult:
- `backend/internal/domain/events/EVENTS_TABLE.md`
- 
Requirements:
- Keep event semantics, schema, and docs aligned.
- If event columns/constraints/indexes change, update `EVENTS_TABLE.md` in the same change.
- Preserve idempotency guarantees around `source + request_id`.
