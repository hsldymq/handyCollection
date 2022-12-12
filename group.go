package handyCollection

import "sort"

type Group[T any] struct {
	c *GeneralCollection[T]
}

func NewGroup[T any]() *Group[T] {
	return &Group[T]{
		c: NewGeneralCollection[T](),
	}
}

func GroupCollectionBy[T any](c *GeneralCollection[T], grouper func(item T, idx int, key string) string) *Group[*GeneralCollection[T]] {
	g := NewGroup[*GeneralCollection[T]]()
	for idx, key := range c.orderedKeys {
		item := c.items[key]
		groupKey := grouper(item, idx, key)
		collection, found := g.Find(groupKey)
		if !found {
			collection = NewGeneralCollection[T]()
			g.Set(groupKey, collection)
		}
		if c.isAutoGenKey(key) {
			collection.Add(item)
		} else {
			collection.AddWithKey(item, key)
		}
	}
	return g
}

func (g *Group[T]) Set(key string, val T) *Group[T] {
	g.c.AddWithKey(val, key)
	return g
}

func (g *Group[T]) Find(key string) (T, bool) {
	return g.c.FindByKey(key)
}

func (g *Group[T]) KeyByIndex(idx int) (string, bool) {
	return g.c.KeyByIndex(idx)
}

func (g *Group[T]) AsSlice() []T {
	return g.c.AsSlice()
}

func (g *Group[T]) AsMap() map[string]T {
	return g.c.AsMap()
}

func (g *Group[T]) SelfSortBy(less func(iKey string, iVal T, jKey string, jVal T) bool) *Group[T] {
	sort.Slice(g.c.orderedKeys, func(i, j int) bool {
		iKey, _ := g.c.KeyByIndex(i)
		iVal := g.c.items[iKey]
		jKey, _ := g.c.KeyByIndex(j)
		jVal := g.c.items[jKey]

		return less(iKey, iVal, jKey, jVal)
	})
	g.c.clearKeysIndex()
	g.c.clearSliceCache()

	return g
}
