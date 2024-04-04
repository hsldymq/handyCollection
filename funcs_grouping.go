package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

type Combined[T1, T2 any] struct {
    First  T1
    Second T2
}

func Join[OuterT, InnerT any, K comparable](
    outer Enumerable[OuterT],
    inner Enumerable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
) Enumerable[*Combined[OuterT, InnerT]] {
    seq := func(yield func(*Combined[OuterT, InnerT]) bool) {
        next, stop := iter.Pull(outer.Iter())
        defer stop()
        outerElem, ok := next()
        if !ok {
            return
        }

        group := groupIterableToMap[InnerT, K](inner, innerKeySelector)
        for ok {
            innerElems, hasAny := group[outerKeySelector(outerElem)]
            if hasAny {
                for _, innerElem := range innerElems {
                    combined := &Combined[OuterT, InnerT]{
                        First:  outerElem,
                        Second: innerElem,
                    }
                    if !yield(combined) {
                        return
                    }
                }
            }
            outerElem, ok = next()
        }
    }
    return NewEnumerator(seq)
}

func JoinAs[OuterT, InnerT any, K comparable, ResultT any](
    outer Enumerable[OuterT],
    inner Enumerable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
    transformer func(OuterT, InnerT) ResultT,
) Enumerable[ResultT] {
    joinedSeq := Join(outer, inner, outerKeySelector, innerKeySelector).Iter()
    transformedSeq := goiter.Transform(joinedSeq, func(combined *Combined[OuterT, InnerT]) ResultT {
        return transformer(combined.First, combined.Second)
    })

    return NewEnumerator(transformedSeq)
}

func GroupJoin[OuterT, InnerT any, K comparable](
    outer Enumerable[OuterT],
    inner Enumerable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
) Enumerable[*Combined[OuterT, Enumerable[InnerT]]] {
    seq := func(yield func(*Combined[OuterT, Enumerable[InnerT]]) bool) {
        next, stop := iter.Pull(outer.Iter())
        defer stop()
        outerElem, ok := next()
        if !ok {
            return
        }

        group := groupIterableToMap[InnerT, K](inner, innerKeySelector)
        for ok {
            innerElems, hasAny := group[outerKeySelector(outerElem)]
            if hasAny {
                combined := &Combined[OuterT, Enumerable[InnerT]]{
                    First:  outerElem,
                    Second: NewEnumerator(goiter.SliceElem(innerElems)),
                }
                if !yield(combined) {
                    return
                }
            }
            outerElem, ok = next()
        }
    }
    return NewEnumerator(seq)
}

func GroupJoinAs[OuterT, InnerT any, K comparable, ResultT any](
    outer Enumerable[OuterT],
    inner Enumerable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
    transformer func(OuterT, Enumerable[InnerT]) ResultT,
) Enumerable[ResultT] {
    joinedSeq := GroupJoin(outer, inner, outerKeySelector, innerKeySelector).Iter()
    transformedSeq := goiter.Transform(joinedSeq, func(combined *Combined[OuterT, Enumerable[InnerT]]) ResultT {
        return transformer(combined.First, combined.Second)
    })

    return NewEnumerator(transformedSeq)
}

func groupIterableToMap[T any, K comparable](iterable Iterable[T], keySelector func(T) K) map[any][]T {
    m := map[any][]T{}
    next, stop := iter.Pull(iterable.Iter())
    defer stop()
    for {
        v, ok := next()
        if !ok {
            return m
        }
        k := keySelector(v)
        m[k] = append(m[k], v)
    }
}
