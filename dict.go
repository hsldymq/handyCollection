package handyCollection

type KVPair[K comparable, V any] struct {
	Key K
	Val V
}

func (p KVPair[K, V]) DistinctKey() any {
	return p.Key
}
