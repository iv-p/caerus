package stats

import "github.com/ip-v/exchange/pkg/stock"

type Factory struct {
}

type Algorithm interface {
	Compute([]stock.Candle) Decision
	GetType() Type
	GetId() string
}

type Type int

const (
	DumbAlgorithm   Type = iota
	RandomAlgorithm Type = iota
)

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) NewAlgorithm(t Type) Algorithm {
	switch t {
	case DumbAlgorithm:
		return NewDumb()
	case RandomAlgorithm:
		return NewRandom()
	default:
		return NewDumb()
	}
}
