# Backend Docs Agent Guide

## Scope

- This file applies to all work under `backend/docs/`.

## OpenAPI Artifacts

- Treat the checked-in files in `backend/docs/` as generated artifacts unless the task explicitly requires manual edits.
- Keep `docs.go`, `swagger.json`, and `swagger.yaml` aligned with their source definitions and handler behavior.
- If generated docs change, mention what source change caused the regeneration.

## Validation

- When touching API contracts or generated docs, verify the affected artifacts are still in sync.
