package dto

import (
	"database/sql"
)

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
	ID          string
	FirstName   string
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

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
