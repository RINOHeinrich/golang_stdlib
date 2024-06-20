package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Replace the connection parameters with your own
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres  sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("Connected to PostgreSQL database!")

	return db
}
func CreateBookTable(db *sql.DB) error {
	// Create the book table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS book (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			published_date DATE NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Book table created!")
	return nil
}

var DatabaseConnection = Connect()
