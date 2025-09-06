package gui

import (
	"monthly-expenses-handler/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// GUI
// Holds the application and all its windows, acting as a controller.
type GUI struct {
	app        fyne.App
	homeWindow fyne.Window
	addWindow  fyne.Window
}

// NewGUI
// Creates and initializes the entire application, including all windows.
func NewGUI(config *config.GUIConfig) *GUI {
	a := app.New()
	g := &GUI{app: a}

	g.homeWindow = a.NewWindow("Rastreador de Transações")
	g.addWindow = a.NewWindow("Adicionar Transação")

	g.homeWindow.SetContent(makeHomeContent(g))
	g.addWindow.SetContent(makeAddContent(g))

	g.homeWindow.Resize(fyne.NewSize(config.Width, config.Height))
	g.homeWindow.SetFixedSize(!config.Resizable)
	g.addWindow.Resize(fyne.NewSize(config.Width, config.Height))
	g.addWindow.SetFixedSize(!config.Resizable)

	g.addWindow.SetCloseIntercept(func() {
		g.ShowHomeScreen()
	})

	g.homeWindow.SetCloseIntercept(func() {
		a.Quit()
	})

	return g
}

// Start
// Shows the initial window and runs the application.
func (g *GUI) Start() {
	g.homeWindow.ShowAndRun()
}

// --- Navigation Methods ---

// ShowAddScreen hides the home window and shows the add window.
func (g *GUI) ShowAddScreen() {
	g.addWindow.Show()
	g.homeWindow.Hide()
}

// ShowHomeScreen hides the add window and shows the home window.
func (g *GUI) ShowHomeScreen() {
	g.homeWindow.Show()
	g.addWindow.Hide()
}
