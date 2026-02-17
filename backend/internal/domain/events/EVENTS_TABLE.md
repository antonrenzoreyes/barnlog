# EVENTS_TABLE.md

This document is the feature-local reference for the `events` table semantics and usage.
Use it when implementing event writes, idempotency checks, event queries, and replay logic.

## Source of Truth

- Canonical schema evolution: `backend/db/migrations/000001_init.up.sql`
- Generated snapshot: `backend/db/schema.sql`

If the table meaning changes, update migration/schema/docs together in the same PR.

## Table Purpose

`events` is an append-only event log.
It is the current backend source of truth for state changes.

## Columns

- `id` (`TEXT PRIMARY KEY`): unique event ID.
- `aggregate_type` (`TEXT NOT NULL`): aggregate category (example: `animal`).
- `aggregate_id` (`TEXT NOT NULL`): specific aggregate instance ID.
- `event_type` (`TEXT NOT NULL`): semantic event name.
- `created_by` (`TEXT NOT NULL`): actor/user that initiated the event.
- `source` (`TEXT NOT NULL`): producer channel/system.
- `request_id` (`TEXT NOT NULL`): idempotency key for request retries.
- `event_version` (`INTEGER NOT NULL DEFAULT 1`): payload schema/event contract version.
- `payload_json` (`TEXT NOT NULL`): event data payload.
- `metadata_json` (`TEXT`): optional trace/context metadata.
- `occurred_at` (`TEXT NOT NULL`): business event timestamp.
- `created_at` (`TEXT NOT NULL DEFAULT datetime('now')`): persistence timestamp.

## Constraints and Indexes

- Non-empty checks on key routing/idempotency fields.
- `UNIQUE (source, request_id)` for idempotent writes.
- Index `(aggregate_type, aggregate_id, occurred_at)` for aggregate stream reads/replay.
- Index `(event_type, occurred_at)` for event-type timeline queries.

## Write Rules

- Inserts are append-only. Do not update/delete event rows in application logic.
- Always set `source` + `request_id` from inbound command context.
- On unique conflict (`source`, `request_id`), treat as idempotent retry behavior.

## Read Rules

- Aggregate replay: filter by `aggregate_type`, `aggregate_id`, order by `occurred_at`.
- Analytics/timeline: filter by `event_type`, `occurred_at` window.

## Future Expansion

Derived projection tables can be added later when read patterns require faster current-state queries.
