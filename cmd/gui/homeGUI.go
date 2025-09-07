package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// makeHomeContent creates the UI for the home screen.
// It takes the GUI controller to link button actions to navigation methods.
func makeHomeContent(g *GUI) fyne.CanvasObject {
	lblTitle := widget.NewLabel("Rastreador de Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true

	addBtn := widget.NewButton("Adicionar Transação", g.ShowAddScreen)

	listBtn := widget.NewButton("Listar Transações", g.ShowListScreen)

	ctrBtns := container.NewGridWithRows(2, addBtn, listBtn)

	return container.NewBorder(lblTitle, nil, nil, nil, ctrBtns)
}
