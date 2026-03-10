#!/usr/bin/env bash
set -euo pipefail

SKILL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MANIFEST="${SKILL_DIR}/references/skills-manifest.txt"
INSTALLER="/Users/areyes/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py"

if [[ ! -f "${MANIFEST}" ]]; then
  echo "Manifest not found: ${MANIFEST}" >&2
  exit 1
fi

if [[ ! -f "${INSTALLER}" ]]; then
  echo "Installer not found: ${INSTALLER}" >&2
  exit 1
fi

while IFS= read -r line; do
  [[ -z "${line}" ]] && continue
  [[ "${line}" =~ ^# ]] && continue

  repo="$(awk '{print $1}' <<<"${line}")"
  skill_path="$(awk '{print $2}' <<<"${line}")"

  if [[ -z "${repo}" || -z "${skill_path}" ]]; then
    echo "Invalid manifest line: ${line}" >&2
    exit 1
  fi

  skill_name="$(basename "${skill_path}")"
  rm -rf "/Users/areyes/.codex/skills/${skill_name}"

  python3 "${INSTALLER}" \
    --repo "${repo}" \
    --path "${skill_path}" \
    --method git

done < "${MANIFEST}"

echo "Skill update complete."
