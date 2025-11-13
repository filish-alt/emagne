package storage

import (
	"context"

	"github.com/filagot/emagne/internal/database/models/dto"
)

// AuthStorage defines the interface for authentication data operations
type AuthStorage interface {
	CreateUser(ctx context.Context, user dto.CreateUserParams) (*dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
	GetUserByID(ctx context.Context, id string) (*dto.User, error)
	UpdateUser(ctx context.Context, user *dto.UpdateUserParams) (*dto.User, error)
	DeleteUser(ctx context.Context, id string) error
}

