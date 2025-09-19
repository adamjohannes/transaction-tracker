package gui

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/chart"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/shopspring/decimal"
)

func makeDataContent(g *GUI) fyne.CanvasObject {
	title := widget.NewLabel("Transaction Data")
	viewContainer := container.NewMax()

	// Combo box for switching between categories and sub categories
	viewSelector := widget.NewSelect([]string{"Category", "Sub-Category"}, func(selected string) {
		var newView fyne.CanvasObject
		if selected == "Category" {
			newView = buildCategoryView(g)
		} else {
			newView = buildSubCategoryView(g)
		}
		viewContainer.Objects = []fyne.CanvasObject{newView}
		viewContainer.Refresh()
	})

	topContent := container.NewBorder(nil, nil, title, viewSelector)
	viewSelector.SetSelected("Category")

	return container.NewBorder(topContent, nil, nil, nil, viewContainer)
}

func buildCategoryView(g *GUI) fyne.CanvasObject {
	transactionCounts, err := g.transactionController.GetTransactionCount("category")
	if err != nil {
		log.Printf("Error getting transaction counts: %v", err)
		return container.NewCenter(widget.NewLabel("Error loading data"))
	}

	return container.NewAppTabs(
		container.NewTabItem("Debit", createCountPieChart("Debit by Category", transactionCounts["Debit"])),
		container.NewTabItem("Credit", createCountPieChart("Credit by Category", transactionCounts["Credit"])),
		container.NewTabItem("Refund", createCountPieChart("Refund by Category", transactionCounts["Refund"])),
	)
}

func buildSubCategoryView(g *GUI) fyne.CanvasObject {
	allData, err := g.transactionController.GetSubCategoryAmounts()

	if err != nil {
		log.Printf("Error getting sub-category amounts: %v", err)
		return container.NewCenter(widget.NewLabel("Error loading data"))
	}

	categoryTabs := container.NewAppTabs()
	typeSelector := widget.NewSelect([]string{"Debit", "Credit", "Refund"}, func(selectedType string) {
		dataForType := make(map[string]map[string]decimal.Decimal)
		categories := make([]string, 0)

		for _, item := range allData {
			if item.TransactionType == selectedType {
				if _, ok := dataForType[item.CategoryName]; !ok {
					dataForType[item.CategoryName] = make(map[string]decimal.Decimal)
					categories = append(categories, item.CategoryName)
				}
				dataForType[item.CategoryName][item.SubCategoryName] = item.TotalAmount
			}
		}
		sort.Strings(categories)

		var newTabs []*container.TabItem

		if len(categories) > 0 {
			for _, categoryName := range categories {
				chartTitle := fmt.Sprintf("%s Breakdown", categoryName)
				tabContent := createAmountPieChart(chartTitle, dataForType[categoryName])
				newTabs = append(newTabs, container.NewTabItem(categoryName, tabContent))
			}
		} else {
			noDataLabel := container.NewCenter(widget.NewLabel("No sub-category data for this transaction type."))
			newTabs = append(newTabs, container.NewTabItem("No Data", noDataLabel))
		}

		categoryTabs.SetItems(newTabs)
	})

	typeSelector.SetSelected("Debit")

	return container.NewBorder(container.NewVBox(widget.NewLabel("Grouped by Amount"), typeSelector), nil, nil, nil, categoryTabs)
}

func createCountPieChart(title string, dataMap map[string]int) fyne.CanvasObject {
	if len(dataMap) == 0 {
		return container.NewCenter(widget.NewLabel("No data available."))
	}
	var chartData []data.ProportionalDataPoint
	for name, count := range dataMap {
		chartData = append(chartData, data.ProportionalDataPoint{C: name, Val: float64(count), Col: randomColor()})
	}
	return buildChart(title, chartData)
}

func createAmountPieChart(title string, dataMap map[string]decimal.Decimal) fyne.CanvasObject {
	if len(dataMap) == 0 {
		return container.NewCenter(widget.NewLabel("No data available."))
	}
	var chartData []data.ProportionalDataPoint
	for name, amount := range dataMap {
		floatAmount, _ := amount.Float64()
		chartData = append(chartData, data.ProportionalDataPoint{C: name, Val: floatAmount, Col: randomColor()})
	}
	return buildChart(title, chartData)
}

func buildChart(title string, chartData []data.ProportionalDataPoint) fyne.CanvasObject {
	pieChart := chart.NewPolarProportionalChart()
	pieChart.SetTitle(title)
	_, err := pieChart.AddProportionalSeries(title, chartData)
	if err != nil {
		log.Printf("Error adding data to pie chart: %v", err)
	}
	return pieChart
}

func randomColor() color.Color {
	return color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
}
