package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"toronto_time_api/db"
)

// CurrentTimeHandler handles the /current-time endpoint
func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().In(time.FixedZone("EST", -5*60*60)) // Adjust to Toronto timezone
	log.Printf("Received request for current time: %v", currentTime)

	// Log time in the database
	_, err := db.DB.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		log.Printf("Error logging time to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send the response
	response := map[string]string{
		"current_time": currentTime.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully responded with current time: %v", currentTime)
}

// GetLoggedTimesHandler handles the /logged-times endpoint
func GetLoggedTimesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		http.Error(w, "Failed to retrieve logged times", http.StatusInternalServerError)
		log.Printf("Error retrieving logged times: %v", err)
		return
	}
	defer rows.Close()

	var times []map[string]interface{}
	for rows.Next() {
		var id int
		var timestamp time.Time
		err = rows.Scan(&id, &timestamp)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		times = append(times, map[string]interface{}{
			"id":        id,
			"timestamp": timestamp,
		})
	}

	if len(times) == 0 {
		log.Println("No logged times found.")
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(times); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
