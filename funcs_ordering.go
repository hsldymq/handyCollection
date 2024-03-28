package handy

import (
	"cmp"
	"github.com/hsldymq/goiter"
)

func Order[T cmp.Ordered](e Iterable[T]) Enumerable[T] {
	return newEnumerator(goiter.Order(e.Iter()))
}

func orderBy[T any](e Iterable[T], cmp func(T, T) int) Enumerable[T] {
	return newEnumerator(goiter.OrderBy(e.Iter(), cmp))
}
