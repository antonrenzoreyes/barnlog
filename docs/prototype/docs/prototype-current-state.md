# Current Prototype Documentation

## Snapshot Metadata
- Date: 2026-02-19
- Source URL: `http://127.0.0.1:4173`
- Capture artifacts: `output/playwright/prototype-state`
- Route inventory: `output/playwright/prototype-state/routes-captured.md`
- User stories index: `output/playwright/prototype-state/user-stories/README.md`
- Viewport standard: iPad portrait `768x1024`

## Screen Inventory
- Home: `/index.html`
- Animal Detail: `/animal-detail.html`
- Add Animal: `/add-animal.html`
- Edit Animal: `/edit-animal.html`
- Log Event: `/log-event.html`

## Navigation and Primary Flows
- Home -> Animal Detail:
  - Animal cards on Home link to `animal-detail.html`.
- Home -> Add Animal:
  - `Add animal` action links to `add-animal.html`.
- Animal Detail -> Edit Animal:
  - `Edit animal` action links to `edit-animal.html`.
- Animal Detail -> Log Event:
  - `Add log event` action links to `log-event.html`.
- Form exits:
  - Add/Edit forms support `Cancel` and `Save` links.
  - Log Event supports `Cancel` and `Save Event` links.

## Dynamic UI Behavior
### Log Event (`/log-event.html`)
The form changes based on selected event type:
- Feed: field `Feed Amount (optional)`
- Medication: fields `Medication Type`, `Dosage`
- Weight: field `Weight`
- Note: field `Note`

Variant screenshots:
- `log-event-feed.png`
- `log-event-medication.png`
- `log-event-weight.png`
- `log-event-note.png`

## Screenshots
- `output/playwright/prototype-state/home.png`
- `output/playwright/prototype-state/animal-detail.png`
- `output/playwright/prototype-state/add-animal.png`
- `output/playwright/prototype-state/edit-animal.png`
- `output/playwright/prototype-state/log-event-feed.png`
- `output/playwright/prototype-state/log-event-medication.png`
- `output/playwright/prototype-state/log-event-weight.png`
- `output/playwright/prototype-state/log-event-note.png`

## Observed Console Signals During Capture
- Warning: Tailwind CDN runtime warning (prototype mode)
- Error: `favicon.ico` missing (404)

Console logs are available under:
- `output/playwright/prototype-state/.playwright-cli/console-*.log`
