package handy

import (
    "github.com/hsldymq/goiter"
)

func Reduce[T, R any](e Iterable[T], init R, reducer func(R, T) R) R {
    return goiter.Reduce(e.Iter(), init, reducer)
}

func Scan[T, R any](e Iterable[T], init R, folder func(R, T) R) Enumerable[R] {
    return newEnumerator(goiter.Scan(e.Iter(), init, folder))
}
