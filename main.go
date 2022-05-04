package main

import (
	"pheasant-api/app/models"
	"pheasant-api/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	ginEngine := routes.SetupApiRouter()

	// Initialize database
	models.Initialize()

	// Start the HTTP API
	ginEngine.Run()
}
