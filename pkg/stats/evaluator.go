package stats

import (
	"github.com/ip-v/exchange/pkg/stock"
)

type ResultType int

const (
	BUY  ResultType = iota
	HOLD ResultType = iota
	SELL ResultType = iota
)

type Decision struct {
	Type      ResultType
	Certainty float64
}

type Performance struct {
	Value float64
}

type Result struct {
	Decision    Decision
	Performance Performance
}

type Evaluator interface {
	Evaluate([]Algorithm, []stock.Candle) ResultBundle
}

type ResultBundle map[string]Result

type SequentialEvaluator struct {
}

func NewSequentialEvaluator() *SequentialEvaluator {
	return &SequentialEvaluator{}
}

func (e *SequentialEvaluator) Evaluate(bundle Bundle, prices []stock.Candle) ResultBundle {
	r := make(ResultBundle)
	for _, algorithm := range bundle.Algorithms {
		result := algorithm.Compute(prices)
		performance := Performance{}
		r[algorithm.GetId()] = Result{result, performance}
	}
	return r
}
