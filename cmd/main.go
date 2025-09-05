package main

import (
	"monthly-expenses-handler/cmd/gui"
	"monthly-expenses-handler/internal/config"
)

func main() {
	guiConfig := config.NewGUIConfig(400, 800, false)
	gui := gui.NewGUI(guiConfig)
	gui.Start()
}
