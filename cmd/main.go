package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	var version string
	if err := pool.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Fetching all currencies...")

	// The SQL query to select all codes from the currencies table.
	querySQL := "SELECT code FROM currencies"

	// Use pool.Query() for queries that return multiple rows.
	rows, err := pool.Query(context.Background(), querySQL)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	// It's important to close the rows when you're done with them to free the connection.
	defer rows.Close()

	// Create a slice to hold the results.
	var currenciesstring []string

	// Loop through the returned rows.
	for rows.Next() {
		var code string
		// Scan the value from the current row into the 'code' variable.
		if err := rows.Scan(&code); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		// Add the currency code to our slice.
		currenciesstring = append(currenciesstring, code)
	}

	// After the loop, check for any errors that may have occurred during iteration.
	if err := rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
	}

	// Print the results.
	fmt.Println("Found currencies:", currenciesstring)

	log.Println("Successfully connected to:", version)
}
