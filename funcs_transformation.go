package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func Transform[In any, Out any](e Iterable[In], transformer func(In) Out) Enumerable[Out] {
    return NewEnumerator(goiter.Transform(e.Iter(), transformer))
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
    return NewEnumerator(seq)
}
