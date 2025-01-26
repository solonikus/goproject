package main

import (
	"fmt"
)

// calculateEMA рассчитывает экспоненциальную скользящую среднюю (EMA).
func IndicatorEMA(candles []Candle, period int) ([]float64, error) {
	if len(candles) < period {
		return nil, fmt.Errorf("недостаточно данных для расчета EMA (необходимо %d, получено %d)", period, len(candles))
	}

	emaValues := make([]float64, len(candles))
	k := 2.0 / float64(period+1)

	// Рассчитываем первое значение EMA как SMA
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += candles[i].Close
	}
	emaValues[period-1] = sum / float64(period)

	// Рассчитываем последующие значения EMA
	for i := period; i < len(candles); i++ {
		emaValues[i] = (candles[i].Close * k) + (emaValues[i-1] * (1 - k))
	}

	return emaValues, nil
}
