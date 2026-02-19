# EPIC-01: Simple Barn Logging

## Outcome
Deliver a minimal, reliable online barn logging workflow for core day-to-day records.

## Problem Statement
Barn operators need a simple way to capture day-to-day records (starting with animals) with a fast, dependable online workflow.

## Scope
### In Scope
- Core record capture for barn activity.
- Online animal record creation and listing through backend APIs and UI.
- Optional animal photo upload and linkage by `photo_id`.
- Basic validation for required fields before records are accepted.

### Out of Scope
- Offline data entry and local persistence.
- Deferred/offline sync and reconciliation flows.
- Advanced conflict-resolution UI.
- Full analytics/reporting dashboards.
- Bulk import/export workflows.
- Multi-barn role/permission administration.

## Constraints
- System requires network connectivity for record operations in this phase.

## Success Criteria
- A user can create animals through the online workflow.
- Required field validation prevents invalid animal records from being persisted.
- Newly created animals are visible in the list after successful save.
- Optional animal photos can be uploaded and linked during animal creation.

## Linked User Stories
- `US-01` - Add Animals (`docs/user-stories/US-01.md`)
- `US-02` - List Animals (`docs/user-stories/US-02.md`)
- `US-03` - Home Page (`docs/user-stories/US-03.md`)

## Risks and Open Questions
- Offline-first requirements will need a future epic to introduce local persistence and sync safely.
