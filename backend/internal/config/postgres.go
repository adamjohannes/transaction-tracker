package config

import (
	"fmt"
	"os"
)

// PostgresConfig
// Holds all configuration variables for the database connection.
type PostgresConfig struct {
	Host          string
	Port          string
	User          string
	Password      string
	Name          string
	EncryptionKey string
	JWTSecret     string
	SearchHashKey string
}

// Load
// Loads the PostgreSQL configuration from environment variables.
func Load() (*PostgresConfig, error) {
	cfg := &PostgresConfig{
		Host:          getEnv("DB_HOST", "localhost"),
		Port:          getEnv("DB_PORT", "5432"),
		User:          getEnv("DB_USER", "postgres"),
		Password:      getEnv("DB_PASSWORD", "password"),
		Name:          getEnv("DB_NAME", "postgres"),
		EncryptionKey: getEnv("ENCRYPTION_KEY", ""),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		SearchHashKey: getEnv("SEARCH_HASH_KEY", ""),
	}

	// Validate that essential variables are present
	if cfg.Host == "" || cfg.Name == "" || cfg.User == "" || cfg.Password == "" {
		return nil, fmt.Errorf("database configuration is incomplete; ensure DB_HOST, DB_NAME, DB_USER, and DB_PASSWORD are set")
	}

	// Validate the encryption key
	if cfg.EncryptionKey == "" {
		return nil, fmt.Errorf("security configuration is incomplete; ensure ENCRYPTION_KEY is set")
	}

	// Validate JWT secret
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("security configuration is incomplete; ensure JWT_SECRET is set")
	}

	// Validate Search Hash Key
	if cfg.SearchHashKey == "" {
		return nil, fmt.Errorf("security configuration is incomplete; ensure SEARCH_HASH_KEY is set")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
