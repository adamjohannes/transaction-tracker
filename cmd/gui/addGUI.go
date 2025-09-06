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

// makeAddContent creates the UI for the "add transaction" screen.
// It takes the GUI controller to link the Cancel button to the navigation method.
func makeAddContent(g *GUI, config *config.GUIConfig) fyne.CanvasObject {
	// Build temp transaction values
	tempTransaction := make(map[string]any)
	tempTransaction["amount"] = nil
	tempTransaction["date"] = nil
	tempTransaction["essential"] = false
	tempTransaction["status"] = nil
	tempTransaction["currency"] = nil
	tempTransaction["category"] = nil

	// Build form widgets
	ctrAmount := amount(tempTransaction)
	ctrDate := date(tempTransaction)
	ctrEssential := essential(tempTransaction)
	ctrStatus := buildRadioGroupAccordion("Status (Required)", "status", float32(120), tempTransaction, []string{"Pending", "Completed", "Failed"})
	ctrCurrency := buildRadioGroupAccordion("Currency (Required)", "currency", float32(120), tempTransaction, []string{"BRL", "USD", "EUR"})
	ctrCategory := buildRadioGroupAccordion("Category (Required)", "category", float32(200), tempTransaction, []string{"Category 1", "Category 2", "Category 3", "Category 4", "Category 5"})
	ctrSubCategory := buildRadioGroupAccordion("Sub Category (Required)", "subCategory", float32(200), tempTransaction, []string{"Sub Category 1", "Sub Category 2", "Sub Category 3"})
	ctrDescription := description(tempTransaction)

	// Build form
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
	}

	// Build buttons
	saveButton := widget.NewButton("Save", func() {
		log.Println(fmt.Sprintf("Transaction to save: %v", tempTransaction))
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
	return container.NewBorder(nil, buttons, nil, nil, scrollCtr)
}

// --- Helper functions for creating form widgets ---

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
	entryDate.OnChanged = func(t *time.Time) {
		transaction["date"] = t.Format("2006-01-02")
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

func buildRadioGroupAccordion(lbl string, dbEntry string, minHeight float32, transaction map[string]any, options []string) fyne.CanvasObject {
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
