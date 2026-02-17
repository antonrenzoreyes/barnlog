---
apply: always
---

# Git Standards

## Required Workflow
- Use Git MCP tools for all git operations (branching, committing, rebasing, pushing, PR preparation).
- Do not use terminal `git` commands unless Git MCP cannot perform the required operation.
- If terminal `git` is required, explicitly state why before proceeding.

## Commit Message Format
- Every commit message must start with: `Issue #<issue_number> `.
- Example: `Issue #9 ci(go): remove push-to-master trigger from Go CI`.
- This requirement applies to normal commits, fixup commits, and amended commits.

## History Hygiene
- Keep commits scoped to a single concern when possible.
- When integrating review fixes, prefer fixup + autosquash into the relevant existing commit.
- Do not rewrite shared history unless explicitly requested or required for branch hygiene.
