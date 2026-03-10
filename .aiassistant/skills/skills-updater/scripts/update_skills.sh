#!/usr/bin/env bash
set -euo pipefail

SKILL_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MANIFEST="${SKILL_DIR}/references/skills-manifest.txt"
CODEX_HOME="${CODEX_HOME:-${HOME}/.codex}"
TARGET_SKILLS_DIR="${CODEX_SKILLS_DIR:-${CODEX_HOME}/skills}"
INSTALLER="${SKILL_INSTALLER_PATH:-${CODEX_HOME}/skills/.system/skill-installer/scripts/install-skill-from-github.py}"

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
  target_dir="${TARGET_SKILLS_DIR}/${skill_name}"
  temp_root="$(mktemp -d "${TMPDIR:-/tmp}/skills-updater.${skill_name}.XXXXXX")"
  backup_dir=""

  python3 "${INSTALLER}" \
    --repo "${repo}" \
    --path "${skill_path}" \
    --method git \
    --dest "${temp_root}"

  if [[ ! -d "${temp_root}/${skill_name}" ]]; then
    echo "Installed skill not found in temp dir: ${temp_root}/${skill_name}" >&2
    rm -rf "${temp_root}"
    exit 1
  fi

  if [[ -e "${target_dir}" ]]; then
    backup_dir="${target_dir}.bak.$(date +%s)"
    mv "${target_dir}" "${backup_dir}"
  fi

  if mv "${temp_root}/${skill_name}" "${target_dir}"; then
    rm -rf "${temp_root}"
    if [[ -n "${backup_dir}" && -e "${backup_dir}" ]]; then
      rm -rf "${backup_dir}"
    fi
  else
    rm -rf "${temp_root}"
    if [[ -n "${backup_dir}" && -e "${backup_dir}" ]]; then
      mv "${backup_dir}" "${target_dir}"
    fi
    echo "Failed to replace ${target_dir}" >&2
    exit 1
  fi

done < "${MANIFEST}"

echo "Skill update complete."
