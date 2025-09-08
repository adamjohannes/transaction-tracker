package gui

import (
	"fmt"
	domain "monthly-expenses-handler/internal/domain/transaction"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	// -1 means no column is sorted initially.
	var currentSortColumn = -1
	var isSortAscending = true
	var table *widget.Table

	headers := []string{"Amount", "Date", "Type", "Currency", "Category", "Sub Category", "Description"}
	headerButtons := make([]fyne.CanvasObject, len(headers))

	updateHeaderLabels := func() {
		for i, h := range headers {
			btn := headerButtons[i].(*widget.Button)
			label := h
			if i == currentSortColumn {
				if isSortAscending {
					label += " ▲" // Ascending order
				} else {
					label += " ▼" // Descending order
				}
			}
			btn.SetText(label)
		}
	}

	for i, header := range headers {
		colIndex := i
		headerButtons[i] = widget.NewButton(header, func() {
			if currentSortColumn == colIndex {
				isSortAscending = !isSortAscending // Flip direction
			} else {
				currentSortColumn = colIndex // New column, sort ascending
				isSortAscending = true
			}

			sortTransactions(transactions, currentSortColumn, isSortAscending)
			updateHeaderLabels()
			table.Refresh()
		})
	}

	customHeader := container.NewGridWithColumns(len(headers), headerButtons...)
	table = widget.NewTable(
		func() (int, int) {
			return len(transactions), len(headers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			tx := transactions[id.Row]
			label := cell.(*widget.Label)
			var value string
			switch id.Col {
			case 0:
				value = tx.Amount.StringFixed(2)
			case 1:
				value = tx.Date.Format("2006-01-02")
			case 2:
				value = tx.Type.Name
			case 3:
				value = tx.Currency.Code
			case 4:
				value = tx.Category.Name
			case 5:
				value = tx.SubCategory.Name
			case 6:
				value = tx.Description
			}
			label.SetText(value)
		},
	)

	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 120)
	table.SetColumnWidth(5, 200)
	table.SetColumnWidth(6, 400)

	btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)
	btnRefresh := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameViewRefresh), func() {
		g.listWindow.SetContent(makeListContent(g, nil))
	})
	btnFilter := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameSearch), g.ShowFilterScreen)
	ctrBtns := container.NewBorder(nil, nil, btnFilter, btnRefresh, btnReturn)

	content := container.NewBorder(lblTitle, ctrBtns, nil, nil, container.NewBorder(customHeader, nil, nil, nil, table))

	return content
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
