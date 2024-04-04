//go:build goexperiment.rangefunc

package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
)

func union[T any](e1, e2 Iterable[T]) Enumerable[T] {
    return distinct[T](NewEnumerator(goiter.Concat(e1.Iter(), e2.Iter())))
}

func unionBy[T any](e1, e2 Iterable[T], keySelector func(T) any) Enumerable[T] {
    return distinctBy(NewEnumerator(goiter.Concat(e1.Iter(), e2.Iter())), keySelector)
}

func intersect[T any](e1, e2 Iterable[T]) Enumerable[T] {
    typeComparable := isTypeComparable[T]()
    _, comparableImpl := any(zVal[T]()).(Comparable)
    if comparableImpl {
        return intersectBy(e1, e2, func(v T) any { return any(v).(Comparable).ComparingKey() })
    } else if typeComparable {
        return intersectBy(e1, e2, func(v T) any { return v })
    } else {
        return NewEnumerator(goiter.Empty[T]())
    }
}

func intersectBy[T any](e1, e2 Iterable[T], keySelector func(T) any) Enumerable[T] {
    seq := func(yield func(T) bool) {
        e2Keys := map[any]bool{}
        next2, stop2 := iter.Pull(e2.Iter())
        defer stop2()
        for {
            v, ok := next2()
            if !ok {
                break
            }
            k := keySelector(v)
            e2Keys[k] = true
        }

        e1Keys := map[any]bool{}
        next1, stop1 := iter.Pull(e1.Iter())
        defer stop1()
        for {
            v, ok := next1()
            if !ok {
                return
            }
            k := keySelector(v)
            if !e2Keys[k] {
                continue
            }
            if e1Keys[k] {
                continue
            }
            e1Keys[k] = true
            if !yield(v) {
                return
            }
        }
    }
    return NewEnumerator(seq)
}

func except[T any](e1, e2 Iterable[T]) Enumerable[T] {
    typeComparable := isTypeComparable[T]()
    _, comparableImpl := any(zVal[T]()).(Comparable)
    if comparableImpl {
        return exceptBy(e1, e2, func(v T) any { return any(v).(Comparable).ComparingKey() })
    } else if typeComparable {
        return exceptBy(e1, e2, func(v T) any { return v })
    } else {
        return NewEnumerator(e1.Iter())
    }
}

func exceptBy[T any](e1, e2 Iterable[T], keySelector func(T) any) Enumerable[T] {
    seq := func(yield func(T) bool) {
        e2Keys := map[any]bool{}
        next2, stop2 := iter.Pull(e2.Iter())
        defer stop2()
        for {
            v, ok := next2()
            if !ok {
                break
            }
            k := keySelector(v)
            e2Keys[k] = true
        }

        e1Keys := map[any]bool{}
        next1, stop1 := iter.Pull(e1.Iter())
        defer stop1()
        for {
            v, ok := next1()
            if !ok {
                return
            }
            k := keySelector(v)
            if e2Keys[k] {
                continue
            }
            if e1Keys[k] {
                continue
            }
            e1Keys[k] = true
            if !yield(v) {
                return
            }
        }
    }

    return NewEnumerator(seq)
}
