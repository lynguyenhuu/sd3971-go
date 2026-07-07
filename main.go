package main

import (
	"log"

	"sd3971-go/config"
	"sd3971-go/internal/app"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize application
	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v\n", err)
	}
	defer application.Close()

	// Start server
	if err := application.Start(); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
