package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func Transform[In any, Out any](e Iterable[In], transformer func(In) Out) Enumerable[Out] {
    return newEnumerator(goiter.Transform(e.Iter(), transformer))
}

func TransformExpand[In any, Out any](e Iterable[In], transformer func(In) Iterable[Out]) Enumerable[Out] {
    seq := func(yield func(Out) bool) {
        next, stop := iter.Pull(e.Iter())
        defer stop()
        for {
            v, ok := next()
            if !ok {
                return
            }
            for each := range transformer(v).Iter() {
                if !yield(each) {
                    return
                }
            }
        }
    }
    return newEnumerator(seq)
}

func ToList[T any](e Iterable[T]) *List[T] {
    l := NewList[T]()
    for each := range e.Iter() {
        l.Add(each)
    }
    return l
}
