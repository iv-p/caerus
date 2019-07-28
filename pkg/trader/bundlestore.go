package trader

type BundleStore struct {
}

type Bundle struct {
	Traders map[string]Trader
}
