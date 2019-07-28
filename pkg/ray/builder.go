package ray

import (
	"github.com/ip-v/exchange/pkg/exchange"
	"github.com/ip-v/exchange/pkg/kline"
	"github.com/ip-v/exchange/pkg/trader"
)

type Builder struct {
	exchangeFactory  *exchange.Factory
	algorithmFactory *algorithm.Factory

	ray *Ray
}

func NewBuilder(
	algorithmFactory *algorithm.Factory,
	algorithmEvaluator algorithm.Evaluator,
	exchangeFactory *exchange.Factory,
	klineRepository kline.Repository) *Builder {
	ray := NewRay()
	ray.algorithmEvaluator = algorithmEvaluator
	ray.klineRepository = klineRepository
	return &Builder{
		exchangeFactory:  exchangeFactory,
		algorithmFactory: algorithmFactory,
		ray:              ray,
	}
}

func (b *Builder) WithExchange(t exchange.Type) *Builder {
	b.ray.exchange = b.exchangeFactory.NewExchange(t)
	return b
}

func (b *Builder) WithAlgorithm(t algorithm.Type) *Builder {
	b.ray.algorithms = append(b.ray.algorithms, b.algorithmFactory.NewAlgorithm(t))
	return b
}

func (b *Builder) WithPair(p pair.Pair) *Builder {
	b.ray.pair = p
	return b
}

func (b *Builder) WithTrader(t trader.Trader) *Builder {
	b.ray.trader = t
	return b
}

func (b *Builder) Build() *Ray {
	return b.ray
}
