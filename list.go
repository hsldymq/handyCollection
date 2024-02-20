package handyCollection

import (
	"github.com/hsldymq/goiter"
	"iter"
)

type Sortable[T any] interface {
}

type ICollection[T any] interface {
	Add(T)
	Remove(T) bool
	Count() int
	Contains(T) bool
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
	FilterSelf(filter func(each ItemInfo1[T]) bool)
	Sort(comparer func(a T, b T) bool) IList[T]
	SortSelf(comparer func(a T, b T) bool)
	Shuffle() IList[T]
	ShuffleSelf()
}

type List[T any] struct {
	coll   []T
	defVal T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Add(item T) {
	l.coll = append(l.coll, item)
}

func (l *List[T]) Clear() {
	l.coll = l.coll[:0]
}

func (l *List[T]) Get(idx int) (T, bool) {
	if idx >= len(l.coll) {
		return l.defVal, false
	}
	return l.coll[idx], true
}

func (l *List[T]) IndexOf(item T) int {
	for idx, each := range l.coll {
		if shallowEqual(each, item) {
			return idx
		}
	}

	return -1
}

func (l *List[T]) Contains(v T) bool {
	for _, each := range l.coll {
		if shallowEqual(each, v) {
			return true
		}
	}

	return false
}

func (l *List[T]) Count() int {
	return len(l.coll)
}

func (l *List[T]) Filter(filterFunc func(item ItemInfo1[T]) bool) *List[T] {
	newList := NewList[T]()
	for idx, item := range l.coll {
		if filterFunc(ItemInfo1[T]{Index: idx, Item: item}) {
			newList.Add(item)
		}
	}

	return newList
}

func (l *List[T]) FilterSelf(filterFunc func(item ItemInfo1[T]) bool) {
	newColl := make([]T, 0)
	for idx, item := range l.coll {
		if filterFunc(ItemInfo1[T]{Index: idx, Item: item}) {
			newColl = append(newColl, item)
		}
	}
	l.coll = newColl
}

func (l *List[T]) Iter() iter.Seq2[int, T] {
	return goiter.SliceIter(l.coll)
}

func (l *List[T]) Clone() *List[T] {
	clonedList := NewList[T]()
	clonedList.coll = make([]T, 0, l.Count())
	copy(clonedList.coll, l.coll)
	return clonedList
}

func shallowEqual(v1, v2 any) bool {
	return v1 == v2
}
