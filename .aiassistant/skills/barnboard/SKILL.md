---
name: barnboard
description: Create and maintain Barn Log planning artifacts using a lightweight docs-plus-issues workflow (epic docs, story docs, and linked Epic/Story/Task GitHub issues).
---

# Barnboard

Use this skill when the user asks to plan, create, or organize feature work through epics, user stories, and tasks.

## Scope

- Maintain docs as a source of truth:
- `docs/epics/EPIC-XX.md`
- `docs/user-stories/US-XX.md`
- Keep GitHub issues lightweight execution trackers:
- Epic issue → epic doc and linked story issues
- Story issue → story doc and linked task issues
- Task issue → concrete technical goal linked to a story

## Tooling

- Use `github` MCP for issue creation, updates, linking, and tracking.
- Use `filesystem` MCP for doc and template edits.
- Use `git` MCP for local change review and commit workflow.
- Use GitHub native sub-issues (via `github` MCP) for Epic → Story and Story → Task hierarchy.

## Workflow

1. Confirm the planning unit requested:
- Epic only
- Story only
- Task only
- Full chain (epic → stories → tasks)

2. Assign IDs before creating artifacts:
- Determine next epic ID by scanning existing epic docs/issues for `EPIC-\d+` and incrementing the highest value.
- Determine next story ID by scanning existing story docs/issues for `US-\d+` and incrementing the highest value.
- Use zero-padded 2-digit format (`EPIC-01`, `US-01`).
- Only ask the user for an ID when no prior IDs exist.

3. Create or update docs first:
- Epic doc captures outcome, scope, and linked story IDs.
- Story doc captures the user story and acceptance criteria.

4. Create or update linked issues from templates:
- `.github/ISSUE_TEMPLATE/epic.md`
- `.github/ISSUE_TEMPLATE/story.md`
- `.github/ISSUE_TEMPLATE/task.md`
- Always read the current template file contents immediately before creating or updating an issue.
- Build issue bodies from the current template structure; do not use hardcoded or remembered template text.
- If a template has changed, the issue body must reflect the latest template version.

5. Enforce minimal linking (one-direction chain):
- Create native sub-issue links:
- Add each Story issue as a sub-issue of its Epic issue.
- Add each Task issue as a sub-issue of its Story issue.
- Do not maintain manual parent-child checklists in issue bodies for hierarchy tracking.

6. Keep issue content minimal:
- Avoid duplicating full acceptance criteria from docs in issues unless explicitly requested.

## Guardrails

- Prefer stable IDs (`EPIC-XX`, `US-XX`) in titles and docs.
- Do not close story/task issues unless done criteria in the issue are checked.
- If docs and issues conflict, treat docs as product truth and update issues.

## Deliverable Checklist

- Required docs exist and are linked.
- The required issue hierarchy exists and is linked.
- No unnecessary cross-linking in both directions.
- Templates remain lightweight and consistent.
