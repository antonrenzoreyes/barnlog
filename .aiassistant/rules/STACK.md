---
apply: always
---

# Barn Log - Finalized Technical Stack

This document defines the finalized, locked-in technology stack and architectural decisions for the Barn Log project.

The goal is clarity, simplicity, and long-term maintainability, not trend-chasing.

---

# Project Goals

- Backend-first development
- Event-driven (without overengineering)
- Local-first deployment
- Offline-capable frontend
- Explicit architecture (no hidden magic)
- Production-credible patterns
- Low operational overhead

---

# Backend

## Language

Go 1.26.x

Why:
- Native HTTP server in stdlib
- Excellent concurrency model
- Low memory footprint
- Explicit and predictable behavior
- Strong industry adoption

## HTTP Layer

net/http + chi

- net/http is the foundation (stdlib)
- chi provides lightweight routing
- Middleware is explicit
- No framework lock-in
- No annotation system
- No hidden lifecycle

## Database

SQLite

- Zero setup
- File-based
- ACID compliant
- Perfect for local-first development
- Same SQL model as Postgres

## Query Layer

sqlc

- Write real SQL
- Generate typed Go code
- Compile-time query validation
- No ORM magic
- Excellent for an event-store pattern

## Logging

slog

- Structured logging
- Standard library
- Minimal overhead

## Concurrency

- goroutines
- context
- errgroup

No job frameworks. No message brokers. No Kafka.

---

# Frontend

## Framework

Svelte + TypeScript

- Minimal boilerplate
- Reactive without hooks complexity
- Clean mental model
- Suitable for event timeline UI
- Lower cognitive overhead than React

## Build Tool

Vite

- Fast dev server
- Clean production build
- Minimal abstraction

## Styling

Tailwind CSS

- Utility-first styling
- Fast iteration
- Mobile-friendly

## Offline Storage

IndexedDB (Dexie by default)

- Local-first capability
- Device-level persistence
- Sync-friendly

---

# Architecture Model

Pragmatic event-driven design:

Devices -> HTTP API -> events (append-only table) -> derived_state tables

Principles:
- Events are the source of truth
- State is derived from events
- No heavy CQRS
- No distributed complexity
- No microservices

---

# Repository Shape

Monorepo with separate frontend and backend directories.

- Single repository
- Clear boundary between frontend and backend
- Simple CI/CD and release management

---

# Implementation Standards

Go architecture, layering, testing, and coding rules are defined in:

`.aiassistant/rules/GO_STANDARDS.md`

---

# Final Summary

Backend: Go + net/http + chi + SQLite + sqlc + slog

Frontend: Svelte + Vite + Tailwind + IndexedDB

Architecture: Event-driven, append-only events, derived state

This stack is intentional, minimal, and scalable.
