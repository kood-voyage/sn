package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"social-network/internal/app/config"
	"social-network/internal/store/sqlstore"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

// @title social-network API
// @version 1.0
// @description API Server for social-network project

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func Start(config *config.Config) error {
	if os.Getenv("JWT_KEY") == "" {
		log.Fatal("JWT_KEY not imported")
	}
	db, err := newDB(config.DatabaseURL, config.Migrations, config.Driver)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)

	if os.Getenv("JWT_KEY") == "" {
		return errors.New("JWT KEY not set")
	}

	srv := newServer(store, WithConfig(config))

	srv.logger.Printf("The server is running on the port %v", config.Port)

	return http.ListenAndServe(config.Port, srv)
}

func WithConfig(cfg *config.Config) Option {
	return func(c *config.Config) {
		*c = *cfg
	}
}

func newDB(databaseURL, migrationSource, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("instance: %w", err)
	}

	fileSource, err := (&file.File{}).Open(migrationSource)
	if err != nil {
		return nil, fmt.Errorf("fileSource: %w", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, driver, instance)
	if err != nil {
		return nil, fmt.Errorf("migrations new: %w", err)
	}

	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, fmt.Errorf("migrations run: %w", err)
		}
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
