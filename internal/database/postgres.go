package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB establishes a connection pool to the PostgreSQL database.
// It builds the connection URI from individual environment variables.
func ConnectDB() *pgxpool.Pool {
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	dbname := getEnv("DB_NAME", "postgres")

	// Construct the full database URI
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	// Log the connection details for debugging (without the password)
	log.Printf("Attempting to connect to: postgresql://%s:***@%s:%s/%s", user, host, port, dbname)

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// Verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")
	return pool
}

// getEnv is a helper function to read an environment variable with a fallback value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
