package model

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //lint:ignore blank-imports
	_ "github.com/golang-migrate/migrate/v4/source/file"       //lint:ignore blank-imports
	"github.com/pkg/errors"
	"go-rest-template-app/config"
)

// MigrateDB performs database migrations for a GORM DB instance.
func MigrateDB(dbCfg *config.Database) error {
	m, err := migrate.New("file://migration", dbCfg.URL)
	if err != nil {
		return errors.Wrap(err, "Failed to create migration instance")
	}

	// Apply all pending migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "Failed to apply migrations")
	}

	return nil
}
