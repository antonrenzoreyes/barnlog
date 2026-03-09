# UI Components Agent Guide

## Scope

- This file applies to all work under `frontend/src/lib/ui/components`.
- Always consult the svelte MCP server when implementing changes under this directory.

## Component Conventions

- Use Svelte 5 runes and typed `$props`.
- Prefer concise prop interfaces (`*Props`) in each component file.
- Prefix shared UI primitives with `B` (`BButton`, `BCard`, etc.).
- Use kebab-case filenames for components/fixtures/specs (for example: `b-field.svelte`).
- Use a component root class with `b-` naming (for example: `.b-button`).
- Keep implementation patterns consistent across components in this directory; if you intentionally diverge, state the reason explicitly.
- Keep defaults explicit in prop destructuring.

## Props And API Shape

- Wrapper components with direct HTML counterparts should extend `svelte/elements` attributes.
  - Example: button-like components extend `HTMLButtonAttributes`.
- Non-wrapper primitives should use a closed API (declare only supported props).
- Do not add broad passthrough props unless there is a concrete consumer need.

## Styling

- Keep component-specific styles in the component `<style>` block.
- Move shared, cross-component utility styles to `src/routes/layout.css`.
- Avoid inline style attributes in component markup.

## Documentation

- Add concise JSDoc above each `*Props` interface (purpose + defaults/important behavior).
- Keep `/dev/components` visual-first; avoid long contract prose there.
- Component files are the source of truth for API contracts.
- When adding a component, update `src/lib/ui/components/index.ts`, `src/lib/ui/index.ts`, and `/dev/components`.

## Testing

- Add/update a fixture and a spec for component behavior changes:
  - `*.fixture.svelte`
  - `*.svelte.spec.ts`
- Prefer assertions on user-visible behavior (role/text/state/classes), not implementation internals.
- Prefer semantic HTML and accessible labels/roles for interactive UI.

## Validation Before Handoff

- Run `bun run check`.
- Run `bun run lint`.
- Run the smallest relevant unit test command for changed components.
