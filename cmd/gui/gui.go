package gui

import (
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/controller/transaction"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

// GUI
// Holds the application and all its windows, acting as a controller.
type GUI struct {
	app                   fyne.App
	homeWindow            fyne.Window
	addWindow             fyne.Window
	transactionController *transaction.TransactionController
}

// NewGUI
// Creates and initializes the entire application, including all windows.
func NewGUI(guiConfig *config.GUIConfig, pool *pgxpool.Pool) *GUI {
	a := app.New()
	g := &GUI{app: a}

	g.transactionController = transaction.NewTransactionController(pool, context.Background())

	g.homeWindow = a.NewWindow("Rastreador de Transações")
	g.addWindow = a.NewWindow("Adicionar Transação")

	g.homeWindow.SetContent(makeHomeContent(g))
	g.addWindow.SetContent(makeAddContent(g))

	g.homeWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.homeWindow.SetFixedSize(!guiConfig.Resizable)
	g.addWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.addWindow.SetFixedSize(!guiConfig.Resizable)

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
