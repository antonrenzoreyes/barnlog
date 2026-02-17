---
apply: by file patterns
patterns: *_test.go
---

# GO_TEST_STANDARDS.md
## Go Testing Standards

These standards define how tests should be structured and what level of behavior each layer should verify.

---

## 🎯 Testing Principles

1. Test behavior, not implementation details.
2. Keep test setup explicit and local to the test package.
3. Prefer small, deterministic tests over broad end-to-end tests.
4. Use table-driven tests when multiple cases share the same execution flow.

---

## 🧪 Layer Testing Responsibilities

- Domain -> unit tests
- Application -> use-case/service tests
- Transport (HTTP handlers) -> handler tests
- Infrastructure -> integration tests

Avoid mocking databases.

---

## 🌐 HTTP Handler Test Standards

Use `net/http/httptest` for transport tests.

Minimum assertions per endpoint:
- expected HTTP status code
- expected `Content-Type`
- JSON response can be decoded
- endpoint-specific payload assertions

Prefer one table-driven test per route group when endpoints share setup and assertion flow.

Use separate test files for endpoints only when the behavior is substantially different.

---

## 📋 Table-Driven Test Pattern

Use this pattern for repetitive endpoint checks:

1. Define test cases as a slice of structs.
2. Include per-case input and expected output.
3. Use `t.Run` with case name.
4. Keep shared assertions in helpers.

This reduces duplication and keeps route contract checks consistent.

---

## 🧱 Test Organization

- Keep `_test.go` files near the package they test.
- Reuse helpers only when they reduce duplication without hiding intent.
- Name tests with behavior-focused names (`TestRoutes`, `TestWriteJSONMarshalError`).

---

## ⚙️ Tooling Expectations

- Run `go test ./...` locally and in CI.
- Keep tests compatible with race testing (`go test -race ./...`).
- Keep fixtures minimal and explicit.
