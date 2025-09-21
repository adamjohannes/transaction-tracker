package main

import (
	"log"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/database"
	"monthly-expenses-handler/internal/logger"
)

func main() {
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer pool.Close()

	slogLogger := logger.New(pool)
	slogLogger.Info("Successfully connected to the database and logger initialized")

	api.StartServer(pool, slogLogger)
}
