package gui

import (
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/sub_category"
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
	listWindow            fyne.Window
	filterWindow          fyne.Window
	dataWindow            fyne.Window
	transactionController *transaction.TransactionController
	categoryController    *category.CategoryController
	subCategoryController *sub_category.SubCategoryController
}

// NewGUI
// Creates and initializes the entire application, including all windows.
func NewGUI(guiConfig *config.GUIConfig, pool *pgxpool.Pool) *GUI {
	a := app.New()
	g := &GUI{app: a}

	g.transactionController = transaction.NewTransactionController(pool, context.Background())
	g.categoryController = category.NewCategoryController(pool, context.Background())
	g.subCategoryController = sub_category.NewSubCategoryController(pool, context.Background())

	g.homeWindow = a.NewWindow("Rastreador de Transações")
	g.addWindow = a.NewWindow("Adicionar Transação")
	g.listWindow = a.NewWindow("Listar Transações")
	g.filterWindow = a.NewWindow("Filtrar Transações")
	g.dataWindow = a.NewWindow("Gráficos das Transações")

	g.homeWindow.SetContent(makeHomeContent(g))
	g.addWindow.SetContent(makeAddContent(g))
	g.listWindow.SetContent(makeListContent(g, nil))
	g.filterWindow.SetContent(makeFilterContent(g))
	g.dataWindow.SetContent(makeDataContent(g))

	g.homeWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.homeWindow.SetFixedSize(!guiConfig.Resizable)
	g.addWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.addWindow.SetFixedSize(!guiConfig.Resizable)
	g.listWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.listWindow.SetFixedSize(!guiConfig.Resizable)
	g.filterWindow.Resize(fyne.NewSize(guiConfig.Width, guiConfig.Height))
	g.filterWindow.SetFixedSize(!guiConfig.Resizable)
	g.dataWindow.Resize(fyne.NewSize(guiConfig.Height, guiConfig.Height))
	g.dataWindow.SetFixedSize(!guiConfig.Resizable)

	g.addWindow.SetCloseIntercept(func() {
		g.ShowHomeScreen()
	})

	g.listWindow.SetCloseIntercept(func() {
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

func (g *GUI) ShowAddScreen() {
	g.addWindow.Show()
	g.homeWindow.Hide()
}

func (g *GUI) ShowListScreen(hideHome bool) {
	if hideHome {
		// If coming from the home screen, refresh the list with all transactions
		g.listWindow.SetContent(makeListContent(g, nil))
		g.homeWindow.Hide()
	}

	// If coming from the filter screen, the content has already been set
	g.listWindow.Show()
	g.filterWindow.Hide()
}

func (g *GUI) ShowFilterScreen() {
	g.filterWindow.Show()
	g.listWindow.Hide()
}

func (g *GUI) ShowDataScreen() {
	g.dataWindow.Show()
	g.homeWindow.Hide()
}

// ShowHomeScreen hides the add window and shows the home window.
func (g *GUI) ShowHomeScreen() {
	g.homeWindow.Show()
	g.addWindow.Hide()
	g.listWindow.Hide()
	g.filterWindow.Hide()
	g.dataWindow.Hide()
}
