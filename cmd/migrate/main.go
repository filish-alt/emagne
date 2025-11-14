package main

import (
	"log"

	"github.com/filagot/emagne/initator"
	"github.com/filagot/emagne/internal/config"
)

func main() {
	log.Println("Starting database migration...")
	
	// Load configuration
	cfg := config.Load()
	
	// Run migrations
	initator.InitiateMigration(cfg)
	
	log.Println("Migration completed successfully!")
}







