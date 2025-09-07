package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeListContent(g *GUI) fyne.CanvasObject {
	lblTitle := widget.NewLabel("Lista das Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true
	btnReturn := widget.NewButton("Voltar", g.ShowHomeScreen)

	// - Max number of columns is 10 (9 fields plus the index column)
	//   - The index column is NOT the row ID in the DB, its just the
	//     number of the row in the current view

	numCols := 10
	numRows := 20
	table := widget.NewTable(
		func() (int, int) {
			return numRows, numCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				switch id.Col {
				case 0:
					label.SetText("ID")
				case 1:
					label.SetText("Amount")
				case 2:
					label.SetText("Date")
				case 3:
					label.SetText("Essential")
				case 4:
					label.SetText("Type")
				case 5:
					label.SetText("Status")
				case 6:
					label.SetText("Currency")
				case 7:
					label.SetText("Category")
				case 8:
					label.SetText("SubCategory")
				case 9:
					label.SetText("Description")
				case 10:
					label.SetText("Description")
				}
				return
			}
			label.SetText(fmt.Sprintf("C%d R%d", id.Col+1, id.Row+1))
		},
	)

	for i := 0; i < numCols; i++ {
		table.SetColumnWidth(i, 120)
	}

	return container.NewBorder(lblTitle, btnReturn, nil, nil, table)
}
