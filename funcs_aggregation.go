package handy

import "github.com/hsldymq/goiter"

func Reduce[T, R any](e Iterable[T], init R, reducer func(R, T) R) R {
	return goiter.Fold(e.Iter(), init, reducer)
}
