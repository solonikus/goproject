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
	line9 := DrawEMALine(candle_data, kline, 9, "blue")
	line20 := DrawEMALine(candle_data, kline, 20, "red")
	line50 := DrawEMALine(candle_data, kline, 50, "green")
	line100 := DrawEMALine(candle_data, kline, 100, "orange")

	line, err := DrawEMALine(candle_data, kline, 9)
	if err != nil {
		panic(err)
	}
	// page := components.NewPage()
	kline.Overlap(line9)
	kline.Overlap(line20)
	kline.Overlap(line50)
	kline.Overlap(line100)

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
