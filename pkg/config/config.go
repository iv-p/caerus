package config

type Config struct {
	Exchanges  []Exchange
	Pairs      []Pair
	Algorithms []Algorithm
}

type Exchange struct {
	ID   string
	Name string
	Type int
}

type Pair struct {
	ID           string
	Name         string
	ExchangeID   string
	AlgorithmIDs []string
}

type Algorithm struct {
	ID   string
	Name string
	Type int
}

var Fake = Config{
	Exchanges: []Exchange{
		Exchange{"1", "FakeExchange", 0},
	},
	Algorithms: []Algorithm{
		Algorithm{"1", "FakeAlgorithm", 0},
	},
	Pairs: []Pair{
		Pair{"1", "BTCETH", "1", []string{"1"}},
	},
}
