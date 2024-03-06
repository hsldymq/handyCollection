//go:build goexperiment.rangefunc

package handyCollection

import (
	"github.com/hsldymq/goiter"
	"iter"
	"reflect"
	"slices"
)

type List[T any] struct {
	elems      []T
	comparable bool
}

func NewList[T any]() *List[T] {
	return &List[T]{
		comparable: reflect.ValueOf(zeroVal[T]()).Comparable(),
	}
}

func (l *List[T]) Get(idx int) (T, bool) {
	if idx >= len(l.elems) || idx < 0 {
		return zeroVal[T](), false
	}
	return l.elems[idx], true
}

func (l *List[T]) Index(item T) (result int) {
	result = -1
	if !l.comparable {
		return
	}

	defer func() { recover() }()
	for idx, elem := range l.elems {
		if shallowEqual(elem, item) {
			result = idx
			break
		}
	}

	return
}

func (l *List[T]) Contains(v T) bool {
	return l.Index(v) >= 0
}

func (l *List[T]) Add(items ...T) {
	l.elems = slices.Concat(l.elems, items)
}

func (l *List[T]) Concat(tl ...*List[T]) {
	if len(tl) == 0 {
		return
	}

	s := make([][]T, 0, len(tl)+1)
	s = append(s, l.elems)
	for _, each := range tl {
		s = append(s, each.elems)
	}
	l.elems = slices.Concat(s...)
}

func (l *List[T]) ConcatSlices(sl ...[]T) {
	if len(sl) == 0 {
		return
	}

	s := make([][]T, 0, len(sl)+1)
	s = append(s, l.elems)
	s = append(s, sl...)
	l.elems = slices.Concat(s...)
}

func (l *List[T]) Clone() *List[T] {
	clonedList := NewList[T]()
	clonedList.elems = make([]T, 0, l.Count())
	copy(clonedList.elems, l.elems)
	return clonedList
}

func (l *List[T]) Clear() {
	l.elems = l.elems[:0]
}

func (l *List[T]) Count() int {
	return len(l.elems)
}

func (l *List[T]) Iter() iter.Seq[T] {
	return goiter.SliceElem(l.elems)
}

func (l *List[T]) Iter2() iter.Seq2[int, T] {
	return goiter.Slice(l.elems)
}
