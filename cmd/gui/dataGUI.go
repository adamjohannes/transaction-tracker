package gui

import (
	"image/color"
	"log"
	"math/rand"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/chart"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

func makeDataContent(g *GUI) fyne.CanvasObject {
	title := widget.NewLabel("Transaction Data")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true

	// A container to hold the charts, allowing us to refresh it later
	chartContainer := container.NewMax()

	// The callback function that will run when the combo box selection changes
	updateCharts := func(selected string) {
		// Convert "Sub-Category" to "sub_category" for the backend
		groupByKey := strings.ToLower(strings.Replace(selected, "-", "_", 1))

		transactionCounts, err := g.transactionController.GetTransactionCount(groupByKey)
		if err != nil {
			log.Printf("Error getting transaction counts: %v", err)
			dialog.ShowError(err, g.dataWindow)
			chartContainer.Objects = []fyne.CanvasObject{widget.NewLabel("Error loading data")}
			chartContainer.Refresh()
			return
		}

		tabs := container.NewAppTabs(
			container.NewTabItem("Debit", createPieChart("Debit by "+selected, transactionCounts["Debit"])),
			container.NewTabItem("Credit", createPieChart("Credit by "+selected, transactionCounts["Credit"])),
			container.NewTabItem("Refund", createPieChart("Refund by "+selected, transactionCounts["Refund"])),
		)

		// Replace the content of the chart container and refresh it
		chartContainer.Objects = []fyne.CanvasObject{tabs}
		chartContainer.Refresh()
	}

	// Create the combo box (Select widget)
	groupBySelect := widget.NewSelect([]string{"Category", "Sub-Category"}, updateCharts)

	// Set the initial value, which will trigger the first call to updateCharts
	groupBySelect.SetSelected("Category")

	// Create a top bar with the title and the select widget
	topContent := container.NewBorder(nil, nil, title, groupBySelect)

	return container.NewBorder(topContent, nil, nil, nil, chartContainer)
}

// This helper function remains unchanged
func createPieChart(title string, dataMap map[string]int) fyne.CanvasObject {
	pieChart := chart.NewPolarProportionalChart()
	pieChart.SetTitle(title)

	if len(dataMap) == 0 {
		return container.NewCenter(widget.NewLabel("No data available for this transaction type."))
	}

	var chartData []data.ProportionalDataPoint
	for category, count := range dataMap {
		chartData = append(chartData, data.ProportionalDataPoint{
			C:   category,
			Val: float64(count),
			Col: randomColor(),
		})
	}

	_, err := pieChart.AddProportionalSeries(title, chartData)
	if err != nil {
		log.Printf("Error adding data to pie chart: %v", err)
	}
	return pieChart
}

// This helper function remains unchanged
func randomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
