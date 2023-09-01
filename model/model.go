package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-rest-template-app/config"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL dialects
)

// Service is the interface of all model service.
type Service interface {
}

type service struct {
}

// New returns a Service instance for operating all model service.
func New(dbCfg *config.Database) (Service, error) {
	_, err := newDB(dbCfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialize db of gorm")
	}

	if err := MigrateDB(dbCfg); err != nil {
		log.Fatalf("Failed to migrate the database: %s\n", err)
	}

	serv := &service{}
	return serv, nil
}

func newDB(dbCfg *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbCfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	db.DB().SetMaxOpenConns(dbCfg.MaxActive)
	db.DB().SetMaxIdleConns(dbCfg.MaxIdle)

	db.LogMode(dbCfg.LogMode)
	return db, nil
}
