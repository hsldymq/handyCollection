//go:build goexperiment.rangefunc

package handy

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

func (l *List[T]) Merge(tl ...*List[T]) {
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

func (l *List[T]) MergeSlices(sl ...[]T) {
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

func (l *List[T]) Iter() iter.Seq[T] {
	return goiter.SliceElem(l.elems)
}

func (l *List[T]) Count() int {
	return len(l.elems)
}

func (l *List[T]) Any(predicate func(T) bool) bool {
	for _, each := range l.elems {
		if predicate(each) {
			return true
		}
	}
	return false
}

func (l *List[T]) All(predicate func(T) bool) bool {
	for _, each := range l.elems {
		if !predicate(each) {
			return false
		}
	}
	return true
}

func (l *List[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return filter(l, predicate)
}

func (l *List[T]) Distinct() Enumerable[T] {
	return distinct[T](l)
}

func (l *List[T]) DistinctBy(keySelector func(T) any) Enumerable[T] {
	return distinctBy(l, keySelector)
}

func (l *List[T]) Union(target Enumerable[T]) Enumerable[T] {
	return union(l, target)
}

func (l *List[T]) UnionBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return unionBy(l, target, keySelector)
}

func (l *List[T]) Intersect(target Enumerable[T]) Enumerable[T] {
	return intersect(l, target)
}

func (l *List[T]) IntersectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return intersectBy(l, target, keySelector)
}

func (l *List[T]) Expect(target Enumerable[T]) Enumerable[T] {
	return except(l, target)
}

func (l *List[T]) ExpectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return exceptBy(l, target, keySelector)
}

func (l *List[T]) SequenceEqual(target Enumerable[T]) bool {
	return sequenceEqual[T](l, target)
}

func (l *List[T]) SequenceEqualBy(target Enumerable[T], keySelector func(T) any) bool {
	return sequenceEqualBy(l, target, keySelector)
}

func (l *List[T]) Concat(iterables ...Iterable[T]) Enumerable[T] {
	return concat(l, iterables...)
}

func (l *List[T]) OrderBy(cmp func(T, T) int) Enumerable[T] {
	return orderBy(l, cmp)
}
