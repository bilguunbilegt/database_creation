package main

import (
	"database/sql"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTestDatabase() (*sql.DB, error) {
	// Create a new in-memory SQLite database for testing
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create the schema
	CreateSchema(db)

	return db, nil
}

func TestCreateSchema(t *testing.T) {
	db, err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Failed to setup test database: %s", err)
	}
	defer db.Close()

	// Check if the actors table was created
	_, err = db.Query("SELECT 1 FROM actors LIMIT 1")
	if err != nil {
		t.Errorf("Table actors does not exist: %s", err)
	}
}

func TestPopulateDatabase(t *testing.T) {
	db, err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Failed to setup test database: %s", err)
	}
	defer db.Close()

	// Create a temporary CSV file for testing
	csvFile, err := os.CreateTemp("", "test-actors-*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp CSV file: %s", err)
	}
	defer os.Remove(csvFile.Name())

	// Write test data to the CSV file
	testData := `id,first_name,last_name,gender
1,Emma,Watson,F
2,Jennifer,Lawrence,F`
	_, err = csvFile.WriteString(testData)
	if err != nil {
		t.Fatalf("Failed to write to temp CSV file: %s", err)
	}
	csvFile.Close()

	// Populate the database with the test data
	PopulateDatabase(db, csvFile.Name())

	// Check if data was inserted correctly
	rows, err := db.Query("SELECT id, first_name, last_name, gender FROM actors")
	if err != nil {
		t.Fatalf("Failed to query database: %s", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}
	if count != 2 {
		t.Errorf("Expected 2 rows in the database, got %d", count)
	}
}

func TestQueryDatabase(t *testing.T) {
	db, err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Failed to setup test database: %s", err)
	}
	defer db.Close()

	// Insert test data
	insertStatement := `INSERT INTO actors (id, first_name, last_name, gender) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(insertStatement, "1", "Emma", "Watson", "F")
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		t.Fatalf("Expected 1 row affected, got %d", rowsAffected)
	}

	_, err = db.Exec(insertStatement, "2", "Jennifer", "Lawrence", "F")
	if err != nil {
		t.Fatalf("Failed to insert data: %s", err)
	}

	// Query and verify the data
	rows, err := db.Query("SELECT id, first_name, last_name, gender FROM actors")
	if err != nil {
		t.Fatalf("Failed to query database: %s", err)
	}
	defer rows.Close()

	expectedResults := [][]string{
		{"1", "Emma", "Watson", "F"},
		{"2", "Jennifer", "Lawrence", "F"},
	}

	var i int
	for rows.Next() {
		var id, firstName, lastName, gender string
		err := rows.Scan(&id, &firstName, &lastName, &gender)
		if err != nil {
			t.Fatalf("Failed to scan row: %s", err)
		}

		if id != expectedResults[i][0] || firstName != expectedResults[i][1] || lastName != expectedResults[i][2] || gender != expectedResults[i][3] {
			t.Errorf("Unexpected result. Got: %s %s %s %s, Expected: %s %s %s %s",
				id, firstName, lastName, gender,
				expectedResults[i][0], expectedResults[i][1], expectedResults[i][2], expectedResults[i][3])
		}
		i++
	}

	if i != len(expectedResults) {
		t.Errorf("Expected %d rows, got %d", len(expectedResults), i)
	}
}
