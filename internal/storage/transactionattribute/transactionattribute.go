package transactionattribute

import (
    "context"

    db "github.com/filagot/emagne/internal/database"
    "github.com/filagot/emagne/internal/storage"
    "github.com/google/uuid"
)

type Storage struct {
    queries *db.Queries
}

func New(queries *db.Queries) storage.TransactionItemAttributeStorage {
    return &Storage{queries: queries}
}

func (s *Storage) Insert(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID, value string) (db.TransactionItemAttribute, error) {
    return s.queries.InsertTransactionItemAttribute(ctx, db.InsertTransactionItemAttributeParams{TransactionID: txID, AttributeID: attributeID, Value: value})
}

func (s *Storage) List(ctx context.Context, txID uuid.UUID) ([]db.ListTransactionItemAttributesRow, error) {
    return s.queries.ListTransactionItemAttributes(ctx, txID)
}

func (s *Storage) Get(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID) (db.TransactionItemAttribute, error) {
    return s.queries.GetEscrowItemAttribute(ctx, db.GetEscrowItemAttributeParams{TransactionID: txID, AttributeID: attributeID})
}

func (s *Storage) DeleteAll(ctx context.Context, txID uuid.UUID) error {
    return s.queries.DeleteTransactionItemAttributes(ctx, txID)
}

