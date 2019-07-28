package stats

type BundleStore struct {
}

type Bundle struct {
	Algorithms map[string]Algorithm
}

func NewBundleStore() *BundleStore {
	return &BundleStore{}
}

func (s *BundleStore) Load(id string) *Bundle {
	return &Bundle{}
}
