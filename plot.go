package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// DrawCandlestickChart функция для рисования графика свечей
func DrawCandlestickChart(candles []Candle, ticker string, filename string) *charts.Kline {
	// Создание свечного графика с помощью go-echarts
	kline := charts.NewKLine()
	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "DataZoom(inside&slider)",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "inside",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	candle_data := make([]opts.KlineData, 0)
	xAxisData := make([]string, 0)

	for _, candle := range candles {
		candle_data = append(candle_data, opts.KlineData{Value: []interface{}{candle.Open, candle.Close, candle.Low, candle.High}})
		xAxisData = append(xAxisData, candle.Begin)
	}
	kline.SetXAxis(xAxisData).
		AddSeries(ticker, candle_data).
		SetSeriesOptions(
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{},
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color:        "green", // Цвет восходящей свечи
				Color0:       "red",   // Цвет нисходящей свечи
				BorderColor:  "green", // Цвет границы восходящей свечи
				BorderColor0: "red",   // Цвет границы нисходящей свечи
				Opacity:      0.8,
			}))

	// Сохраняем график в файл
	// f, err := os.Create(filename)
	// if err != nil {
	// 	return nil //fmt.Errorf("failed to create file: %w", err)
	// }
	// defer f.Close()

	// kline.Render(f)
	return kline
}

func DrawEMALine(candle_data []Candle, kline *charts.Kline, period int, color string) *charts.Line {
	line := charts.NewLine()
	ema9, err := IndicatorEMA(candle_data, period)
	if err != nil {
		return nil
	}
	items := make([]opts.LineData, 0)
	for _, d := range ema9 {
		if d == 0 {
			items = append(items, opts.LineData{Value: nil})
			continue
		}
		items = append(items, opts.LineData{Value: d})
	}
	line.AddSeries("Линия", items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
			Smooth:     opts.Bool(true), // Убираем точки
			ShowSymbol: opts.Bool(false),
		}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color: color,
			}))

	return line
}
