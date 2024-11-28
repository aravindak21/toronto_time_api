package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Initialize the database connection by accepting the dsn
func InitDB(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		return err
	}

	// Check if the database is accessible
	err = DB.Ping()
	if err != nil {
		log.Printf("Error pinging the database: %v", err)
		return err
	}

	log.Println("Database connection established.")
	return nil
}
