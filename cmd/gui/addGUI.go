package gui

import (
	"fmt"
	"log"
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

	// 2. Date *
	ctrDate := date()

	// 3. Status *
	ctrStatus := status()

	// 4. Currency *
	ctrCurrency := currency()

	// 5. Category *
	ctrCategory := category()

	// 6. Subcategory
	ctrSubCategory := subCategory()

	// 7. Description
	ctrDescription := description()

	formContent := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: ctrAmount},
			{Text: "", Widget: ctrDate},
			{Text: "", Widget: ctrStatus},
			{Text: "", Widget: ctrCurrency},
			{Text: "", Widget: ctrCategory},
			{Text: "", Widget: ctrSubCategory},
			{Text: "", Widget: ctrDescription},
		},
		OnSubmit: func() {
			log.Println("Form submited.")
		},
		OnCancel: func() {
			log.Println("Form canceled")
		},
		SubmitText: "Save",
		CancelText: "Cancel",
	}

	scrollCtr := container.NewScroll(formContent)
	w.SetContent(scrollCtr)
	return w
}

func amount() fyne.CanvasObject {
	lblAmount := widget.NewLabel("Amount: (Required)")
	entryAmount := widget.NewEntry()
	return container.NewBorder(nil, nil, lblAmount, nil, entryAmount)
}

func date() fyne.CanvasObject {
	lblDate := widget.NewLabel("Date: (Required)")
	entryDate := widget.NewDateEntry()
	return container.NewBorder(nil, nil, lblDate, nil, entryDate)
}

func status() fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		[]string{
			"Pending",
			"Completed",
			"Failed",
		},
		func(s string) {
			fmt.Println(s)
		},
	)

	ctrItems := container.NewScroll(
		radioGroup,
	)
	ctrItems.SetMinSize(fyne.NewSize(radioGroup.MinSize().Width, 200))

	accStatus := widget.NewAccordion(
		widget.NewAccordionItem("Status: (Required)",
			ctrItems,
		),
	)

	return accStatus
}

func currency() fyne.CanvasObject {
	lblCurrency := widget.NewLabel("Currency: (Required)")
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

func category() fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		[]string{
			"Category 1",
			"Category 2",
			"Category 3",
			"Category 4",
			"Category 5",
			"Category 6",
			"Category 7",
			"Category 8",
			"Category 9",
			"Category 10",
		},
		func(s string) {
			fmt.Println(s)
		},
	)
	radioGroup.Required = true

	ctrItems := container.NewScroll(
		radioGroup,
	)
	ctrItems.SetMinSize(fyne.NewSize(radioGroup.MinSize().Width, 200))

	accCategory := widget.NewAccordion(
		widget.NewAccordionItem("Category: (Required)",
			ctrItems,
		),
	)

	return accCategory
}

func subCategory() fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		[]string{
			"Sub Category 1",
			"Sub Category 2",
			"Sub Category 3",
			"Sub Category 4",
			"Sub Category 5",
			"Sub Category 6",
			"Sub Category 7",
			"Sub Category 8",
			"Sub Category 9",
			"Sub Category 10",
		},
		func(s string) {
			fmt.Println(s)
		},
	)

	ctrItems := container.NewScroll(
		radioGroup,
	)
	ctrItems.SetMinSize(fyne.NewSize(radioGroup.MinSize().Width, 200))

	accSubCategory := widget.NewAccordion(
		widget.NewAccordionItem("Sub Category:",
			ctrItems,
		),
	)

	return accSubCategory
}

func description() fyne.CanvasObject {
	lblDescription := widget.NewLabel("Description:")
	entryDescription := widget.NewEntry()
	return container.NewBorder(nil, nil, lblDescription, nil, entryDescription)
}
