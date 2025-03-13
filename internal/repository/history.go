package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Gustrb/kramer/models"
)

type SqliteHistoryRepository struct {
	DB *sql.DB
}

func NewSqliteHistoryRepository(db *sql.DB) *SqliteHistoryRepository {
	return &SqliteHistoryRepository{DB: db}
}

const insertHistorySQL = `
INSERT INTO history (contextId, message, role, createdAt) VALUES (?, ?, ?, ?), (?, ?, ?, ?);
`

func (s *SqliteHistoryRepository) Create(ctx context.Context, req models.CreateHistoryEntry, res models.CreateHistoryEntry) error {
	stmt, err := s.DB.Prepare(insertHistorySQL)
	if err != nil {
		return err
	}

	createdAt := time.Now()
	_, err = stmt.ExecContext(ctx, req.ContextID, req.Message, req.Role, createdAt, res.ContextID, res.Message, res.Role, createdAt)
	if err != nil {
		return err
	}

	return nil
}

const readByContextIDSQL = `
SELECT h.id, h.message, h.role, h.createdAt, c.id as contextID, c.name as contextName, c.createdAt as contextCreatedAt
FROM history h
JOIN context c ON h.contextId = c.id
WHERE c.id = ?
ORDER BY h.createdAt ASC;
`

func (s *SqliteHistoryRepository) ReadByContextID(ctx context.Context, contextID int) ([]models.HistoryEntry, error) {
	stmt, err := s.DB.Prepare(readByContextIDSQL)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, contextID)
	if err != nil {
		return nil, err
	}

	var history []models.HistoryEntry
	for rows.Next() {
		h := models.HistoryEntry{}
		if err := rows.Scan(&h.ID, &h.Message, &h.Role, &h.CreatedAt, &h.Context.ID, &h.Context.Name, &h.Context.CreatedAt); err != nil {
			return nil, err
		}
	}

	return history, nil
}
