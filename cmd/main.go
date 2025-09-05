package main

import "monthly-expenses-handler/cmd/gui"

func main() {
	gui := gui.NewGUI()
	gui.Start()
}
