package handy

import "iter"

type Iterable[T any] interface {
	Iter() iter.Seq[T]
}

type Enumerable[T any] interface {
	Iterable[T]
	Count() int
	Any(func(T) bool) bool
	All(func(T) bool) bool
	Filter(func(T) bool) Enumerable[T]
	Distinct() Enumerable[T]
	DistinctBy(func(T) any) Enumerable[T]
	Union(Enumerable[T]) Enumerable[T]
	UnionBy(Enumerable[T], func(T) any) Enumerable[T]
	Intersect(Enumerable[T]) Enumerable[T]
	IntersectBy(Enumerable[T], func(T) any) Enumerable[T]
	Expect(Enumerable[T]) Enumerable[T]
	ExpectBy(Enumerable[T], func(T) any) Enumerable[T]
	SequenceEqual(Enumerable[T]) bool
	SequenceEqualBy(Enumerable[T], func(T) any) bool
	Concat(...Iterable[T]) Enumerable[T]
	OrderBy(func(T, T) int) Enumerable[T]
}

type Comparable interface {
	ComparingKey() any
}
