package ray

import (
	"time"

	"github.com/ip-v/exchange/pkg/exchange"
	"github.com/ip-v/exchange/pkg/stats"
	"github.com/ip-v/exchange/pkg/trader"
)

type Ray struct {
	pair       pair.Pair
	exchange   exchange.Exchange
	algorithms stats.Bundle
	traders    trader.Bundle
}

func NewRay() *Ray {
	return &Ray{}
}

func (r *Ray) Tick() {
	kline := r.exchange.Fetch(r.pair)
	r.klineRepository.Store(r.pair, kline)
	klines := r.klineRepository.Load(r.pair, time.Now(), time.Now())
	evaluation := r.algorithmEvaluator.Evaluate(r.algorithms, klines)
	r.trader.Trade(evaluation)
}
