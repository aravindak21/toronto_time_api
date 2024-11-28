package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB initializes the database connection and logs errors or success.
func InitDB() error {
	dsn := "root:Aravind@sv123@tcp(localhost:3306)/toronto_time"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error connecting to the database: %v", err) // Log the error
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check if the connection is alive
	if err := DB.Ping(); err != nil {
		log.Printf("Error pinging the database: %v", err) // Log the error
		return fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Database connection established successfully") // Log success
	return nil
}
