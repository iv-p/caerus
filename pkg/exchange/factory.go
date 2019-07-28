package exchange

import "github.com/ip-v/exchange/pkg/config"

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

func (f *Factory) NewExchange(c config.Exchange) Exchange {
	switch Type(c.Type) {
	case FakeExchange:
		return NewFake()
	default:
		return NewFake()
	}
}
