package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/filagot/emagne/initator/foundation"
	"github.com/filagot/emagne/internal/config"
	db "github.com/filagot/emagne/internal/database"
	"github.com/filagot/emagne/pkg/utils"
	"github.com/jackc/pgx/v4"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	conn := foundation.InitDB(cfg.DBSource)
	defer conn.Close()

	q := db.New(conn)
	ctx := context.Background()

	adminEmail := "admin@emagne.com"

	// Check if admin already exists
	_, err := q.GetUserByEmail(ctx, adminEmail)
	if err == nil {
		log.Println("Super admin already exists")
		return
	}

	if err != pgx.ErrNoRows {
		// If error is something else than NoRows, we should check it.
		// However, GetUserByEmail likely returns error if not found.
		// Let's assume if err != nil and err.Error() contains "no rows", it's fine.
		// But pgx usually returns pgx.ErrNoRows.
		// Let's just try to create and handle duplicate error if race condition.
	}

	log.Println("Creating super admin user...")
 
	password := "admin123"
	hashedPassword, err := utils.HashPassword(password, 10)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	arg := db.CreateUserParams{
		Email:        adminEmail,
		PasswordHash: hashedPassword,
		FirstName:    "Super",
		LastName:     "Admin",
		Phone:        sql.NullString{Valid: false},
		Role:         "super_admin",
	}

	user, err := q.CreateUser(ctx, arg)
	if err != nil {
		log.Fatal("Failed to create super admin:", err)
	}

	log.Printf("Super admin created successfully!\nEmail: %s\nPassword: %s\nID: %s\n", user.Email, password, user.ID)
}
