package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func newEnumerator[TIter goiter.SeqX[T], T any](it TIter) *Enumerator[T] {
    return &Enumerator[T]{
        iterator: goiter.Iterator[T](it),
    }
}

type Enumerator[T any] struct {
    iterator goiter.Iterator[T]
}

func (e *Enumerator[T]) Iter() iter.Seq[T] {
    return e.iterator.Seq()
}

func (e *Enumerator[T]) Count() int {
    return e.iterator.Count()
}

func (e *Enumerator[T]) Any(predicate func(T) bool) bool {
    for each := range e.iterator {
        if predicate(each) {
            return true
        }
    }
    return false
}

func (e *Enumerator[T]) All(predicate func(T) bool) bool {
    for each := range e.iterator {
        if !predicate(each) {
            return false
        }
    }
    return true
}

func (e *Enumerator[T]) Filter(predicate func(T) bool) Enumerable[T] {
    return filter(e, predicate)
}

func (e *Enumerator[T]) Take(n int) Enumerable[T] {
    return newEnumerator(goiter.Take(e.Iter(), n))
}

func (e *Enumerator[T]) TakeLast(n int) Enumerable[T] {
    return newEnumerator(goiter.TakeLast(e.Iter(), n))
}

func (e *Enumerator[T]) Skip(n int) Enumerable[T] {
    return newEnumerator(goiter.Skip(e.Iter(), n))
}

func (e *Enumerator[T]) SkipLast(n int) Enumerable[T] {
    return newEnumerator(goiter.SkipLast(e.Iter(), n))
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

func (e *Enumerator[T]) SequenceEqual(target Enumerable[T]) bool {
    return sequenceEqual[T](e, target)
}

func (e *Enumerator[T]) SequenceEqualBy(target Enumerable[T], keySelector func(T) any) bool {
    return sequenceEqualBy(e, target, keySelector)
}

func (e *Enumerator[T]) Concat(iterables ...Iterable[T]) Enumerable[T] {
    return concat(e, iterables...)
}

func (e *Enumerator[T]) OrderBy(cmp func(a, b T) int) Enumerable[T] {
    return orderBy(e, cmp)
}
