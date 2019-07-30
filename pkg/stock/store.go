package stock

import (
	"time"
)

type Candle struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
	Time   time.Time
}

type CandleStore struct {
}

func NewCandleStore() *CandleStore {
	return &CandleStore{}
}

func (s *CandleStore) Store(candle Candle, pair Pair) {
}

func (s *CandleStore) LoadRange(pair Pair, fr, to time.Time) []Candle {
	return []Candle{}
}
