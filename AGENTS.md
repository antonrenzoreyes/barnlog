# Agent Guide

## Scope

- This file applies to all work in this repository.

## General Rules

- I can be wrong. If something seems off, challenge it with evidence.
- Do not make silent assumptions. If ambiguity could change behavior, ask first; otherwise state your assumption explicitly and proceed.
- Do not create commits unless I explicitly ask you to commit.
- Keep security top-of-mind in every action; default to least privilege and ask before any potentially sensitive or destructive step.
- Validate before handoff: run the smallest relevant checks/tests and report exactly what you ran (and what you could not run).
- Explain your reasoning clearly and concisely, including assumptions and tradeoffs.

## Skills

- Repository-owned skills live in `.agents/skills/` and should be versioned with this repo.
- Use `$barnboard` to create and maintain planning artifacts with docs-plus-issues workflow for Barn Log.
- Use `$barnreview` to review local changes for bugs, regressions, architecture drift, and missing tests.
