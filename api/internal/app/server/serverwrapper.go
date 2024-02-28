package server

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"social-network/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL, config.DatabaseSchema)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	srv.logger.Printf("The server is running on the port %v", config.Port)

	return http.ListenAndServe(config.Port, srv)
}

func newDB(databaseURL, dataBaseSchema string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	sqlStmt, err := os.ReadFile(dataBaseSchema)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(string(sqlStmt))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
