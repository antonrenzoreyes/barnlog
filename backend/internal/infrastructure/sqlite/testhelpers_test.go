package sqlite

import (
	"database/sql"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"barnlog/backend/internal/infrastructure/sqlite/sqlc"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func newTestAnimalWriteStore(t *testing.T) (animalWriteStore, *sql.DB) {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "test.sqlite3")
	applyTestMigrations(t, dbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}

	return animalWriteStore{
		queries:  sqlc.New(db),
		photoDir: t.TempDir(),
		now:      time.Now,
	}, db
}

func applyTestMigrations(t *testing.T, dbPath string) {
	t.Helper()

	migrationsPath := testMigrationsPath(t)
	dbURL := (&url.URL{Scheme: "sqlite", Path: dbPath}).String()
	srcURL := (&url.URL{Scheme: "file", Path: migrationsPath}).String()

	m, err := gomigrate.New(srcURL, dbURL)
	if err != nil {
		t.Fatalf("initialize migrate: %v", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			t.Fatalf("close migration source: %v", srcErr)
		}
		if dbErr != nil {
			t.Fatalf("close migration db: %v", dbErr)
		}
	}()

	if err := m.Up(); err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		t.Fatalf("run migrations: %v", err)
	}
}

func testMigrationsPath(t *testing.T) string {
	t.Helper()

	if path := os.Getenv("BARNLOG_TEST_MIGRATIONS_PATH"); path != "" {
		if filepath.IsAbs(path) {
			return path
		}
		return filepath.Clean(filepath.Join(testRepoRoot(t), path))
	}

	return filepath.Join(testRepoRoot(t), "backend", "db", "migrations")
}

func testRepoRoot(t *testing.T) string {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("resolve caller path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", "..", ".."))
}
