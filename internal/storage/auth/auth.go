package auth

import (
    "context"
    "fmt"

    db "github.com/filagot/emagne/internal/database"
    "github.com/filagot/emagne/internal/database/models/dto"
	persistancedb "github.com/filagot/emagne/internal/database/persistancedb"
    "github.com/google/uuid"
	"github.com/filagot/emagne/internal/storage"
)

// AuthStorage implements the storage.AuthStorage interface
type AuthStorage struct {
   persistanceQueries persistancedb.PersistenceDB
}

// NewAuthStorage creates a new auth storage instance
func NewAuthStorage(persistanceQueries persistancedb.PersistenceDB) storage.AuthStorage {
	return &AuthStorage{
		persistanceQueries: persistanceQueries,
	}
}

// CreateUser creates a new user in the database
func (s *AuthStorage) CreateUser(ctx context.Context, user dto.CreateUserParams) (*dto.User, error) {
    dbUser, err := s.persistanceQueries.CreateUser(ctx, db.CreateUserParams{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Phone:        user.Phone,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &dto.User{
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
func (s *AuthStorage) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	dbUser, err := s.persistanceQueries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &dto.User{
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
func (s *AuthStorage) GetUserByID(ctx context.Context, id string) (*dto.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	dbUser, err := s.persistanceQueries.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &dto.User{
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
func (s *AuthStorage) UpdateUser(ctx context.Context, user *dto.UpdateUserParams) (*dto.User, error) {
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	dbUser, err := s.persistanceQueries.UpdateUser(ctx, db.UpdateUserParams{
		ID:         userID,
		FirstName:  user.FirstName,
		Phone:      user.Phone,
		IsVerified: user.IsVerified,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &dto.User{
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

	return s.persistanceQueries.DeleteUser(ctx, userID)
}

