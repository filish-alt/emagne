package module

import (
    "context"

    "github.com/filagot/emagne/internal/database/models/dto"
)

// AuthModule defines the interface for authentication business logic
type AuthModule interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)
	UpdateUser(ctx context.Context, userID string, req *dto.UpdateUserRequest) (*dto.User, error)
	GetUserByID(ctx context.Context, userID string) (*dto.User, error)
}

