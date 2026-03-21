# Frontend Agent Guide

## Scope

- This file applies to all work under `frontend/`.

## Skills

- Use `$critique` to evaluate design effectiveness from a UX perspective.
- Use `$frontend-design` to build distinctive, production-grade frontend interfaces that avoid generic AI aesthetics.
- Use `$svelte-code-writer` whenever creating, editing, or analyzing any Svelte component (`.svelte`) or module (`.svelte.ts`/`.svelte.js`).
- Use `$svelte-core-bestpractices` to enforce modern Svelte reactivity, event handling, styling, and integration patterns.

## Stack Baseline

- Framework: SvelteKit 2 with Svelte 5.
- Language: TypeScript for app and tests.
- Build tooling: Vite.
- Package manager/runtime: Bun (use `bun run <script>` for project scripts).

## Implementation Rules

- Write new Svelte code using Svelte 5 patterns (runes where applicable).
- Prefer strongly typed code; avoid `any` unless unavoidable and documented inline.
- Prefer simple implementations over abstractions unless logic is reused 2+ times.
- Keep components focused and move reusable logic to modules under `src/lib`.
- Follow existing formatting/linting config (Oxc + ESLint for Svelte rules).

## Testing Requirements

- Unit/component tests: use Vitest.
- End-to-end tests: use Playwright.
- For UI changes, add at least 1 happy-path unit/component test or 1 E2E check.
- For API-backed UI changes, also add at least 1 non-happy-path test (error or loading state assertion).

## MCP And Tooling

- For Svelte or SvelteKit documentation and framework-specific guidance, use the `svelte` MCP server.
- For browser automation, UI flow checks, screenshots, and Playwright-driven E2E work, use `playwright-cli`.

## Architecture

- Frontend is CSR-only.
- Do not add SvelteKit server routes/endpoints for business logic.
- All domain and persistence logic lives in the Go backend.

## Environment

- Configure backend base URL with public env vars only (for example, `PUBLIC_API_BASE_URL`).
- Never store secrets in frontend env or source.
- This project does not implement user authentication.

## CI Gates

- Before merge, run `bun run check`, `bun run lint`, and all tests (`bun run test`).

## Test Locations

- Unit/component tests:
  - `src/**/*.test.ts`
  - `src/**/*.spec.ts`
- E2E tests:
  - `e2e/**/*.spec.ts`

## Standard Commands

- Dev server: `bun run dev`
- Type and Svelte checks: `bun run check`
- Lint: `bun run lint`
- Generate OpenAPI types: `bun run gen:api-types`
- Verify OpenAPI generated types are up to date: `bun run check:api-types`
- Unit tests: `bun run test:unit -- --run`
- E2E tests: `bun run test:e2e`
- Full test pass: `bun run test`

## Done Criteria

- Build and checks pass for changed code.
- Changed UI is covered by at least 1 happy-path unit/component test or 1 E2E check.
- API-backed UI changes include at least 1 non-happy-path test (error or loading state assertion).
