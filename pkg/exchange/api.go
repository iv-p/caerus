package exchange

import "time"

type Candle struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
	Time   time.Time
}

type Pair struct {
	Currency string
	Base     string
	Exchange Exchange
}

type Order struct {
	Type OrderType
}

type OrderType int

const (
	BUY  OrderType = iota
	SELL OrderType = iota
)

type API interface {
	FetchPrice(Pair) Candle
	ExecuteOrder(Order)
}

