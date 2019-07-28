package trader

import "github.com/ip-v/exchange/pkg/stats"

type DecisionType int

const (
	BUY  DecisionType = iota
	HOLD DecisionType = iota
	SELL DecisionType = iota
)

type Trader interface {
	Trade(stats.ResultBundle) Decision
}

type Decision struct {
	Type   DecisionType
	Amount float64
}
