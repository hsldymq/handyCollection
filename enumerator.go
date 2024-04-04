package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func newEnumerator[TIter goiter.SeqX[T], T any](it TIter) *enumerator[T] {
    return &enumerator[T]{
        iterator: goiter.Iterator[T](it),
    }
}

type enumerator[T any] struct {
    iterator goiter.Iterator[T]
}

func (e *enumerator[T]) Iter() iter.Seq[T] {
    return e.iterator.Seq()
}

func (e *enumerator[T]) Count() int {
    return e.iterator.Count()
}

func (e *enumerator[T]) Any(predicate func(T) bool) bool {
    for each := range e.iterator {
        if predicate(each) {
            return true
        }
    }
    return false
}

func (e *enumerator[T]) All(predicate func(T) bool) bool {
    for each := range e.iterator {
        if !predicate(each) {
            return false
        }
    }
    return true
}

func (e *enumerator[T]) Filter(predicate func(T) bool) Enumerable[T] {
    return filter(e, predicate)
}

func (e *enumerator[T]) Take(n int) Enumerable[T] {
    return newEnumerator(goiter.Take(e.Iter(), n))
}

func (e *enumerator[T]) TakeLast(n int) Enumerable[T] {
    return newEnumerator(goiter.TakeLast(e.Iter(), n))
}

func (e *enumerator[T]) Skip(n int) Enumerable[T] {
    return newEnumerator(goiter.Skip(e.Iter(), n))
}

func (e *enumerator[T]) SkipLast(n int) Enumerable[T] {
    return newEnumerator(goiter.SkipLast(e.Iter(), n))
}

func (e *enumerator[T]) Distinct() Enumerable[T] {
    return distinct[T](e)
}

func (e *enumerator[T]) DistinctBy(keySelector func(T) any) Enumerable[T] {
    return distinctBy(e, keySelector)
}

func (e *enumerator[T]) Union(target Iterable[T]) Enumerable[T] {
    return union(e, target)
}

func (e *enumerator[T]) UnionBy(target Iterable[T], keySelector func(T) any) Enumerable[T] {
    return unionBy(e, target, keySelector)
}

func (e *enumerator[T]) Intersect(target Iterable[T]) Enumerable[T] {
    return intersect(e, target)
}

func (e *enumerator[T]) IntersectBy(target Iterable[T], keySelector func(T) any) Enumerable[T] {
    return intersectBy(e, target, keySelector)
}

func (e *enumerator[T]) Except(target Iterable[T]) Enumerable[T] {
    return except(e, target)
}

func (e *enumerator[T]) ExceptBy(target Iterable[T], keySelector func(T) any) Enumerable[T] {
    return exceptBy(e, target, keySelector)
}

func (e *enumerator[T]) SequenceEqual(target Iterable[T]) bool {
    return sequenceEqual[T](e, target)
}

func (e *enumerator[T]) SequenceEqualBy(target Iterable[T], keySelector func(T) any) bool {
    return sequenceEqualBy(e, target, keySelector)
}

func (e *enumerator[T]) Concat(iterables ...Iterable[T]) Enumerable[T] {
    return concat(e, iterables...)
}

func (e *enumerator[T]) OrderBy(cmp func(a, b T) int) Enumerable[T] {
    return orderBy(e, cmp)
}
