package gui

import (
	"log"
	"time"

	domain "monthly-expenses-handler/internal/domain/transaction"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func makeFilterContent(g *GUI) fyne.CanvasObject {
	filters := &domain.FilterCriteria{}

	title := widget.NewLabel("Filter Transactions")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true

	subCategorySelect := widget.NewSelect([]string{}, func(s string) {
		if s == "Any" || s == "" {
			filters.SubCategoryName = nil
		} else {
			filters.SubCategoryName = &s
		}
	})
	subCategorySelect.PlaceHolder = "Select a category first"
	subCategorySelect.Disable()

	categories, err := g.categoryController.GetAllCategories()
	if err != nil {
		log.Printf("Failed to load categories for filter screen: %v", err)
	}
	categoryNames := []string{"Any"}
	for _, cat := range categories {
		categoryNames = append(categoryNames, cat.Name)
	}

	categorySelect := widget.NewSelect(categoryNames, func(s string) {
		// Reset sub-category on change
		filters.SubCategoryName = nil
		subCategorySelect.ClearSelected()

		if s == "Any" {
			filters.CategoryName = nil
			subCategorySelect.Disable()
			subCategorySelect.PlaceHolder = "Select a category first"
		} else {
			filters.CategoryName = &s
			subCategories, err := g.subCategoryController.GetSubCategoriesByCategory(s)
			if err != nil {
				log.Printf("Failed to load sub-categories: %v", err)
				subCategorySelect.Disable()
				return
			}
			subCategoryNames := []string{"Any"}
			for _, subCat := range subCategories {
				subCategoryNames = append(subCategoryNames, subCat.Name)
			}
			subCategorySelect.Options = subCategoryNames
			subCategorySelect.Enable()
			subCategorySelect.PlaceHolder = "Any Sub Category"
		}
		subCategorySelect.Refresh()
	})

	startDateEntry := widget.NewEntry()
	startDateEntry.SetPlaceHolder("YYYY-MM-DD")
	endDateEntry := widget.NewEntry()
	endDateEntry.SetPlaceHolder("YYYY-MM-DD")

	essentialRadio := widget.NewRadioGroup([]string{"Any", "Yes", "No"}, func(s string) {
		switch s {
		case "Yes":
			isEssential := true
			filters.Essential = &isEssential
		case "No":
			isEssential := false
			filters.Essential = &isEssential
		default: // "Any"
			filters.Essential = nil
		}
	})
	essentialRadio.SetSelected("Any")
	essentialRadio.Horizontal = true

	typeSelect := widget.NewSelect([]string{"Any", "Debit", "Credit", "Refund"}, func(s string) {
		if s == "Any" {
			filters.TypeName = nil
		} else {
			filters.TypeName = &s
		}
	})

	statusSelect := widget.NewSelect([]string{"Any", "Pending", "Completed", "Failed"}, func(s string) {
		if s == "Any" {
			filters.StatusName = nil
		} else {
			filters.StatusName = &s
		}
	})

	currencySelect := widget.NewSelect([]string{"Any", "BRL", "USD", "EUR"}, func(s string) {
		if s == "Any" {
			filters.CurrencyCode = nil
		} else {
			filters.CurrencyCode = &s
		}
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Category", Widget: categorySelect},
			{Text: "Sub Category", Widget: subCategorySelect},
			{Text: "Start Date", Widget: startDateEntry},
			{Text: "End Date", Widget: endDateEntry},
			{Text: "Essential", Widget: essentialRadio},
			{Text: "Type", Widget: typeSelect},
			{Text: "Status", Widget: statusSelect},
			{Text: "Currency", Widget: currencySelect},
		},
	}

	applyButton := widget.NewButton("Apply Filters", func() {
		log.Printf("Applying filters: %+v", filters)

		// Parse dates from entries, ignoring errors (blank/invalid means no filter)
		if t, err := time.Parse("2006-01-02", startDateEntry.Text); err == nil {
			filters.StartDate = &t
		} else {
			filters.StartDate = nil
		}
		if t, err := time.Parse("2006-01-02", endDateEntry.Text); err == nil {
			filters.EndDate = &t
		} else {
			filters.EndDate = nil
		}

		transactions, err := g.transactionController.GetFilteredTransactions(filters)
		if err != nil {
			log.Printf("Error applying filters: %v", err)
			dialog.ShowError(err, g.filterWindow)
			return
		}

		log.Printf("Found %d transactions matching filters.", len(transactions))
		g.listWindow.SetContent(makeListContent(g, transactions))
		g.ShowListScreen(false)
	})

	cancelButton := widget.NewButton("Cancel", func() {
		g.ShowListScreen(false)
	})

	buttons := container.NewGridWithColumns(2, cancelButton, applyButton)
	return container.NewBorder(title, buttons, nil, nil, container.NewScroll(form))
}
