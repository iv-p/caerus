package order

type Order struct {
	Type OrderType
}

type OrderType int

const (
	BUY  OrderType = iota
	SELL OrderType = iota
)

type Store struct {
}
