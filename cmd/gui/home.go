package gui

import (
	"log"
	"monthly-expenses-handler/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	window fyne.Window
}

func NewGUI(config *config.GUIConfig) *GUI {
	a := app.New()

	addGui := addGUI(a, config)
	home := homeGUI(a, addGui, config)

	return &GUI{home}
}

func homeGUI(a fyne.App, addGui fyne.Window, config *config.GUIConfig) fyne.Window {
	w := a.NewWindow("Rastreador de Transações")
	w.Resize(fyne.NewSize(config.Width, config.Height))
	w.SetFixedSize(!config.Resizable)

	lblTitle := widget.NewLabel("Rastreador de Transações")
	lblTitle.Alignment = fyne.TextAlignCenter
	lblTitle.TextStyle.Bold = true

	addBtn := widget.NewButton("Adicionar Transação", func() {
		addGui.Show()
		w.Hide()
	})

	listBtn := widget.NewButton("Listar Transações", func() {
		log.Println("List")
	})

	ctrBtns := container.NewGridWithRows(2, addBtn, listBtn)

	ctr := container.NewBorder(lblTitle, nil, nil, nil, ctrBtns)

	w.SetContent(ctr)
	return w
}

func (g *GUI) Start() {
	g.window.ShowAndRun()
}
