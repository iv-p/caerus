package stock

import (
	"time"

	"github.com/ip-v/exchange/pkg/exchange"
)

type CandleStore struct {
}

func NewCandleStore() *CandleStore {
	return &CandleStore{}
}

func (s *CandleStore) Store(Candle Candle, pair exchange.Pair) {
}

func (s *CandleStore) LoadRange(pair exchange.Pair, fr, to time.Time) []Candle {
	return []Candle{}
}
