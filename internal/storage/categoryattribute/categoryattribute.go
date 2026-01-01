package categoryattribute

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

func New(queries *db.Queries) storage.CategoryAttributeStorage {
    return &Storage{queries: queries}
}

func (s *Storage) Create(ctx context.Context, categoryID uuid.UUID, name, dataType string, isRequired sql.NullBool) (db.CategoryAttribute, error) {
    return s.queries.CreateCategoryAttribute(ctx, db.CreateCategoryAttributeParams{CategoryID: categoryID, Name: name, DataType: dataType, IsRequired: isRequired})
}

func (s *Storage) Get(ctx context.Context, categoryID uuid.UUID, name string) (db.CategoryAttribute, error) {
    return s.queries.GetCategoryAttribute(ctx, db.GetCategoryAttributeParams{CategoryID: categoryID, Name: name})
}

func (s *Storage) List(ctx context.Context, categoryID uuid.UUID) ([]db.CategoryAttribute, error) {
    return s.queries.ListCategoryAttributes(ctx, categoryID)
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
    return s.queries.DeleteCategoryAttribute(ctx, id)
}

