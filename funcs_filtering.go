package handy

import (
	"github.com/hsldymq/goiter"
	"iter"
)

func filter[T any](e Iterable[T], predicate func(T) bool) Enumerable[T] {
	return newEnumerator(goiter.Filter(e.Iter(), predicate))
}

func distinct[T any](e Enumerable[T]) Enumerable[T] {
	typeComparable := isTypeComparable[T]()
	_, comparableImpl := any(zVal[T]()).(Comparable)
	if comparableImpl {
		return distinctBy(e, func(v T) any { return any(v).(Comparable).ComparableKey() })
	} else if typeComparable {
		return distinctBy(e, func(v T) any { return v })
	} else {
		return e
	}
}

func distinctBy[T any](e Iterable[T], keySelector func(v T) any) Enumerable[T] {
	seq := func(yield func(T) bool) {
		yielded := map[any]bool{}

		next, stop := iter.Pull(e.Iter())
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			k := keySelector(v)
			if yielded[k] {
				continue
			}
			yielded[k] = true
			if !yield(v) {
				return
			}
		}
	}
	return newEnumerator(seq)
}
