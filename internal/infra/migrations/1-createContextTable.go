package migrations

import "database/sql"

const createContextTableSQL = `
CREATE TABLE IF NOT EXISTS context (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	createdAt TIMESTAMP
);
`

func CreateContextTable(db *sql.DB) error {
	stmt, err := db.Prepare(createContextTableSQL)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
