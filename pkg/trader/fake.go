package trader

import (
	"log"

	"github.com/ip-v/exchange/pkg/stats"
)

type Fake struct {
	Trader
}

func NewFake() *Fake {
	return &Fake{}
}

func (t *Fake) Trade(evaluation stats.ResultBundle) Decision {
	log.Println(evaluation)
	return Decision{SELL, 1.0}
}
