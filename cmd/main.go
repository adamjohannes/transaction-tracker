package main

import (
	"monthly-expenses-handler/cmd/gui"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/database"
)

func main() {
	pool := database.ConnectDB()
	defer pool.Close()

	guiConfig := config.NewGUIConfig(607, 1080, true)
	g := gui.NewGUI(guiConfig, pool)
	g.Start()
}
