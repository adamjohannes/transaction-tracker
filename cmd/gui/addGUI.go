package gui

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// makeAddContent creates the UI for the "add transaction" screen.
// It takes the GUI controller to link the Cancel button to the navigation method.
func makeAddContent(g *GUI) fyne.CanvasObject {
	// Build temp transaction values
	tempTransaction := make(map[string]any)
	resetTransaction(tempTransaction)

	// Fetch categories
	categories, err := g.categoryController.GetAllCategories()

	if err != nil {
		log.Println(fmt.Sprintf("Failed to load categories: %v", err))
		dialog.ShowError(err, g.addWindow)
	}

	// Convert the list of Category objects into a simple list of names for the widget.
	categoryNames := make([]string, len(categories))
	for i, cat := range categories {
		categoryNames[i] = cat.Name
	}

	// Build form widgets
	ctrAmount := buildEntryCtr("Amount (Required)", "amount", tempTransaction)
	ctrDate := buildDateCtr("Date (Required)", "date", tempTransaction)
	ctrEssential := buildCheckCtr("Essential", "essential", tempTransaction)
	ctrType := buildRadioGroupAccordion("Type (Required)", "type", float32(120), tempTransaction, []string{"Debit", "Credit", "Refund"})
	ctrStatus := buildRadioGroupAccordion("Status (Required)", "status", float32(120), tempTransaction, []string{"Pending", "Completed", "Failed"})
	ctrCurrency := buildRadioGroupAccordion("Currency (Required)", "currency", float32(120), tempTransaction, []string{"BRL", "USD", "EUR"})
	ctrCategory := buildRadioGroupAccordion("Category (Required)", "category", float32(200), tempTransaction, categoryNames)
	ctrSubCategory := buildRadioGroupAccordion("Sub Category (Required)", "subCategory", float32(200), tempTransaction, []string{"Eletrônicos", "Sub Category 2", "Sub Category 3"})
	ctrDescription := buildEntryCtr("Description", "description", tempTransaction)

	// Build form
	formContent := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: ctrAmount},
			{Text: "", Widget: ctrDate},
			{Text: "", Widget: ctrEssential},
			{Text: "", Widget: ctrType},
			{Text: "", Widget: ctrStatus},
			{Text: "", Widget: ctrCurrency},
			{Text: "", Widget: ctrCategory},
			{Text: "", Widget: ctrSubCategory},
			{Text: "", Widget: ctrDescription},
		},
	}

	// Build buttons
	saveButton := widget.NewButton("Save", func() {
		log.Printf("Attempting to save transaction: %v", tempTransaction)
		_, err := g.transactionController.NewTransaction(tempTransaction)
		if err != nil {
			log.Printf("Error saving transaction: %v", err)
			dialog.ShowError(fmt.Errorf("failed to save transaction: %w", err), g.addWindow)
			return
		}

		log.Println("Transaction saved successfully!")
		dialog.ShowInformation("Success", "Transaction saved successfully!", g.addWindow)
		g.ShowHomeScreen()
	})

	cancelButton := widget.NewButton("Cancel", g.ShowHomeScreen)

	formContent.SetOnValidationChanged(func(err error) {
		if err == nil {
			saveButton.Enable()
		} else {
			saveButton.Disable()
		}
	})

	if formContent.Validate() != nil {
		saveButton.Disable()
	}

	buttons := container.NewGridWithColumns(2, cancelButton, saveButton)
	scrollCtr := container.NewScroll(formContent)
	title := widget.NewLabel("Adicionar Transação")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	return container.NewBorder(title, buttons, nil, nil, scrollCtr)
}

// --- Helper functions for creating form widgets ---

func resetTransaction(transaction map[string]any) {
	transaction["amount"] = nil
	transaction["date"] = nil
	transaction["type"] = nil
	transaction["essential"] = false
	transaction["status"] = nil
	transaction["currency"] = nil
	transaction["category"] = nil
	transaction["subCategory"] = nil
	transaction["description"] = nil
}

func buildEntryCtr(lbl, dbEntry string, transaction map[string]any) fyne.CanvasObject {
	lblCtr := widget.NewLabel(lbl)
	entryDescription := widget.NewEntry()
	entryDescription.OnChanged = func(s string) {
		transaction[dbEntry] = s
	}
	return container.NewBorder(nil, nil, lblCtr, nil, entryDescription)
}

func buildCheckCtr(lbl, dbEntry string, transaction map[string]any) fyne.CanvasObject {
	lblCtr := widget.NewLabel(lbl)
	checkEssential := widget.NewCheck("", func(checked bool) {
		transaction[dbEntry] = checked
	})
	checkEssential.SetChecked(false)
	transaction[dbEntry] = false
	return container.NewBorder(nil, nil, lblCtr, nil, checkEssential)
}

func buildDateCtr(lbl, dbEntry string, transaction map[string]any) fyne.CanvasObject {
	lblDate := widget.NewLabel(lbl)
	entryDate := widget.NewDateEntry()
	entryDate.OnChanged = func(t *time.Time) {
		transaction[dbEntry] = t.Format("2006-01-02")
	}
	return container.NewBorder(nil, nil, lblDate, nil, entryDate)
}

func buildRadioGroupAccordion(lbl, dbEntry string, minHeight float32, transaction map[string]any, options []string) fyne.CanvasObject {
	radioGroup := widget.NewRadioGroup(
		options,
		func(s string) {
			transaction[dbEntry] = s
		},
	)
	ctrItems := container.NewScroll(radioGroup)
	ctrItems.SetMinSize(fyne.NewSize(radioGroup.MinSize().Width, minHeight))
	acc := widget.NewAccordion(
		widget.NewAccordionItem(lbl, ctrItems),
	)
	return acc
}
