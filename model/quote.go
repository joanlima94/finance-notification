package model

import "time"

type HistoricalDataPrice struct {
	Date          int64   `json:"date"`
	Open          float64 `json:"open"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	AdjustedClose float64 `json:"adjustedClose"`
}

type Quote struct {
	Symbol                     string                `json:"symbol"`
	ShortName                  string                `json:"shortName"`
	RegularMarketPrice         float64               `json:"regularMarketPrice"`
	RegularMarketChange        float64               `json:"regularMarketChange"`
	RegularMarketChangePercent float64               `json:"regularMarketChangePercent"`
	RegularMarketVolume        int64                 `json:"regularMarketVolume"`
	HistoricalDataPrice        []HistoricalDataPrice `json:"historicalDataPrice"`
}

type BrapiResponse struct {
	Results []Quote `json:"results"`
}

type Stock struct {
	Stock  string  `json:"stock"`
	Name   string  `json:"name"`
	Close  float64 `json:"close"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

type StockListResponse struct {
	Stocks []Stock `json:"stocks"`
}

type HistoricalDataPriceOutput struct {
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	Close         float64 `json:"close"`
	Volume        int64   `json:"volume"`
	AdjustedClose float64 `json:"adjustedClose"`
}

type QuoteOutput struct {
	Symbol                     string                     `json:"symbol"`
	ShortName                  string                     `json:"shortName"`
	RegularMarketPrice         float64                    `json:"regularMarketPrice"`
	RegularMarketChange        float64                    `json:"regularMarketChange"`
	RegularMarketChangePercent float64                    `json:"regularMarketChangePercent"`
	RegularMarketVolume        int64                      `json:"regularMarketVolume"`
	HistoricalDataPrice        *HistoricalDataPriceOutput `json:"historicalDataPrice,omitempty"`
}

func unixToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("02/01/2006")
}

func ToQuoteOutput(q Quote) QuoteOutput {
	output := QuoteOutput{
		Symbol:                     q.Symbol,
		ShortName:                  q.ShortName,
		RegularMarketPrice:         q.RegularMarketPrice,
		RegularMarketChange:        q.RegularMarketChange,
		RegularMarketChangePercent: q.RegularMarketChangePercent,
		RegularMarketVolume:        q.RegularMarketVolume,
	}
	if len(q.HistoricalDataPrice) > 0 {
		price := q.HistoricalDataPrice[len(q.HistoricalDataPrice)-1]
		output.HistoricalDataPrice = &HistoricalDataPriceOutput{
			Date:          unixToDate(price.Date),
			Open:          price.Open,
			Close:         price.Close,
			Volume:        price.Volume,
			AdjustedClose: price.AdjustedClose,
		}
	}
	return output
}