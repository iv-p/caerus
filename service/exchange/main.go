package main

import (
	"github.com/ip-v/exchange/pkg/algorithm"
	"github.com/ip-v/exchange/pkg/exchange"
	"github.com/ip-v/exchange/pkg/kline"
	"github.com/ip-v/exchange/pkg/ray"
	"github.com/ip-v/exchange/pkg/trader"
)

func main() {
	algorithmFactory := algorithm.NewFactory()
	algorithmEvaluator := algorithm.NewSequentialEvaluator()
	exchangeFactory := exchange.NewFactory()
	klineRepository := kline.NewInMemoryRepository()
	trader := trader.NewFake()

	rayBuilder := ray.NewBuilder(algorithmFactory, algorithmEvaluator, exchangeFactory, klineRepository)
	ray := rayBuilder.
		WithAlgorithm(algorithm.DumbAlgorithm).
		WithExchange(exchange.FakeExchange).
		WithTrader(trader).
		Build()

	ray.Tick()
}
