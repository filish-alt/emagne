package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/module"
	"github.com/filagot/emagne/internal/storage"
	"github.com/filagot/emagne/pkg/utils"
)

// AuthModule implements the module.AuthModule interface
type AuthModule struct {
	authStorage storage.AuthStorage
	config      *config.Config
}

// NewAuthModule creates a new auth module instance
func NewAuthModule(authStorage storage.AuthStorage, cfg *config.Config) module.AuthModule {
	return &AuthModule{
		authStorage: authStorage,
		config:      cfg,
	}
}

// Register handles user registration
func (m *AuthModule) Register(ctx context.Context, req *module.RegisterRequest) (*module.AuthResponse, error) {
	// Check if user already exists
	_, err := m.authStorage.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password, m.config.PasswordCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	phone := sql.NullString{}
	if req.Phone != "" {
		phone.String = req.Phone
		phone.Valid = true
	}

	user, err := m.authStorage.CreateUser(ctx, &storage.CreateUserParams{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        phone,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, "user", m.config.JWTSecret, m.config.JWTExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &module.AuthResponse{
		Token: token,
		User: &module.User{
			ID:           user.ID,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Phone:        user.Phone,
			IsVerified:   user.IsVerified,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	}, nil
}

// Login handles user login
func (m *AuthModule) Login(ctx context.Context, req *module.LoginRequest) (*module.AuthResponse, error) {
	// Get user by email
	user, err := m.authStorage.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if err := utils.CheckPassword(req.Password, user.PasswordHash); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, "user", m.config.JWTSecret, m.config.JWTExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &module.AuthResponse{
		Token: token,
		User: &module.User{
			ID:           user.ID,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Phone:        user.Phone,
			IsVerified:   user.IsVerified,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	}, nil
}
