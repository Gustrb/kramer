package infra

import (
	"database/sql"

	"github.com/Gustrb/kramer/internal/infra/migrations"
	_ "github.com/mattn/go-sqlite3"
)

// Connect with sqlite
func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := migrations.Run(db); err != nil {
		return nil, err
	}

	return db, nil
}
