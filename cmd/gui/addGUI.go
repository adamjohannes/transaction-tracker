package gui

import (
	"log"
	"monthly-expenses-handler/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
	ctrStatus := buildRadioGroupAccordion("Status (Required)", float32(120), []string{"Pending", "Completed", "Failed"})

	// 4. Currency *
	ctrCurrency := buildRadioGroupAccordion("Currency (Required)", float32(120), []string{"BRL", "USD", "EUR"})

	// 5. Category *
	ctrCategory := buildRadioGroupAccordion("Category (Required)", float32(200), []string{
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
	})

	// 6. Subcategory
	ctrSubCategory := buildRadioGroupAccordion("Sub Category (Required)", float32(200), []string{
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
	})

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

func description() fyne.CanvasObject {
	lblDescription := widget.NewLabel("Description:")
	entryDescription := widget.NewEntry()
	return container.NewBorder(nil, nil, lblDescription, nil, entryDescription)
}

func buildRadioGroupAccordion(lbl string, minHeight float32, options []string) fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		options,
		func(s string) {
			log.Println(s)
		},
	)

	ctrItems := container.NewScroll(
		radioGroup,
	)
	ctrItems.SetMinSize(fyne.NewSize(radioGroup.MinSize().Width, minHeight))

	acc := widget.NewAccordion(
		widget.NewAccordionItem(lbl, ctrItems),
	)

	return acc
}
