package exchange

import (
	"github.com/ip-v/exchange/pkg/order"
	"github.com/ip-v/exchange/pkg/stock"
)

type Exchange struct {
	api   API
	pairs []stock.Pair
}

func NewExchange(api API, pairs []stock.Pair) *Exchange {
	return &Exchange{api, pairs}
}

func (e *Exchange) FetchPairData(pair stock.Pair) stock.Candle {
	return stock.Candle{}
}

func (e *Exchange) ExecuteOrder(order order.Order) {

}

func (e *Exchange) GetPairs() []stock.Pair {
	return e.pairs
}
