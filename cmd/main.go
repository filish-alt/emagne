package main

import (
	"log"

	"github.com/filagot/emagne/initator"
)

func main() {
	// Initialize the application using the initiator pattern
	app, err := initator.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application
	app.Start()
}
