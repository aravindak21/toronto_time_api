package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"toronto_time_api/db"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the timezone to Toronto
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		log.Printf("Error loading timezone: %v", err)
		return
	}

	currentTime := time.Now().In(loc)

	// Log time to the database
	err = db.LogTime(currentTime)
	if err != nil {
		http.Error(w, "Failed to log time to database", http.StatusInternalServerError)
		log.Printf("Error logging time: %v", err)
		return
	}

	// Respond with current time in JSON
	response := TimeResponse{CurrentTime: currentTime.Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
