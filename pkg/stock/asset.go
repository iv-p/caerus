package stock

type Asset struct {
	Name   string
	Symbol string
}

type AssetPair struct {
	Quote Asset
	Base  Asset
}
