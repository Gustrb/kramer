package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Gustrb/kramer/models"
)

type SqliteContextRepository struct {
	DB *sql.DB
}

func NewSqliteContextRepository(db *sql.DB) *SqliteContextRepository {
	return &SqliteContextRepository{DB: db}
}

const InsertContextIntoDBSQL = `
INSERT INTO context (name, createdAt) VALUES (?, ?)
`

func (s *SqliteContextRepository) Create(ctx context.Context, data models.CreateContext) (models.Context, error) {
	var context models.Context

	stmt, err := s.DB.Prepare(InsertContextIntoDBSQL)
	if err != nil {
		return context, err
	}

	createdAt := time.Now()
	res, err := stmt.ExecContext(ctx, data.Name, createdAt)
	if err != nil {
		return context, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return context, err
	}

	context.CreatedAt = createdAt
	context.ID = int(id)
	context.Name = data.Name

	return context, nil
}

const findContextByNameSQL = `
SELECT id, name, createdAt FROM context WHERE name = ?;
`

var ErrContextNotFound = errors.New("context not found")

func (s *SqliteContextRepository) FindContextByName(ctx context.Context, name string) (models.Context, error) {
	var context models.Context

	stmt, err := s.DB.Prepare(findContextByNameSQL)
	if err != nil {
		return context, err
	}

	row := stmt.QueryRowContext(ctx, name)
	if err := row.Scan(&context.ID, &context.Name, &context.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return context, ErrContextNotFound
		}

		return context, err
	}

	return context, nil
}
