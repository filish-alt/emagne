package storage

import (
	"context"
	"database/sql"
)

// AuthStorage defines the interface for authentication data operations
type AuthStorage interface {
	CreateUser(ctx context.Context, user *CreateUserParams) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *UpdateUserParams) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

// CreateUserParams represents parameters for creating a user
type CreateUserParams struct {
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Phone        sql.NullString
}

// UpdateUserParams represents parameters for updating a user
type UpdateUserParams struct {
	ID         string
	FirstName  string
	LastName    string
	Phone       sql.NullString
	IsVerified  sql.NullBool
}

// User represents a user entity in storage layer
type User struct {
	ID           string         `json:"id"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"password_hash"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Phone        sql.NullString `json:"phone"`
	IsVerified   sql.NullBool   `json:"is_verified"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
}