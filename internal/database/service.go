package database

import (
	"database/sql"
	"time"
)

type DatabaseService interface {
	SelectInstruments() ([]Instrument, error)
	InsertInstrument(item Instrument) error
	DeleteInstrument(symbol string) error
	UpdateInstrument(symbol string, item Instrument) error
	CreateInstrumentTable() error
	BulkInsertInstruments(items []Instrument) error
	TruncateAndBulkInsertInstruments(items []Instrument) error

	SelectPermittedInstruments() ([]Instrument, error)
	SelectNonPermittedInstruments() ([]Instrument, error)
	UpdatePermittedList(permitted []string) error
	SelectOpenOrders() ([]string, error)
	Close() error

	SelectLatestTrades(since time.Time) ([]PositionState, error)
}

type Instrument struct {
	Symbol           sql.NullString `json:"Symbol"`
	IsPermitted      bool           `json:"IsPermitted"`
	IsDefaultSetting bool           `json:"IsDefaultSettings"`
}

type PositionState struct {
	LaunchID             string  `json:"LaunchId"`
	DateTime             string  `json:"Datetime"`
	Symbol               string  `json:"Symbol"`
	Status               string  `json:"Status"`
	Side                 string  `json:"Side"`
	BuyCount             int64   `json:"BuyCount"`
	Quantity             float64 `json:"Quantity"`
	AveragePrice         float64 `json:"AveragePrice"`
	TakeProfitPrice      float64 `json:"TakeProfitPrice"`
	StopLossPrice        float64 `json:"StopLossPrice"`
	TakeProfitLimitPrice string  `json:"TakeProfitLimitPrice"`
	Reason               string  `json:"Reason"`
}
