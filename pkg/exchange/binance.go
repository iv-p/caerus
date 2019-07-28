package exchange

type BinanceAPI struct {
	url string

	API
}

func NewBinanceAPI() *BinanceAPI {
	return &BinanceAPI{}
}

func (a *BinanceAPI) FetchPair(pair Pair) Candle {
	return Candle{}
}

func (a *BinanceAPI) ExecuteOrder(Order) {
	
}