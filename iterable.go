package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func NewSliceIterable[T any](slice []T) Iterable[T] {
    return &sliceIterable[T]{s: slice}
}

type Iterable[T any] interface {
    Iter() iter.Seq[T]
}

type sliceIterable[T any] struct {
    s []T
}

func (si *sliceIterable[T]) Iter() iter.Seq[T] {
    return goiter.SliceElem(si.s).Seq()
}
