package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/internal/storage"
	"github.com/google/uuid"
)

// AuthStorage implements the storage.AuthStorage interface
type AuthStorage struct {
	queries *database.Queries
}

// NewAuthStorage creates a new auth storage instance
func NewAuthStorage(queries *database.Queries) storage.AuthStorage {
	return &AuthStorage{
		queries: queries,
	}
}

// CreateUser creates a new user in the database
func (s *AuthStorage) CreateUser(ctx context.Context, user *storage.CreateUserParams) (*storage.User, error) {
	dbUser, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        user.Phone,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &storage.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		Phone:        dbUser.Phone,
		IsVerified:   dbUser.IsVerified,
		CreatedAt:    dbUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    dbUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetUserByEmail retrieves a user by email
func (s *AuthStorage) GetUserByEmail(ctx context.Context, email string) (*storage.User, error) {
	dbUser, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &storage.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		Phone:        dbUser.Phone,
		IsVerified:   dbUser.IsVerified,
		CreatedAt:    dbUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    dbUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthStorage) GetUserByID(ctx context.Context, id string) (*storage.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	dbUser, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &storage.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		Phone:        dbUser.Phone,
		IsVerified:   dbUser.IsVerified,
		CreatedAt:    dbUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    dbUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// UpdateUser updates a user in the database
func (s *AuthStorage) UpdateUser(ctx context.Context, user *storage.UpdateUserParams) (*storage.User, error) {
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	dbUser, err := s.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:         userID,
		FirstName:  user.FirstName,
		Phone:      user.Phone,
		IsVerified: user.IsVerified,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &storage.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		Phone:        dbUser.Phone,
		IsVerified:   dbUser.IsVerified,
		CreatedAt:    dbUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    dbUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// DeleteUser deletes a user from the database
func (s *AuthStorage) DeleteUser(ctx context.Context, id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	return s.queries.DeleteUser(ctx, userID)
}
