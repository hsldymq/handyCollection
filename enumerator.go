package handy

import (
	"github.com/hsldymq/goiter"
	"iter"
)

func newEnumerator[T any](seq iter.Seq[T]) *Enumerator[T] {
	return &Enumerator[T]{
		seq: seq,
	}
}

type Enumerator[T any] struct {
	seq iter.Seq[T]
}

func (e *Enumerator[T]) Iter() iter.Seq[T] {
	return e.seq
}

func (e *Enumerator[T]) Count() int {
	return goiter.Count(e.seq)
}

func (e *Enumerator[T]) Any(predicate func(T) bool) bool {
	for each := range e.seq {
		if predicate(each) {
			return true
		}
	}
	return false
}

func (e *Enumerator[T]) All(predicate func(T) bool) bool {
	for each := range e.seq {
		if !predicate(each) {
			return false
		}
	}
	return true
}

func (e *Enumerator[T]) Filter(predicate func(T) bool) Enumerable[T] {
	return filter(e, predicate)
}

func (e *Enumerator[T]) Distinct() Enumerable[T] {
	return distinct[T](e)
}

func (e *Enumerator[T]) DistinctBy(keySelector func(T) any) Enumerable[T] {
	return distinctBy(e, keySelector)
}

func (e *Enumerator[T]) Union(target Enumerable[T]) Enumerable[T] {
	return union(e, target)
}

func (e *Enumerator[T]) UnionBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return unionBy(e, target, keySelector)
}

func (e *Enumerator[T]) Intersect(target Enumerable[T]) Enumerable[T] {
	return intersect(e, target)
}

func (e *Enumerator[T]) IntersectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return intersectBy(e, target, keySelector)
}

func (e *Enumerator[T]) Expect(target Enumerable[T]) Enumerable[T] {
	return except(e, target)
}

func (e *Enumerator[T]) ExpectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
	return exceptBy(e, target, keySelector)
}
