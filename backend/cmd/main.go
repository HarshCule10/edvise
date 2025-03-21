package main

import (
	"backend/api"
	"backend/config"
	"log"
	"net/http"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Connect to Google Cloud Storage
	config.ConnectGCS()

	// Register API routes
	api.RegisterRoutes()

	// Start the server
	log.Println("🚀 Server starting on :8080 OwO")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
