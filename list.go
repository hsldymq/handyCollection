package handyCollection

type Sortable[T any] interface {
}

type ICollection[T any] interface {
	Count() int
	Contains(T) bool
	Add(T)
	Remove(T) bool
}

type ItemInfo1[T any] struct {
	Index int
	Item  T
}

type IList[T any] interface {
	ICollection[T]

	Get(index int) (T, bool)
	IndexOf(item T) int
	Insert(item T) error
	RemoveAt(index int) error
	Clear()

	Filter(filter func(each ItemInfo1[T]) bool) IList[T]
	SelfFilter(filter func(each ItemInfo1[T]) bool)
	Sort(comparer func(a T, b T) bool) IList[T]
	SelfSort(comparer func(a T, b T) bool)
	Shuffle() IList[T]
	SelfShuffle()
}

type List[T any] struct {
	coll []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Add(item T) {
	l.coll = append(l.coll, item)
}

func (l *List[T]) Filter(filterFunc func(item ItemInfo1[T]) bool) IList[T] {
	newList := NewList[T]()
	for idx, item := range newList.coll {
		if filterFunc(ItemInfo1[T]{Index: idx, Item: item}) {
			newList.Add(item)
		}
	}
	return newList
}
