package handy

func NewMap[TK comparable, TV any]() *Map[TK, TV] {
    return &Map[TK, TV]{
        m: make(map[TK]TV),
    }
}

type Map[TK comparable, TV any] struct {
    m map[TK]TV
}

func (m *Map[TK, TV]) HasKey(key TK) bool {
    _, ok := m.m[key]
    return ok
}

func (m *Map[TK, TV]) Set(key TK, value TV) {
    m.m[key] = value
}

func (m *Map[TK, TV]) SetOnce(key TK, value TV) bool {
    if _, ok := m.m[key]; !ok {
        m.m[key] = value
        return true
    }
    return false
}

func (m *Map[TK, TV]) Merge(maps ...*Map[TK, TV]) {
    for _, each := range maps {
        for k, v := range each.m {
            m.m[k] = v
        }
    }
}
