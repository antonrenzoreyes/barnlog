#!/usr/bin/env bash
set -euo pipefail

MIGRATIONS_DIR="${MIGRATIONS_DIR:-backend/db/migrations}"
SCHEMA_OUT="${SCHEMA_OUT:-backend/db/schema.sql}"

if ! command -v migrate >/dev/null 2>&1; then
  echo "migrate CLI not found. Install with: go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.0"
  exit 1
fi

if ! command -v sqlite3 >/dev/null 2>&1; then
  echo "sqlite3 CLI not found. Install sqlite3 and retry."
  exit 1
fi

tmp_dir="$(mktemp -d)"
trap 'rm -rf "${tmp_dir}"' EXIT

db_file="${tmp_dir}/schema.sqlite3"
db_url="sqlite3://${db_file}"

migrate -path "${MIGRATIONS_DIR}" -database "${db_url}" up >/dev/null

mkdir -p "$(dirname "${SCHEMA_OUT}")"

sqlite3 "${db_file}" <<'SQL' > "${SCHEMA_OUT}"
SELECT sql || ';'
FROM sqlite_master
WHERE sql IS NOT NULL
  AND type IN ('table', 'index', 'trigger', 'view')
  AND name NOT LIKE 'sqlite_%'
  AND name != 'schema_migrations'
ORDER BY
  CASE type
    WHEN 'table' THEN 1
    WHEN 'index' THEN 2
    WHEN 'trigger' THEN 3
    WHEN 'view' THEN 4
    ELSE 5
  END,
  name;
SQL

echo "Generated ${SCHEMA_OUT} from ${MIGRATIONS_DIR}"
