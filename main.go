package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	// Setup the SQLite database
	db, err := sql.Open("sqlite", "./actors.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the database schema
	CreateSchema(db)

	// Populate the database with CSV data
	PopulateDatabase(db, "IMDB-actors.csv")

	// Perform SQL queries on the SQLite database
	PerformQuery(db)
}

// CreateSchema creates the actors table schema
func CreateSchema(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS actors (
		id INTEGER,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		gender TEXT NOT NULL
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating schema: %s", err)
	}
	log.Println("Schema created successfully.")
}

// PopulateDatabase populates the actors table with data from a CSV file
func PopulateDatabase(db *sql.DB, csvFile string) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV file:", err)
	}

	// Iterate over the records and insert the data into the database
	for i, record := range records {
		if i == 0 {
			continue // Skip the header
		}

		// Insert the data into the database, including the ID column
		insertStatement := `INSERT OR REPLACE INTO actors (id, first_name, last_name, gender) VALUES (?, ?, ?, ?)`
		_, err := db.Exec(insertStatement, record[0], record[1], record[2], record[3])
		if err != nil {
			log.Fatalf("Error inserting data: %s", err)
		}
	}

	log.Println("Data inserted successfully.")
}

// PerformQuery performs SQL queries on the actors table
func PerformQuery(db *sql.DB) {
	query := `SELECT id, first_name, last_name, gender FROM actors limit 10`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error querying database: %s", err)
	}
	defer rows.Close()

	fmt.Println("ID | First Name | Last Name | Gender")
	fmt.Println("-----------------------------------")
	for rows.Next() {
		var id int
		var firstName, lastName, gender string
		err := rows.Scan(&id, &firstName, &lastName, &gender)
		if err != nil {
			log.Fatalf("Error scanning row: %s", err)
		}
		fmt.Printf("%d | %s | %s | %s\n", id, firstName, lastName, gender)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error with rows: %s", err)
	}
}
