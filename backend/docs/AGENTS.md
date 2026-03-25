# Backend Docs Agent Guide

## Scope

- This file applies to all work under `backend/docs/`.

## OpenAPI Artifacts

- Treat the checked-in files in `backend/docs/` as generated artifacts unless the task explicitly requires manual edits.
- Keep `backend/docs/swagger.json` and `backend/docs/swagger.yaml` generated from `backend/openapi/openapi.yaml` via `make api-docs`.
- If generated docs change, mention what source change caused the regeneration.

## Validation

- When touching API contracts or generated docs, verify the affected artifacts are still in sync.
