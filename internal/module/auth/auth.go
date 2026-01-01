package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/database/models/dto"
	"github.com/filagot/emagne/internal/module"
	"github.com/filagot/emagne/internal/storage"
	"github.com/filagot/emagne/pkg/utils"
	"github.com/google/uuid"
)

// AuthModule implements the module.AuthModule interface
type authModule struct {
	authStorage storage.AuthStorage
	config      *config.Config
}

// NewAuthModule creates a new auth module instance
func NewAuthModule(authStorage storage.AuthStorage, cfg *config.Config) module.AuthModule {
	return &authModule{
		authStorage: authStorage,
		config:      cfg,
	}
}

// Register handles user registration
func (m *authModule) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
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

	user, err := m.authStorage.CreateUser(ctx, dto.CreateUserParams{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Phone:        phone,
		Role:         "user",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate token
    userUUID, err := uuid.Parse(user.ID)
    if err != nil {
        return nil, fmt.Errorf("invalid user id: %w", err)
    }
    token, err := utils.GenerateToken(userUUID, user.Email, user.Role, m.config.JWTSecret, m.config.JWTExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.AuthResponse{
		Token: token,
		User: &dto.User{
			ID:           user.ID,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Phone:        user.Phone,
			IsVerified:   user.IsVerified,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			Role:         user.Role,
		},
	}, nil
}

// Login handles user login
func (m *authModule) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
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
    userUUID, err := uuid.Parse(user.ID)
    if err != nil {
        return nil, fmt.Errorf("invalid user id: %w", err)
    }
    token, err := utils.GenerateToken(userUUID, user.Email, user.Role, m.config.JWTSecret, m.config.JWTExpiration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.AuthResponse{
		Token: token,
		User: &dto.User{
			ID:           user.ID,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Phone:        user.Phone,
			IsVerified:   user.IsVerified,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
			Role:         user.Role,
		},
	}, nil
}

// UpdateUser updates the authenticated user's profile
func (m *authModule) UpdateUser(ctx context.Context, userID string, req *dto.UpdateUserRequest) (*dto.User, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	if req == nil {
		return nil, errors.New("update request is required")
	}

	params := &dto.UpdateUserParams{
		ID: userID,
	}

	if req.FirstName != nil {
		params.FirstName = sql.NullString{String: *req.FirstName, Valid: true}
	}

	if req.LastName != nil {
		params.LastName = sql.NullString{String: *req.LastName, Valid: true}
	}

	if req.Phone != nil {
		params.Phone = sql.NullString{String: *req.Phone, Valid: true}
	}

	updatedUser, err := m.authStorage.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// GetUserByID retrieves a user by ID
func (m *authModule) GetUserByID(ctx context.Context, userID string) (*dto.User, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}

	user, err := m.authStorage.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
