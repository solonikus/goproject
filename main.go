package main

import (
	"fmt"
	"net/http"
	"time"
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
	kline := DrawCandlestickChart(candle_data, symbol, "candle.html")
	line := DrawEMALine(candle_data, kline, 9)
	// page := components.NewPage()
	kline.Overlap(line)
	err = kline.Render(w)
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
