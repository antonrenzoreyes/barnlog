# Database Agent Guide

## Scope

- This file applies to all work under `backend/db/`.

## Migrations

- Migrations are SQL-first and live in `backend/db/migrations`.
- Migration filenames must follow `000001_name.up.sql` and `000001_name.down.sql`.
- Keep migration changes minimal and reversible unless the task explicitly requires otherwise.

## Schema And Queries

- Keep `backend/db/schema.sql` in sync when migrations change.
- Keep handwritten SQL queries in `backend/db/queries`.
- Prefer updating the SQL source files rather than generated outputs.

## SQLC

- Keep SQLC output in `backend/internal/infrastructure/sqlite/sqlc` aligned with `backend/sqlc.yaml`.
- Do not hand-edit generated SQLC output unless the generation workflow explicitly requires it.

## Validation

- When migrations change, run the smallest relevant schema or generation checks.
- Report whether you updated or verified `backend/db/schema.sql`.

## Standard Commands

- Generate schema snapshot: `make db-schema`
- Verify schema snapshot: `make db-schema-check`
- Generate SQLC code: `sqlc generate -f backend/sqlc.yaml`
