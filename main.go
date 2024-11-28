package main

import (
	"log"
	"net/http"
	"toronto_time_api/db"
	"toronto_time_api/handlers"
)

func main() {

	db.InitDB()
	// Ensure the database connection is closed when the program exits after
	defer db.DB.Close()

	http.HandleFunc("/current-time", handlers.CurrentTimeHandler)

	log.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
