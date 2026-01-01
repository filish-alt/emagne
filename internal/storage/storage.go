package storage

import (
	"context"
	"database/sql"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/google/uuid"
)

// AuthStorage defines the interface for authentication data operations
type AuthStorage interface {
	CreateUser(ctx context.Context, user dto.CreateUserParams) (*dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	GetUserByID(ctx context.Context, id string) (*dto.User, error)
	UpdateUser(ctx context.Context, user *dto.UpdateUserParams) (*dto.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type ItemCategoryStorage interface {
	Create(ctx context.Context, name string, description sql.NullString) (db.ItemCategory, error)
	GetByID(ctx context.Context, id uuid.UUID) (db.ItemCategory, error)
	List(ctx context.Context) ([]db.ItemCategory, error)
	Update(ctx context.Context, id uuid.UUID, name string, description sql.NullString) (db.ItemCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CategoryAttributeStorage interface {
	Create(ctx context.Context, categoryID uuid.UUID, name, dataType string, isRequired sql.NullBool) (db.CategoryAttribute, error)
	Get(ctx context.Context, categoryID uuid.UUID, name string) (db.CategoryAttribute, error)
	List(ctx context.Context, categoryID uuid.UUID) ([]db.CategoryAttribute, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TransactionStorage interface {
	Create(ctx context.Context, params db.CreateTransactionParams) (db.Transaction, error)
	GetWithCategory(ctx context.Context, id uuid.UUID) (db.GetTransactionWithCategoryRow, error)
	ListByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int32) ([]db.ListTransactionsByCategoryRow, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status sql.NullString) (db.Transaction, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TransactionItemAttributeStorage interface {
	Insert(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID, value string) (db.TransactionItemAttribute, error)
	List(ctx context.Context, txID uuid.UUID) ([]db.ListTransactionItemAttributesRow, error)
	Get(ctx context.Context, txID uuid.UUID, attributeID uuid.UUID) (db.TransactionItemAttribute, error)
	DeleteAll(ctx context.Context, txID uuid.UUID) error
}
