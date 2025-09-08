package gui

import (
	"fmt"
	domain "monthly-expenses-handler/internal/domain/transaction"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeListContent(g *GUI, transactions []*domain.Transaction) fyne.CanvasObject {
	lblTitle := widget.NewLabel("Lista das Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true

	if transactions == nil {
		var err error
		transactions, err = g.transactionController.GetAllTransactions()
		if err != nil {
			go func() {
				dialog.ShowError(fmt.Errorf("failed to load transactions: %w", err), g.listWindow)
			}()
			btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)
			return container.NewBorder(lblTitle, btnReturn, nil, nil)
		}
	}

	var currentSortColumn = -1
	var isSortAscending = true

	headers := []string{"Amount", "Date", "Type", "Currency", "Category", "Sub Category", "Description"}
	numCols := len(headers)
	grid := container.NewGridWithColumns(numCols)
	scrollableGrid := container.NewScroll(grid)

	var rebuildGrid func()
	rebuildGrid = func() {
		allWidgets := make([]fyne.CanvasObject, 0, (len(transactions)+1)*numCols)

		updateHeaderLabels := func(headerButtons []*widget.Button) {
			for i, h := range headers {
				label := h
				if i == currentSortColumn {
					if isSortAscending {
						label += " ▲"
					} else {
						label += " ▼"
					}
				}
				headerButtons[i].SetText(label)
			}
		}

		headerWidgets := make([]*widget.Button, numCols)
		for i := range headers {
			colIndex := i
			btn := widget.NewButton("", func() {
				if currentSortColumn == colIndex {
					isSortAscending = !isSortAscending
				} else {
					currentSortColumn = colIndex
					isSortAscending = true
				}
				sortTransactions(transactions, currentSortColumn, isSortAscending)
				rebuildGrid()
			})
			headerWidgets[i] = btn
			allWidgets = append(allWidgets, btn)
		}
		updateHeaderLabels(headerWidgets)

		for _, tx := range transactions {
			allWidgets = append(allWidgets, widget.NewLabel(tx.Amount.StringFixed(2)))
			allWidgets = append(allWidgets, widget.NewLabel(tx.Date.Format("2006-01-02")))
			allWidgets = append(allWidgets, widget.NewLabel(tx.Type.Name))
			allWidgets = append(allWidgets, widget.NewLabel(tx.Currency.Code))
			allWidgets = append(allWidgets, widget.NewLabel(tx.Category.Name))
			allWidgets = append(allWidgets, widget.NewLabel(tx.SubCategory.Name))

			descLabel := widget.NewLabel(tx.Description)
			descLabel.Truncation = fyne.TextTruncateEllipsis
			allWidgets = append(allWidgets, descLabel)
		}

		grid.Objects = allWidgets
		grid.Refresh()
	}
	rebuildGrid()

	btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)
	btnRefresh := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameViewRefresh), func() {
		g.listWindow.SetContent(makeListContent(g, nil))
	})

	btnFilter := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameSearch), g.ShowFilterScreen)
	btnCharts := widget.NewButtonWithIcon("Ver Gráficos", theme.Icon(theme.IconNameSettings), g.ShowDataScreenFromList)

	bottomButtons := container.NewHBox(
		btnReturn,
		layout.NewSpacer(),
		btnFilter,
		btnRefresh,
		btnCharts,
	)

	return container.NewBorder(lblTitle, bottomButtons, nil, nil, scrollableGrid)
}

// sortTransactions
// Sorts the slice of transactions in place based on the selected column and direction.
func sortTransactions(transactions []*domain.Transaction, colIndex int, ascending bool) {
	sort.Slice(transactions, func(i, j int) bool {
		var less bool
		switch colIndex {
		case 0: // Amount (Decimal)
			comparison := transactions[i].Amount.Cmp(transactions[j].Amount)
			less = comparison < 0
		case 1: // Date (time.Time)
			less = transactions[i].Date.Before(transactions[j].Date)
		case 2: // Type (string)
			less = transactions[i].Type.Name < transactions[j].Type.Name
		case 3: // Currency (string)
			less = transactions[i].Currency.Code < transactions[j].Currency.Code
		case 4: // Category (string)
			less = transactions[i].Category.Name < transactions[j].Category.Name
		case 5: // Sub Category (string)
			less = transactions[i].SubCategory.Name < transactions[j].SubCategory.Name
		case 6: // Description (string)
			less = transactions[i].Description < transactions[j].Description
		}

		if ascending {
			return less
		}
		return !less
	})
}
