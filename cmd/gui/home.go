package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	window fyne.Window
}

func NewGUI() *GUI {
	a := app.New()
	w := a.NewWindow("Fyne UI")
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(true)

	addBtn := widget.NewButton("Add", func() {
		fmt.Println("Add")
	})

	listBtn := widget.NewButton("List", func() {
		fmt.Println("List")
	})

	ctr := container.NewGridWithRows(2, addBtn, listBtn)
	w.SetContent(ctr)
	return &GUI{w}
}

func (g *GUI) Start() {
	g.window.ShowAndRun()
}
