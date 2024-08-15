package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func NewIterableFromSlice[T any](slice []T) Iterable[T] {
    return &sliceIterable[T]{s: slice}
}

func NewIterableFromSeq[TIter goiter.SeqX[T], T any](iter TIter) Iterable[T] {
    return &seqIterable[TIter, T]{iter: iter}
}

type Iterable[T any] interface {
    Iter() iter.Seq[T]
}

type sliceIterable[T any] struct {
    s []T
}

func (si *sliceIterable[T]) Iter() iter.Seq[T] {
    return goiter.SliceElems(si.s).Seq()
}

type seqIterable[TIter goiter.SeqX[T], T any] struct {
    iter TIter
}

func (si *seqIterable[TIter, T]) Iter() iter.Seq[T] {
    return iter.Seq[T](si.iter)
}
