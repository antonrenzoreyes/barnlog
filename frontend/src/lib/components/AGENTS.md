# Components Agent Guide

## Scope

- This file applies to all files under `src/lib/components/`.

## Component Skills

- Use `$arrange` to improve component layout composition, spacing rhythm, and hierarchy.
- Use `$typeset` to improve typography scale, readability, and text hierarchy.
- Use `$colorize` to strengthen color usage and semantic emphasis.
- Use `$bolder` to increase visual impact when components feel too generic.
- Use `$quieter` to reduce visual intensity when components feel too aggressive.
- Use `$distill` to remove unnecessary complexity and visual clutter.
- Use `$polish` for a final pass on alignment, consistency, and finish quality.
- Use `$normalize` to align component design with system conventions.
- Use `$extract` to convert repeated patterns into reusable components and tokens.
- Use `$harden` to improve resilience for unusual content and interaction edge cases.

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
