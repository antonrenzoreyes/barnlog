# Prototype Standards (Simple)

## Goal
Keep documentation lightweight and consistent.

## Required Outputs
- Screenshots in: `output/playwright/prototype-state/`
- One master file: `output/playwright/prototype-state/MASTER.md`

## Screenshot Rules
- Device: iPad portrait `768x1024`
- Capture product views only (exclude reference pages like `components.html` unless requested).
- For dynamic screens, capture each meaningful state as a separate screenshot.
- Use clear, stable filenames:
  - Default state: `<page>.png`
  - State variant: `<page>-<state>.png`

## Master File Rules
`MASTER.md` must include, for each screenshot:
- filename
- route/page
- brief description (1-2 lines)

Recommended format: a simple table.

## Annotation Guidance
- Prefer text annotations in `MASTER.md` over drawing directly on images.
- Only annotate images visually when explicitly requested.

## Definition of Done
- Screenshots are updated in `output/playwright/prototype-state/`.
- `MASTER.md` exists and briefly describes every screenshot included.
