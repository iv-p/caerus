package stats

import "github.com/ip-v/exchange/pkg/stock"

type Dumb struct {
	ID string

	Algorithm
}

func NewDumb() *Dumb {
	return &Dumb{}
}

func (d *Dumb) Compute(prices []stock.Candle) Decision {
	return Decision{
		Type:      BUY,
		Certainty: 0.56,
	}
}

func (d *Dumb) GetType() Type {
	return DumbAlgorithm
}

func (d *Dumb) GetID() string {
	return d.ID
}
