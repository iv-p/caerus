package stock

import "github.com/ip-v/exchange/pkg/config"

type Stock struct {
	Symbol string
}

type Pair struct {
	Quote Stock
	Base  Stock
}

func NewPair(c config.Pair) Pair {
	return Pair{Stock{c.Quote}, Stock{c.Base}}
}
