package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/components"
)

// Основная функция
func handler(w http.ResponseWriter, r *http.Request) {
	symbol := "SBER" // Тикер Сбербанка
	// endDate := time.Now()
	// // endDate1 := endDate.AddDate(0, 0, 0)
	// startDate := endDate.AddDate(0, 0, -1)
	startDate := time.Date(2025, 1, 12, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 1, 13, 0, 0, 0, 0, time.UTC)
	candle_data, err := GetCandlesData(startDate, endDate, symbol, "10")
	if err != nil {
		return
	}
	candle_data60, err := GetCandlesData(startDate, endDate, symbol, "60")
	if err != nil {
		return
	}

	kline := DrawCandlestickChart(candle_data, symbol, "candle.html")

	kline60 := DrawCandlestickChart(candle_data60, symbol, "candle.html")

	line9, err := DrawEMALine(candle_data, kline, 9, "blue")
	if err != nil {
		return
	}

	line9_60, err := DrawEMALine(candle_data60, kline60, 9, "blue")
	if err != nil {
		return
	}

	line20, err := DrawEMALine(candle_data, kline, 20, "red")
	if err != nil {
		return
	}
	line50, err := DrawEMALine(candle_data, kline, 50, "green")
	if err != nil {
		return
	}
	line100, err := DrawEMALine(candle_data, kline, 100, "orange")
	if err != nil {
		panic(err)
	}

	page := components.NewPage()
	page.AddCharts(kline, kline60)

	kline60.Overlap(line9_60)

	kline.Overlap(line9)
	kline.Overlap(line20)
	kline.Overlap(line50)
	kline.Overlap(line100)

	err = page.Render(w)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
