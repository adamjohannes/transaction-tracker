package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// --- Test Case 1: Successful load with all variables set ---
	t.Run("Successful load", func(t *testing.T) {
		// Set environment variables for the test
		os.Setenv("DB_HOST", "testhost")
		os.Setenv("DB_PORT", "1234")
		os.Setenv("DB_USER", "testuser")
		os.Setenv("DB_PASSWORD", "testpass")
		os.Setenv("DB_NAME", "testdb")
		// Unset them after the test
		defer func() {
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_PORT")
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_NAME")
		}()

		cfg, err := Load()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		if cfg.Host != "testhost" {
			t.Errorf("expected host 'testhost', got '%s'", cfg.Host)
		}
		if cfg.User != "testuser" {
			t.Errorf("expected user 'testuser', got '%s'", cfg.User)
		}
	})

	// --- Test Case 2: Error on missing essential variable ---
	t.Run("Error on missing variable", func(t *testing.T) {
		// Ensure a critical variable is not set
		os.Unsetenv("DB_USER")

		_, err := Load()
		if err == nil {
			t.Fatal("expected an error for missing config, but got none")
		}
	})

	// --- Test Case 3: Successful load with fallback values ---
	t.Run("Successful load with fallbacks", func(t *testing.T) {
		// Set only the required variables
		os.Setenv("DB_HOST", "testhost")
		os.Setenv("DB_USER", "testuser")
		os.Setenv("DB_PASSWORD", "testpass")
		os.Setenv("DB_NAME", "testdb")
		// Unset the optional one
		os.Unsetenv("DB_PORT")
		defer func() {
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_USER")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_NAME")
		}()

		cfg, err := Load()
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		// Check that the port has fallen back to the default
		if cfg.Port != "5432" {
			t.Errorf("expected fallback port '5432', got '%s'", cfg.Port)
		}
	})
}
