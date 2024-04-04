package handy

import (
    "cmp"
    "github.com/hsldymq/goiter"
)

func Order[T cmp.Ordered](e Iterable[T], desc ...bool) Enumerable[T] {
    return NewEnumerator(goiter.Order(e.Iter(), desc...))
}

func orderBy[T any](e Iterable[T], cmp func(T, T) int) Enumerable[T] {
    return NewEnumerator(goiter.OrderBy(e.Iter(), cmp))
}
