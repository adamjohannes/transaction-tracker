package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeListContent(g *GUI) fyne.CanvasObject {
	lblTitle := widget.NewLabel("Lista das Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true

	btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)
	btnRefresh := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameViewRefresh), func() {
		g.listWindow.SetContent(makeListContent(g))
	})
	btnFilter := widget.NewButtonWithIcon("", theme.Icon(theme.IconNameSearch), func() {
		log.Println("user wants to filter transactions")
	})
	ctrBtns := container.NewBorder(nil, nil, btnFilter, btnRefresh, btnReturn)

	transactions, err := g.transactionController.GetAllTransactions()
	if err != nil {
		go func() {
			dialog.ShowError(fmt.Errorf("failed to load transactions: %w", err), g.listWindow)
		}()
		return container.NewBorder(lblTitle, btnReturn, nil, nil)
	}

	// Convert the slice of Transaction structs into a 2D slice of strings for the table
	var tableData [][]string
	for _, tx := range transactions {
		row := []string{
			tx.Amount.StringFixed(2), // Format amount to 2 decimal places
			tx.Date.Format("2006-01-02"),
			tx.Type.Name,
			tx.Currency.Code,
			tx.Category.Name,
			tx.SubCategory.Name,
			tx.Description,
		}
		tableData = append(tableData, row)
	}

	headers := []string{
		"Amount", "Date", "Type", "Currency",
		"Category", "Sub Category", "Description",
	}

	table := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(tableData), len(headers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText(tableData[id.Row][id.Col]) // Populate cells with fetched data
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

	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 100)
	table.SetColumnWidth(4, 120)
	table.SetColumnWidth(5, 200)
	table.SetColumnWidth(6, 200)

	table.ShowHeaderColumn = false

	return container.NewBorder(lblTitle, ctrBtns, nil, nil, table)
}
