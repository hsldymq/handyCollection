package handy

type KVPair[K comparable, V any] struct {
	Key K
	Val V
}

func (p KVPair[K, V]) ComparableKey() any {
	return p.Key
}

func NewDict[K comparable, V any]() *Dict[K, V] {
	return &Dict[K, V]{}
}

type Dict[K comparable, V any] struct {
	// TODO
}

func (d *Dict[K, V]) Set(key K, val V) {
	// TODO
}

func (d *Dict[K, V]) SetOnce(key K, val V) bool {
	// TODO
	return true
}
