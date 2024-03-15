package handyCollection

import "iter"

type Enumerable[T any] interface {
	Filter(func(T) bool) Enumerable[T]
	Iter() iter.Seq[T]
	Count() int
	Any(func(T) bool) bool
	All(func(T) bool) bool
}
