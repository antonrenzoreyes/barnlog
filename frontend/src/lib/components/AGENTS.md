# Components Agent Guide

## Scope

- This file applies to all files under `src/lib/components/`.

## Inheritance

- Follow `frontend/AGENTS.md` for shared project rules (stack baseline, testing requirements, tooling, CI gates, commands, and skill workflow).
- This file only defines component-specific standards to avoid duplication.

## Component-Specific Standards

### Design System Alignment

- Keep the Barn visual language consistent: warm neutrals, restrained green accents, practical/utilitarian tone.
- Use design tokens from `src/routes/layout.css` for colors, shadows, and semantic states.
- Prefer canonical Tailwind token syntax (for example, `text-(--ui-color-text-muted)`).
- Avoid hardcoded hex values in component files.

### API Shape

- Add customization props only when there is a concrete, active use case.
- Use `B*` naming for Barn-owned wrapper components.

### Accessibility

- Keep explicit label-control linking (`for` + `id`).
- Preserve visible keyboard focus affordances.

### Bits UI Wrapping

- Use Bits UI for accessibility-heavy behavior primitives (for example, labels/date pickers).
- Wrap library primitives with Barn components for consistent API and styling.
- Keep library internals encapsulated; expose Barn-level props only.
