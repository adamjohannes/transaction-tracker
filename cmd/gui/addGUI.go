package gui

import (
	"fmt"
	"log"
	"monthly-expenses-handler/internal/config"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addGUI(a fyne.App, config *config.GUIConfig) fyne.Window {
	w := a.NewWindow("Adicionar Transação")
	w.Resize(fyne.NewSize(config.Width, config.Height))
	w.SetFixedSize(!config.Resizable)

	// Build temp transaction values
	tempTransaction := make(map[string]any)
	tempTransaction["amount"] = nil
	tempTransaction["date"] = nil
	tempTransaction["essential"] = false
	tempTransaction["status"] = nil
	tempTransaction["currency"] = nil
	tempTransaction["category"] = nil

	// 1. Amount *
	ctrAmount := amount(tempTransaction)

	// 2. Date *
	ctrDate := date(tempTransaction)

	// 3. Essential *
	ctrEssential := essential(tempTransaction)

	// 4. Status *
	ctrStatus := buildRadioGroupAccordion("Status (Required)", float32(120), tempTransaction, []string{"Pending", "Completed", "Failed"})

	// 5. Currency *
	ctrCurrency := buildRadioGroupAccordion("Currency (Required)", float32(120), tempTransaction, []string{"BRL", "USD", "EUR"})

	// 6. Category *
	ctrCategory := buildRadioGroupAccordion("Category (Required)", float32(200), tempTransaction, []string{
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

	// 7. Subcategory
	ctrSubCategory := buildRadioGroupAccordion("Sub Category (Required)", float32(200), tempTransaction, []string{
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

	// 8. Description
	ctrDescription := description(tempTransaction)

	formContent := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: ctrAmount},
			{Text: "", Widget: ctrDate},
			{Text: "", Widget: ctrEssential},
			{Text: "", Widget: ctrStatus},
			{Text: "", Widget: ctrCurrency},
			{Text: "", Widget: ctrCategory},
			{Text: "", Widget: ctrSubCategory},
			{Text: "", Widget: ctrDescription},
		},
		OnSubmit: func() {
			log.Println(fmt.Sprintf("Transaction amount: %v", tempTransaction["amount"]))
			log.Println(fmt.Sprintf("Transaction date: %v", tempTransaction["date"]))
			log.Println(fmt.Sprintf("Transaction essential: %v", tempTransaction["essential"]))
			log.Println(fmt.Sprintf("Transaction status: %v", tempTransaction["status"]))
			log.Println(fmt.Sprintf("Transaction currency: %v", tempTransaction["currency"]))
			log.Println(fmt.Sprintf("Transaction category: %v", tempTransaction["category"]))
			log.Println(fmt.Sprintf("Transaction subCategory: %v", tempTransaction["subCategory"]))
			log.Println(fmt.Sprintf("Transaction description: %v", tempTransaction["description"]))
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

func amount(transaction map[string]any) fyne.CanvasObject {
	lblAmount := widget.NewLabel("Amount: (Required)")
	entryAmount := widget.NewEntry()
	entryAmount.OnChanged = func(s string) {
		transaction["amount"] = s
	}
	return container.NewBorder(nil, nil, lblAmount, nil, entryAmount)
}

func date(transaction map[string]any) fyne.CanvasObject {
	lblDate := widget.NewLabel("Date: (Required)")
	entryDate := widget.NewDateEntry()
	entryDate.OnChanged = func(time *time.Time) {
		transaction["date"] = time.Format("2006-01-02")
	}
	return container.NewBorder(nil, nil, lblDate, nil, entryDate)
}

func essential(transaction map[string]any) fyne.CanvasObject {
	lblEssential := widget.NewLabel("Essential:")
	checkEssential := widget.NewCheck("", func(checked bool) {
		transaction["essential"] = checked
	})
	return container.NewBorder(nil, nil, lblEssential, nil, checkEssential)
}

func description(transaction map[string]any) fyne.CanvasObject {
	lblDescription := widget.NewLabel("Description:")
	entryDescription := widget.NewEntry()
	entryDescription.OnChanged = func(s string) {
		transaction["description"] = s
	}
	return container.NewBorder(nil, nil, lblDescription, nil, entryDescription)
}

func buildRadioGroupAccordion(lbl string, minHeight float32, transaction map[string]any, options []string) fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		options,
		func(s string) {
			transaction[s] = s
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
