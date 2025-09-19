package main

import (
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/database"
)

func main() {
	pool := database.ConnectDB()
	defer pool.Close()

	api.StartServer(pool)
}
