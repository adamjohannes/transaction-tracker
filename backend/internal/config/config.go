package config

import (
	"io"
	"monthly-expenses-handler/internal/logger"
	"os"
)

type Config struct {
	Logger   *loggerConfig   `json:"logger"`
	Postgres *postgresConfig `json:"postgres"`
}

// Load
// Loads the project's configuration from environment variables.
func Load() *Config {
	return &Config{
		Logger:   loadLoggerConfig(),
		Postgres: loadPostgresConfig(),
	}
}

// --- Logger

func loadLoggerConfig() *loggerConfig {
	out := loadLoggerOutputStream()
	level := loadLoggerLevel()

	return &loggerConfig{
		out,
		level,
	}
}

// --- Postgres

func loadPostgresConfig() *postgresConfig {
	cfg := &postgresConfig{
		Host:          getEnv("DB_HOST", "localhost"),
		Port:          getEnv("DB_PORT", "5432"),
		User:          getEnv("DB_USER", "postgres"),
		Password:      getEnv("DB_PASSWORD", "password"),
		Name:          getEnv("DB_NAME", "postgres"),
		EncryptionKey: getEnv("ENCRYPTION_KEY", ""),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		SearchHashKey: getEnv("SEARCH_HASH_KEY", ""),
	}

	return cfg
}

// --- Helpers

func loadLoggerOutputStream() io.Writer {
	// TODO: add a check for other output streams here
	return os.Stdout
}

func loadLoggerLevel() logger.Level {
	if value, ok := os.LookupEnv("LEVEL"); ok {
		switch value {
		case "debug":
			return logger.DEBUG
		case "info":
			return logger.INFO
		case "warn":
			return logger.WARNING
		case "error":
			return logger.ERROR
		case "fatal":
			return logger.FATAL
		}
	}
	return logger.INFO
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
