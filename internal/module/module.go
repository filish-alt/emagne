package module

import (
	"context"
	"database/sql"

	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/google/uuid"
)

// AuthModule defines the interface for authentication business logic
type AuthModule interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)
	UpdateUser(ctx context.Context, userID string, req *dto.UpdateUserRequest) (*dto.User, error)
	GetUserByID(ctx context.Context, userID string) (*dto.User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type Catagory interface {
	CreateCategory(ctx context.Context, name string, description sql.NullString) (db.ItemCategory, error)
	GetCategory(ctx context.Context, id uuid.UUID) (db.ItemCategory, error)
	ListCategories(ctx context.Context) ([]db.ItemCategory, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, name string, description sql.NullString) (db.ItemCategory, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error

	CreateAttribute(ctx context.Context, categoryID uuid.UUID, name, dataType string, isRequired sql.NullBool) (db.CategoryAttribute, error)
	GetAttribute(ctx context.Context, categoryID uuid.UUID, name string) (db.CategoryAttribute, error)
	ListAttributes(ctx context.Context, categoryID uuid.UUID) ([]db.CategoryAttribute, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) error
}
