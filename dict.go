package handy

type KV[K comparable, V any] struct {
	K K
	V V
}

func (p KV[K, V]) ComparableKey() any {
	return p.K
}

func NewDict[K comparable, V any]() *Dict[K, V] {
	return &Dict[K, V]{
		m: make(map[K]V),
	}
}

type Dict[K comparable, V any] struct {
	m map[K]V
}

func (d *Dict[K, V]) Set(key K, val V) {
	d.m[key] = val
}

func (d *Dict[K, V]) SetOnce(key K, val V) bool {
	if _, ok := d.m[key]; ok {
		return false
	}
	d.m[key] = val
	return true
}
