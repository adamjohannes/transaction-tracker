package main

import (
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/database"
	"monthly-expenses-handler/internal/logger"
)

func main() {
	log := logger.New()

	pool := database.ConnectDB()
	defer pool.Close()

	api.StartServer(pool, log)
}
