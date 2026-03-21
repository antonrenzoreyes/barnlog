# Event Store Agent Guide

## Scope

- This file applies to all work under `backend/internal/domain/events/`.

## Source Of Truth

- Use `EVENTS_TABLE.md` as the source of truth when changing event-store persistence, event queries, or event table behavior.
- Keep event-store changes aligned with the append-only model unless the task explicitly changes that design.

## Validation

- Call out any behavioral or schema change that affects event ordering, event immutability, or query assumptions.
