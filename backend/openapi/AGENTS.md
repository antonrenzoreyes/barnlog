# Backend OpenAPI Agent Guide

## Scope

- This file applies to all work under `backend/openapi/`.

## Canonical Contract

- Treat [`openapi.yaml`](./openapi.yaml) as the canonical API contract document.
- Generate docs artifacts from this file with `make api-docs`.
- Prefer spec edits here over endpoint-specific normalization patches when contract behavior is intentional.

## oapi-codegen

- Generate backend contract models with `make openapi-generate-types`.
- Generated files are written to:
  - `backend/internal/contracts/openapi/types.gen.go`
  - `backend/internal/contracts/openapi/server.gen.go`
- Do not hand-edit generated files.
