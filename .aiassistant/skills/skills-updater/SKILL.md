---
name: skills-updater
description: Update selected Codex skills from GitHub into ~/.codex/skills using the official skill-installer script. Use when the user asks to refresh installed skills, sync tracked skill versions, or maintain a small list of skills that should be updated together.
---

# Skills Updater

Use this skill to refresh a tracked set of skills from GitHub.

## Tracked Skills

- Source list: `references/skills-manifest.txt`
- Current entries:
- `sveltejs/ai-tools tools/skills/svelte-code-writer`
- `sveltejs/ai-tools tools/skills/svelte-core-bestpractices`

## Update Workflow

1. Review the manifest before updating:
- `cat .aiassistant/skills/skills-updater/references/skills-manifest.txt`

2. Run the updater script:
- `.aiassistant/skills/skills-updater/scripts/update_skills.sh`

3. Report what was updated:
- List skill names updated.
- Mention any failures and the failing entry.

4. Remind the user:
- `Restart Codex to pick up updated skills.`

## Guardrails

- Only update skills listed in `references/skills-manifest.txt` unless the user explicitly asks to modify the list.
- Use `--method git` to avoid HTTPS certificate issues seen with zip download flows.
- Do not silently skip failed entries.
