package exchange

import (
	"log"

	"github.com/ip-v/exchange/pkg/stockprice"
)

type Fake struct {
	ID   string
	Name string

	Exchange
}

func NewFake() *Fake {
	return &Fake{ID: "1", Name: "FakeExchange"}
}

func (e *Fake) GetID() string {
	return e.ID
}

func (e *Fake) FetchPrice(pair Pair) stockprice.StockPrice {
	log.Println("Ferching price")
	return stockprice.StockPrice{}
}

func (e *Fake) ExecuteOrder(order Order) {
	log.Println("Executing order")
}
