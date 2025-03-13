package migrations

import "database/sql"

const createHistoryTableSQL = `
CREATE TABLE IF NOT EXISTS history (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	message TEXT,
	role TEXT,
	contextId INTEGER,
	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	FOREIGN KEY(contextId) REFERENCES context(id)
);
`

func CreateHistoryTable(db *sql.DB) error {
	stmt, err := db.Prepare(createHistoryTableSQL)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
