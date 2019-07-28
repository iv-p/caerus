package exchange

type Exchange interface {
	
}

type SimpleExchange struct {
	api API
	pairs []Pair

	Exchange
}
