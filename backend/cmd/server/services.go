package main

import (
	"database/sql"

	"barnlog/backend/internal/application"
	"barnlog/backend/internal/infrastructure/config"
	sqliteinfra "barnlog/backend/internal/infrastructure/sqlite"
)

// Services groups application services wired at process startup.
type Services struct {
	AnimalWriter application.AnimalWriter
}

func newServices(cfg config.Config, db *sql.DB) Services {
	store := sqliteinfra.NewAnimalWriteStore(db, cfg.FileDir)
	return Services{
		AnimalWriter: application.NewCreateAnimalWriter(store),
	}
}
