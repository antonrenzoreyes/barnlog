.PHONY: db-schema db-schema-check api-docs openapi-generate-types openapi-check install-hooks

db-schema:
	./scripts/generate_schema_snapshot.sh

db-schema-check: db-schema
	git diff --exit-code -- backend/db/schema.sql

api-docs:
	go run ./backend/cmd/openapi-export -in backend/openapi/openapi.yaml -json backend/docs/swagger.json -yaml backend/docs/swagger.yaml

openapi-generate-types:
	go tool oapi-codegen --config backend/openapi/oapi-codegen.yaml backend/openapi/openapi.yaml
	go tool oapi-codegen --config backend/openapi/oapi-codegen-server.yaml backend/openapi/openapi.yaml

openapi-check:
	$(MAKE) api-docs
	$(MAKE) openapi-generate-types
	git diff --exit-code -- backend/docs/swagger.json backend/docs/swagger.yaml backend/internal/contracts/openapi/types.gen.go backend/internal/contracts/openapi/server.gen.go

install-hooks:
	mkdir -p .git/hooks
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
