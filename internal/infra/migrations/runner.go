package migrations

import (
	"database/sql"
	"errors"
)

func Run(db *sql.DB) error {
	if err := CreateContextTable(db); err != nil {
		return errors.New("error creating context table: " + err.Error())
	}

	if err := CreateHistoryTable(db); err != nil {
		return errors.New("error creating context table: " + err.Error())
	}

	return nil
}
