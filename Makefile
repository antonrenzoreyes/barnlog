.PHONY: db-schema db-schema-check install-hooks

db-schema:
	./scripts/generate_schema_snapshot.sh

db-schema-check: db-schema
	git diff --exit-code -- backend/db/schema.sql

install-hooks:
	mkdir -p .git/hooks
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
