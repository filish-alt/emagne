package category

import (
	"context"
	"database/sql"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/module"
	"github.com/filagot/emagne/internal/storage"
	"github.com/google/uuid"
)


type catagoryModule struct {
    cat storage.ItemCategoryStorage
    attr storage.CategoryAttributeStorage
}

func New(cat storage.ItemCategoryStorage, attr storage.CategoryAttributeStorage) module.Catagory {
    return &catagoryModule{
        cat: cat,
        attr: attr}
}

func (m *catagoryModule) CreateCategory(ctx context.Context, name string, description sql.NullString) (db.ItemCategory, error) {
    return m.cat.Create(ctx, name, description)
}

func (m *catagoryModule) GetCategory(ctx context.Context, id uuid.UUID) (db.ItemCategory, error) {
    return m.cat.GetByID(ctx, id)
}

func (m *catagoryModule) ListCategories(ctx context.Context) ([]db.ItemCategory, error) {
    return m.cat.List(ctx)
}

func (m *catagoryModule) UpdateCategory(ctx context.Context, id uuid.UUID, name string, description sql.NullString) (db.ItemCategory, error) {
    return m.cat.Update(ctx, id, name, description)
}

func (m *catagoryModule) DeleteCategory(ctx context.Context, id uuid.UUID) error {
    return m.cat.Delete(ctx, id)
}

func (m *catagoryModule) CreateAttribute(ctx context.Context, categoryID uuid.UUID, name, dataType string, isRequired sql.NullBool) (db.CategoryAttribute, error) {
    return m.attr.Create(ctx, categoryID, name, dataType, isRequired)
}

func (m *catagoryModule) GetAttribute(ctx context.Context, categoryID uuid.UUID, name string) (db.CategoryAttribute, error) {
    return m.attr.Get(ctx, categoryID, name)
}

func (m *catagoryModule) ListAttributes(ctx context.Context, categoryID uuid.UUID) ([]db.CategoryAttribute, error) {
    return m.attr.List(ctx, categoryID)
}

func (m *catagoryModule) DeleteAttribute(ctx context.Context, id uuid.UUID) error {
    return m.attr.Delete(ctx, id)
}

