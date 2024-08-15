package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

type Joined[T1, T2 any] struct {
    Outer T1
    Inner T2
}

func Join[OuterT, InnerT any, K comparable](
    outer Iterable[OuterT],
    inner Iterable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
) Enumerable[*Joined[OuterT, InnerT]] {
    seq := func(yield func(*Joined[OuterT, InnerT]) bool) {
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
                    combined := &Joined[OuterT, InnerT]{
                        Outer: outerElem,
                        Inner: innerElem,
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
    outer Iterable[OuterT],
    inner Iterable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
    transformer func(OuterT, InnerT) ResultT,
) Enumerable[ResultT] {
    joinedSeq := Join(outer, inner, outerKeySelector, innerKeySelector).Iter()
    transformedSeq := goiter.Transform(joinedSeq, func(combined *Joined[OuterT, InnerT]) ResultT {
        return transformer(combined.Outer, combined.Inner)
    })

    return NewEnumerator(transformedSeq)
}

func GroupJoin[OuterT, InnerT any, K comparable](
    outer Iterable[OuterT],
    inner Iterable[InnerT],
    outerKeySelector func(OuterT) K,
    innerKeySelector func(InnerT) K,
) Enumerable[*Joined[OuterT, Enumerable[InnerT]]] {
    seq := func(yield func(joined *Joined[OuterT, Enumerable[InnerT]]) bool) {
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
                combined := &Joined[OuterT, Enumerable[InnerT]]{
                    Outer: outerElem,
                    Inner: NewEnumerator(goiter.SliceElems(innerElems)),
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
