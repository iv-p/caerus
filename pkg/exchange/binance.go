package exchange

import (
	"github.com/ip-v/exchange/pkg/config"
	"github.com/ip-v/exchange/pkg/order"
	"github.com/ip-v/exchange/pkg/stock"
)

type BinanceAPI struct {
	url string
}

func NewBinanceAPI(c config.API) API {
	return &BinanceAPI{c.URL}
}

func (a *BinanceAPI) FetchPrice(pair stock.Pair) stock.Candle {
	return stock.Candle{}
}

func (a *BinanceAPI) ExecuteOrder(order order.Order) {

}
