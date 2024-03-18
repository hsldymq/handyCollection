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
}
