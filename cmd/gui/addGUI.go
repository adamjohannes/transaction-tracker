package gui

import (
	"monthly-expenses-handler/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func addGUI(a fyne.App, config *config.GUIConfig) fyne.Window {
	w := a.NewWindow("Adicionar Transação")
	w.Resize(fyne.NewSize(config.Width, config.Height))
	w.SetFixedSize(!config.Resizable)

	// 1. Amount *
	ctrAmount := amount()

	// 2. Category *
	ctrCategory := category()

	// 3. Subcategory
	ctrSubCategory := subCategory()

	// 4. Date *
	ctrDate := date()

	// 5. Description
	ctrDescription := description()

	// 6. Status *
	ctrStatus := status()

	// 7. Currency *
	ctrCurrency := currency()

	content := container.NewVBox(
		ctrAmount, widget.NewSeparator(),
		ctrCategory, widget.NewSeparator(),
		ctrSubCategory, widget.NewSeparator(),
		ctrDate, widget.NewSeparator(),
		ctrDescription, widget.NewSeparator(),
		ctrStatus, widget.NewSeparator(),
		ctrCurrency)

	scrollCtr := container.NewScroll(content)
	w.SetContent(scrollCtr)
	return w
}

func amount() fyne.CanvasObject {
	lblAmount := widget.NewLabel("Amout:")
	entryAmount := widget.NewEntry()
	return container.NewBorder(nil, nil, lblAmount, nil, entryAmount)
}

func category() fyne.CanvasObject {
	lblCategory := widget.NewLabel("Category:")
	entryCategory := widget.NewEntry()
	return container.NewBorder(nil, nil, lblCategory, nil, entryCategory)
}

func subCategory() fyne.CanvasObject {
	lblSubCategory := widget.NewLabel("Subcategory:")
	entrySubCategory := widget.NewEntry()
	return container.NewBorder(nil, nil, lblSubCategory, nil, entrySubCategory)
}

func date() fyne.CanvasObject {
	lblDate := widget.NewLabel("Date:")
	entryDate := widget.NewDateEntry()
	return container.NewBorder(nil, nil, lblDate, nil, entryDate)
}

func description() fyne.CanvasObject {
	lblDescription := widget.NewLabel("Description:")
	entryDescription := widget.NewEntry()
	return container.NewBorder(nil, nil, lblDescription, nil, entryDescription)
}

func status() fyne.CanvasObject {
	lblStatus := widget.NewLabel("Status:")
	errorText := canvas.NewText("Please select only one status.", theme.ErrorColor())
	errorText.Hide()

	checkStatus := widget.NewCheckGroup([]string{"Pending", "Completed", "Failed"}, func(selected []string) {
		if len(selected) > 1 {
			errorText.Show()
		} else {
			errorText.Hide()
		}
	})

	inputContainer := container.NewVBox(checkStatus, errorText)
	return container.NewBorder(nil, nil, lblStatus, nil, inputContainer)
}

func currency() fyne.CanvasObject {
	lblCurrency := widget.NewLabel("Currency:")
	errorText := canvas.NewText("Please select only one currency.", theme.ErrorColor())
	errorText.Hide()

	checkCurrency := widget.NewCheckGroup([]string{"BRL", "USD", "EUR"}, func(selected []string) {
		if len(selected) > 1 {
			errorText.Show()
		} else {
			errorText.Hide()
		}
	})

	// Use a VBox to stack the input widget and its potential error message
	inputContainer := container.NewVBox(checkCurrency, errorText)
	return container.NewBorder(nil, nil, lblCurrency, nil, inputContainer)
}
