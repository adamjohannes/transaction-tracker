package gui

import (
	"image/color"
	"log"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/chart"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

func makeDataContent(g *GUI) fyne.CanvasObject {
	transactionCounts, err := g.transactionController.GetTransactionCountByTypeAndCategory()
	if err != nil {
		log.Printf("Error getting transaction counts: %v", err)
		dialog.ShowError(err, g.dataWindow)
		return container.NewCenter(widget.NewLabel("Error loading data"))
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Debit", createPieChart("Debit Transactions by Category", transactionCounts["Debit"])),
		container.NewTabItem("Credit", createPieChart("Credit Transactions by Category", transactionCounts["Credit"])),
		container.NewTabItem("Refund", createPieChart("Refund Transactions by Category", transactionCounts["Refund"])),
	)

	title := widget.NewLabel("Transaction Data")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true

	return container.NewBorder(title, nil, nil, nil, tabs)
}

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

func randomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
