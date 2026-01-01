package transaction

import (
	"context"

	"database/sql"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/storage"
	"github.com/google/uuid"
)

type Storage struct {
	queries *db.Queries
}

func New(queries *db.Queries) storage.TransactionStorage {
	return &Storage{queries: queries}
}

func (s *Storage) Create(ctx context.Context, params db.CreateTransactionParams) (db.Transaction, error) {
	return s.queries.CreateTransaction(ctx, params)
}

func (s *Storage) GetWithCategory(ctx context.Context, id uuid.UUID) (db.GetTransactionWithCategoryRow, error) {
	return s.queries.GetTransactionWithCategory(ctx, id)
}

func (s *Storage) ListByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int32) ([]db.ListTransactionsByCategoryRow, error) {
	return s.queries.ListTransactionsByCategory(ctx, db.ListTransactionsByCategoryParams{ItemCategoryID: categoryID, Limit: limit, Offset: offset})
}

func (s *Storage) UpdateStatus(ctx context.Context, id uuid.UUID, status sql.NullString) (db.Transaction, error) {
	return s.queries.UpdateTransactionStatus(ctx, db.UpdateTransactionStatusParams{ID: id, Status: status})
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteTransaction(ctx, id)
}
