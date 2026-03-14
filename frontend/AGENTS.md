# Frontend Agent Guide

## Scope

- This file applies to all work under `frontend/`.

## Stack Baseline

- Framework: SvelteKit 2 with Svelte 5.
- Language: TypeScript for app and tests.
- Build tooling: Vite.
- Package manager/runtime: pnpm + Vite+ (`vp` for scripts/tasks).

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

## API Integration

- Treat the Go service as the single source of truth for domain data.
- Keep API calls in typed client modules under `src/lib/api`.
- Default to OpenAPI-generated types for request/response contracts; document any handwritten type exceptions inline.

## Environment

- Configure backend base URL with public env vars only (for example, `PUBLIC_API_BASE_URL`).
- Never store secrets in frontend env or source.
- This project does not implement user authentication.

## Error Handling

- Handle network errors, timeouts, and non-2xx responses consistently in API client utilities.
- As API utilities are added, normalize errors to one shared app-level format in `src/lib/api` and have pages/components consume that normalized form only.
- Surface clear user-facing loading and error states in pages/components.

## CI Gates

- Before merge, run `vp check`, `vp run lint:svelte`, and all tests (`vp run test`).

## Test Locations

- Unit/component tests:
  - `src/**/*.test.ts`
  - `src/**/*.spec.ts`
- E2E tests:
  - `e2e/**/*.spec.ts`

## Project-Specific Commands

- Generate OpenAPI types: `vp run gen:api-types`
- Verify OpenAPI generated types are up to date: `vp run check:api-types`
- Run E2E tests only: `vp run test:e2e`

## Done Criteria

- Build and checks pass for changed code.
- Changed UI is covered by at least 1 happy-path unit/component test or 1 E2E check.
- API-backed UI changes include at least 1 non-happy-path test (error or loading state assertion).

<!--VITE PLUS START-->

# Using Vite+, the Unified Toolchain for the Web

This project is using Vite+, a unified toolchain built on top of Vite, Rolldown, Vitest, tsdown, Oxlint, Oxfmt, and Vite Task. Vite+ wraps runtime management, package management, and frontend tooling in a single global CLI called `vp`. Vite+ is distinct from Vite, but it invokes Vite through `vp dev` and `vp build`.

## Vite+ Workflow

`vp` is a global binary that handles the full development lifecycle. Run `vp help` to print a list of commands and `vp <command> --help` for information about a specific command.

### Start

- create - Create a new project from a template
- migrate - Migrate an existing project to Vite+
- config - Configure hooks and agent integration
- staged - Run linters on staged files
- install (`i`) - Install dependencies
- env - Manage Node.js versions

### Develop

- dev - Run the development server
- check - Run format, lint, and TypeScript type checks
- lint - Lint code
- fmt - Format code
- test - Run tests

### Execute

- run - Run monorepo tasks
- exec - Execute a command from local `node_modules/.bin`
- dlx - Execute a package binary without installing it as a dependency
- cache - Manage the task cache

### Build

- build - Build for production
- pack - Build libraries
- preview - Preview production build

### Manage Dependencies

Vite+ automatically detects and wraps the underlying package manager such as pnpm, npm, or Yarn through the `packageManager` field in `package.json` or package manager-specific lockfiles.

- add - Add packages to dependencies
- remove (`rm`, `un`, `uninstall`) - Remove packages from dependencies
- update (`up`) - Update packages to latest versions
- dedupe - Deduplicate dependencies
- outdated - Check for outdated packages
- list (`ls`) - List installed packages
- why (`explain`) - Show why a package is installed
- info (`view`, `show`) - View package information from the registry
- link (`ln`) / unlink - Manage local package links
- pm - Forward a command to the package manager

### Maintain

- upgrade - Update `vp` itself to the latest version

These commands map to their corresponding tools. For example, `vp dev --port 3000` runs Vite's dev server and works the same as Vite. `vp test` runs JavaScript tests through the bundled Vitest. The version of all tools can be checked using `vp --version`. This is useful when researching documentation, features, and bugs.

## Common Pitfalls

- **Using the package manager directly:** Do not use pnpm, npm, or Yarn directly. Vite+ can handle all package manager operations.
- **Always use Vite commands to run tools:** Don't attempt to run `vp vitest` or `vp oxlint`. They do not exist. Use `vp test` and `vp lint` instead.
- **Running scripts:** Vite+ commands take precedence over `package.json` scripts. If there is a `test` script defined in `scripts` that conflicts with the built-in `vp test` command, run it using `vp run test`.
- **Do not install Vitest, Oxlint, Oxfmt, or tsdown directly:** Vite+ wraps these tools. They must not be installed directly. You cannot upgrade these tools by installing their latest versions. Always use Vite+ commands.
- **Use Vite+ wrappers for one-off binaries:** Use `vp dlx` instead of package-manager-specific `dlx`/`npx` commands.
- **Import JavaScript modules from `vite-plus`:** Instead of importing from `vite` or `vitest`, all modules should be imported from the project's `vite-plus` dependency. For example, `import { defineConfig } from 'vite-plus';` or `import { expect, test, vi } from 'vite-plus/test';`. You must not install `vitest` to import test utilities.
- **Type-Aware Linting:** There is no need to install `oxlint-tsgolint`, `vp lint --type-aware` works out of the box.

## Review Checklist for Agents

- [ ] Run `vp install` after pulling remote changes and before getting started.
- [ ] Run `vp check` and `vp test` to validate changes.
<!--VITE PLUS END-->
