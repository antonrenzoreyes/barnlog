.PHONY: db-schema db-schema-check api-docs install-hooks

db-schema:
	./scripts/generate_schema_snapshot.sh

db-schema-check: db-schema
	git diff --exit-code -- backend/db/schema.sql

api-docs:
	go tool swag init --v3.1 -g main.go -d backend/cmd/server,backend/internal/adapters/httpapi --parseInternal -o backend/docs

install-hooks:
	mkdir -p .git/hooks
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
