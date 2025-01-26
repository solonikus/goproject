package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Candle структура для хранения данных свечи
type Candle struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Value  float64
	Volume int
	Begin  string
	End    string
}

// CandlesResponse структура для хранения ответа JSON
type CandlesResponse struct {
	Candles struct {
		Columns []string        `json:"columns"`
		Data    [][]interface{} `json:"data"`
	} `json:"candles"`
}

// Возвращает срез свеч за заданные даты, заданной акции и заданного интервала
func GetCandlesData(start time.Time, end time.Time, symbol string, interval string) ([]Candle, error) {
	url := "https://iss.moex.com/iss/engines/stock/markets/shares/securities/" +
		symbol + "/candles.json?iss.meta=off&from=" + start.Format("2006-01-02") + "&to=" + end.Format("2006-01-02") + "&interval=" + interval

	// Создаем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Создаем новый декодер для чтения JSON-ответа
	decoder := json.NewDecoder(resp.Body)

	// Распарсить JSON-ответ
	var result CandlesResponse
	err = decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Преобразовать данные в структуры Candle
	var candles []Candle
	for _, candleData := range result.Candles.Data {
		candle := Candle{
			Open:   candleData[0].(float64),
			Close:  candleData[1].(float64),
			High:   candleData[2].(float64),
			Low:    candleData[3].(float64),
			Value:  candleData[4].(float64),
			Volume: int(candleData[5].(float64)),
			Begin:  candleData[6].(string),
			End:    candleData[7].(string),
		}
		candles = append(candles, candle)
	}
	if len(candles) == 0 {
		return nil, fmt.Errorf("candles nil")
	}

	return candles, nil
}

// ---- Current Data ----

// Возвращает текущие рыночные данные выбранной акцииы
func GetLCurMarketData(symbol string) map[string]interface{} {
	url := fmt.Sprintf("https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/%s.json?iss.meta=off&iss.only=marketdata&iss.json=extended", symbol)
	resp, err := http.Get(url)
	if err != nil {
		// return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// return nil, err
	}

	var response []interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		// return nil, err
	}
	map_ticket := response[1].(map[string]interface{})["marketdata"].([]interface{})[0].(map[string]interface{})
	return map_ticket

	// Второй элемент массива содержит marketdata

	// return nil, nil
}

// ---- Historical Data ----

func GetHistoryData() {
	// Определяем даты для последней недели
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)
	fromDate := startDate.Format("2006-01-02")
	toDate := endDate.Format("2006-01-02")
	url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/TQBR/securities/SBER.json?from=%s&till=%s&interval=60&iss.meta=off&iss.only=history&iss.json=extended", fromDate, toDate)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var response []interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	// map_ticket := response[1].(map[string]interface{})["marketdata"].([]interface{})[0].(map[string]interface{})
}
