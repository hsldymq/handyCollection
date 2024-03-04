//go:build goexperiment.rangefunc

package handyCollection

import (
	"github.com/hsldymq/goiter"
	"iter"
	"slices"
)

type List[T any] struct {
	coll []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Get(idx int) (T, bool) {
	if idx >= len(l.coll) || idx < 0 {
		return zeroVal[T](), false
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
	return l.IndexOf(v) != -1
}

func (l *List[T]) Add(items ...T) {
	l.coll = slices.Concat(l.coll, items)
}

func (l *List[T]) Concat(tl ...*List[T]) {
	if len(tl) == 0 {
		return
	}

	s := make([][]T, 0, len(tl)+1)
	s = append(s, l.coll)
	for _, each := range tl {
		s = append(s, each.coll)
	}
	l.coll = slices.Concat(s...)
}

func (l *List[T]) ConcatSlices(sl ...[]T) {
	if len(sl) == 0 {
		return
	}

	s := make([][]T, 0, len(sl)+1)
	s = append(s, l.coll)
	s = append(s, sl...)
	l.coll = slices.Concat(s...)
}

func (l *List[T]) Clone() *List[T] {
	clonedList := NewList[T]()
	clonedList.coll = make([]T, 0, l.Count())
	copy(clonedList.coll, l.coll)
	return clonedList
}

func (l *List[T]) Clear() {
	l.coll = l.coll[:0]
}

func (l *List[T]) Count() int {
	return len(l.coll)
}

func (l *List[T]) Iter() iter.Seq[T] {
	return goiter.SliceElem(l.coll)
}

func (l *List[T]) Iter2() iter.Seq2[int, T] {
	return goiter.Slice(l.coll)
}
