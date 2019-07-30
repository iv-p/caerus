package config

type Config struct {
	Exchanges  []Exchange
	Algorithms []Algorithm
}

type Exchange struct {
	ID    string
	Name  string
	Pairs []Pair
	API   API
}

type API struct {
	URL  string
	Type APIType
}

type APIType int

const (
	BINANCE APIType = iota
)

type Pair struct {
	Symbol string
	Quote  string
	Base   string
}

type Algorithm struct {
	ID   string
	Name string
	Type int
}

var Fake = Config{
	Exchanges: []Exchange{
		Exchange{"1", "FakeExchange", []Pair{
			Pair{"BTCETH", "BTC", "ETH"},
		}, API{"http://binance.com", BINANCE}},
	},
	Algorithms: []Algorithm{
		Algorithm{"1", "FakeAlgorithm", 0},
	},
}
