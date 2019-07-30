package exchange

import (
	"github.com/ip-v/exchange/pkg/config"
	"github.com/ip-v/exchange/pkg/order"
	"github.com/ip-v/exchange/pkg/stock"
)

type API interface {
	FetchPrice(stock.Pair) stock.Candle
	ExecuteOrder(order.Order)
}

func NewAPI(c config.API) API {
	switch c.Type {
	case config.BINANCE:
		return NewBinanceAPI(c)
	default:
		return NewBinanceAPI(c)
	}
}
