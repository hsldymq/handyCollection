package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func concat[T any](e Iterable[T], iterables ...Iterable[T]) Enumerable[T] {
    seqs := []iter.Seq[T]{}
    for _, each := range iterables {
        seqs = append(seqs, each.Iter())
    }

    return NewEnumerator(goiter.Concat(e.Iter(), seqs...))
}
