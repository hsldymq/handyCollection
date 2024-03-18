package handy

type KVPair[K comparable, V any] struct {
	Key K
	Val V
}

func (p KVPair[K, V]) ComparableKey() any {
	return p.Key
}
