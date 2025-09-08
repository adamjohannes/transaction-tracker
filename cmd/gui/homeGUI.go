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

	listBtn := widget.NewButton("Listar Transações", func() {
		g.ShowListScreen(true)
	})

	dataBtn := widget.NewButton("Visualizar Dados", g.ShowDataScreen)
	ctrBtns := container.NewGridWithRows(3, addBtn, listBtn, dataBtn)

	return container.NewBorder(lblTitle, nil, nil, nil, ctrBtns)
}
