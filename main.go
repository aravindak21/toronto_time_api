package main

import (
	"log"
	"net/http"
	"toronto_time_api/db"
	"toronto_time_api/handlers"
)

func main() {
	// Initialize the database connection
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.DB.Close()

	// Register handlers
	http.HandleFunc("/current-time", handlers.CurrentTimeHandler)
	http.HandleFunc("/logged-times", handlers.GetLoggedTimesHandler) // New handler for logged times

	// Start the server
	log.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
