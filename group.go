package handyCollection

type GroupItemInfo[T any] struct {
	Key  string
	Item T
}

type Group[T any] struct {
	c Collection[T]
}

func NewGroup[T any]() *Group[T] {
	return &Group[T]{
		c: NewGeneralCollection[T](),
	}
}

func GroupCollection[T any](c Collection[T], grouper func(*ItemInfo[T]) string) *Group[Collection[T]] {
	g := NewGroup[Collection[T]]()
	c.ForEach(func(info *ItemInfo[T]) {
		groupKey := grouper(info)
		collection, found := g.Find(groupKey)
		if !found {
			collection = NewGeneralCollection[T]()
			g.Set(groupKey, collection)
		}
		if info.IsAutoGenKey {
			collection.Add(info.Item)
		} else {
			collection.AddWithKey(info.Item, info.Key)
		}
	})
	return g
}

func (g *Group[T]) Set(key string, val T) *Group[T] {
	g.c.AddWithKey(val, key)
	return g
}

func (g *Group[T]) Find(key string) (T, bool) {
	return g.c.FindByKey(key)
}

func (g *Group[T]) Keys() []string {
	keys := make([]string, 0, g.c.Count())
	g.c.ForEach(func(each *ItemInfo[T]) {
		keys = append(keys, each.Key)
	})
	return keys
}

func (g *Group[T]) AsSlice() []T {
	return g.c.AsSlice()
}

func (g *Group[T]) AsMap() map[string]T {
	return g.c.AsMap()
}

func (g *Group[T]) SelfSort(less func(i *GroupItemInfo[T], j *GroupItemInfo[T]) bool) *Group[T] {
	g.c.SelfSort(func(i *ItemInfo[T], j *ItemInfo[T]) bool {
		return less(&GroupItemInfo[T]{Key: i.Key, Item: i.Item}, &GroupItemInfo[T]{Key: j.Key, Item: j.Item})
	})

	return g
}
