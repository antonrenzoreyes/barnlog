# UI Components Guide

## Scope

- This file applies to all work under `frontend/src/lib/ui/`.

## Naming

- Prefix all shared UI components with `B` (for example: `BButton`, `BCard`, `BIcon`).
- Keep component symbols in PascalCase and component file names in kebab-case.
- Export UI modules through local barrels:
  - `src/lib/ui/components/index.ts`
  - `src/lib/ui/index.ts`

## API Source Of Truth

- Keep component API contracts in the component files themselves using named `*Props` types.
- Treat README and showcase pages as guides/examples, not API source-of-truth.
- When changing props, update all of the following in the same change:
  - component file type contract (`src/lib/ui/components/*.svelte`)
  - component showcase (`src/routes/dev/components/+page.svelte`)
  - any affected tests

## Styling Conventions

- Keep component-specific base styles in each component file (`src/lib/ui/components/*.svelte`).
- Keep app-wide layout/theme primitives in `src/routes/layout.css`.
- In component markup, prefer semantic class composition over inline utility strings.
