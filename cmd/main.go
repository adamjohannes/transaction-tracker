package main

import (
	"flag"
	"fmt"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/cmd/gui"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/database"
)

func main() {
	serverMode := flag.Bool("server", false, "Run in headless server mode")
	flag.BoolVar(serverMode, "s", false, "Run in headless server mode (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Transaction Tracker is a desktop GUI application and a headless API.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Run without arguments to start the GUI, or use flags to change the mode.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")

		flag.PrintDefaults()
	}

	flag.Parse()

	pool := database.ConnectDB()
	defer pool.Close()

	if *serverMode {
		api.StartServer(pool)
	} else {
		guiConfig := config.NewGUIConfig(607, 1080, true)
		g := gui.NewGUI(guiConfig, pool)
		g.Start()
	}
}
