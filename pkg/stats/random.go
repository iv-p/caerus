package stats

import "github.com/ip-v/exchange/pkg/stock"

type Random struct {
	ID string

	Algorithm
}

func NewRandom() *Random {
	return &Random{}
}

func (a *Random) Compute(prices []stock.Candle) Decision {
	return Decision{
		Type:      BUY,
		Certainty: 0.98,
	}
}

func (r *Random) GetType() Type {
	return RandomAlgorithm
}

func (r *Random) GetID() string {
	return r.ID
}
