package exchange

import (
	"github.com/ip-v/exchange/pkg/config"
	"github.com/ip-v/exchange/pkg/stock"
)

type Factory struct {
	exchanges map[string]Exchange
}

type Type int

const (
	FakeExchange Type = iota
)

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewExchange(c config.Exchange) *Exchange {
	api := NewAPI(c.API)
	var pairs []stock.Pair
	for _, p := range c.Pairs {
		pairs = append(pairs, stock.NewPair(p))
	}
	return NewExchange(api, pairs)
}
