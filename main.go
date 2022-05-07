package main

import (
	"log"
	"pheasant-api/app/models"
	"pheasant-api/routes"

	"github.com/joho/godotenv"
)

// Entrypoint for app.
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Load ENV failed. Err: %s", err)
	}

	// Load the routes
	ginEngine := routes.SetupApiRouter()

	// Initialize database
	models.Initialize(true)

	// Start the HTTP API
	ginEngine.Run()
}
