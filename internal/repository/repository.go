package repository

import (
	"context"
	"database/sql"

	"github.com/Gustrb/kramer/models"
)

type ContextRepository interface {
	Create(ctx context.Context, data models.CreateContext) (models.Context, error)
	FindContextByName(ctx context.Context, name string) (models.Context, error)
}

type HistoryRepository interface {
	Create(ctx context.Context, req models.CreateHistoryEntry, res models.CreateHistoryEntry) error
	ReadByContextID(ctx context.Context, contextID int) ([]models.HistoryEntry, error)
}

type Store interface {
	Context() ContextRepository
	History() HistoryRepository
}

func StoreFactory(db *sql.DB) Store {
	return &sqliteStore{
		db:                db,
		contextRepository: NewSqliteContextRepository(db),
		historyRepository: NewSqliteHistoryRepository(db),
	}
}

type sqliteStore struct {
	db                *sql.DB
	contextRepository ContextRepository
	historyRepository HistoryRepository
}

func (s *sqliteStore) Context() ContextRepository {
	return s.contextRepository
}

func (s *sqliteStore) History() HistoryRepository {
	return s.historyRepository
}
