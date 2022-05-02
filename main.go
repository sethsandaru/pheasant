package main

import (
	"pheasant-api/database"
	"pheasant-api/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	r := routes.SetupApiRouter()

	// Initialize database
	database.Initialize()

	// Start the HTTP API
	r.Run()
}
