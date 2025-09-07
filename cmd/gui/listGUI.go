package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeListContent(g *GUI) fyne.CanvasObject {
	lblTitle := widget.NewLabel("Lista das Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true
	btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)

	sampleData := [][]string{
		{"100.00", "2025-09-07", "Expense", "USD", "Food", "Groceries", "Bought vegetables"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
		{"250.00", "2025-09-06", "Expense", "EUR", "Transport", "Fuel", "Gas for car"},
		{"500.00", "2025-09-05", "Income", "USD", "Salary", "Main Job", "September paycheck"},
		{"75.50", "2025-09-04", "Expense", "GBP", "Entertainment", "Movies", "Cinema tickets"},
		{"1200.00", "2025-09-01", "Income", "USD", "Freelance", "Project", "Website development"},
	}

	headers := []string{
		"Amount",
		"Date",
		"Type",
		"Currency",
		"Category",
		"Sub Category",
		"Description",
	}

	table := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(sampleData), len(headers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText(sampleData[id.Row][id.Col])
		},
	)

	table.CreateHeader = func() fyne.CanvasObject {
		return widget.NewLabelWithStyle("Header", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	}

	table.UpdateHeader = func(id widget.TableCellID, cell fyne.CanvasObject) {
		label := cell.(*widget.Label)
		if id.Row == -1 && id.Col >= 0 && id.Col < len(headers) {
			label.SetText(headers[id.Col])
		}
	}

	// Set column widths
	table.SetColumnWidth(0, 100) // Amount
	table.SetColumnWidth(1, 120) // Date
	table.SetColumnWidth(2, 100) // Type
	table.SetColumnWidth(3, 100) // Currency
	table.SetColumnWidth(4, 120) // Category
	table.SetColumnWidth(5, 150) // Sub Category
	table.SetColumnWidth(6, 200) // Description

	// Disable table index header
	table.ShowHeaderColumn = false

	return container.NewBorder(lblTitle, btnReturn, nil, nil, table)
}
