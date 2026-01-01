package itemcategory

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

func New(queries *db.Queries) storage.ItemCategoryStorage {
    return &Storage{queries: queries}
}

func (s *Storage) Create(ctx context.Context, name string, description sql.NullString) (db.ItemCategory, error) {
    return s.queries.CreateItemCategory(ctx, db.CreateItemCategoryParams{Name: name, Description: description})
}

func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (db.ItemCategory, error) {
    return s.queries.GetItemCategoryByID(ctx, id)
}

func (s *Storage) List(ctx context.Context) ([]db.ItemCategory, error) {
    return s.queries.ListItemCategories(ctx)
}

func (s *Storage) Update(ctx context.Context, id uuid.UUID, name string, description sql.NullString) (db.ItemCategory, error) {
    return s.queries.UpdateItemCategory(ctx, db.UpdateItemCategoryParams{ID: id, Name: name, Description: description})
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
    return s.queries.DeleteItemCategory(ctx, id)
}

