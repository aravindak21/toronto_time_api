package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:Aravind@sv123@tcp(localhost:3306)/toronto_time"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Connected to database successfully.")
}

func LogTime(currentTime time.Time) error {
	_, err := DB.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	return err
}
