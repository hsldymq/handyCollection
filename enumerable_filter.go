package handyCollection

import (
	"github.com/hsldymq/goiter"
	"iter"
)

type Distinctable interface {
	DistinctKey() any
}

func distinct[T any](e Enumerable[T]) Enumerable[T] {
	typeComparable := isTypeComparable[T]()
	_, distinctable := any(zeroVal[T]()).(Distinctable)
	if distinctable {
		return distinctBy(e, func(v T) any { return any(v).(Distinctable).DistinctKey() })
	} else if typeComparable {
		return distinctBy(e, func(v T) any { return v })
	} else {
		return newEnumerator(e.Iter())
	}
}

func distinctBy[T any](e Enumerable[T], keySelector func(v T) any) Enumerable[T] {
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

func filter[T any](e Enumerable[T], predicate func(T) bool) Enumerable[T] {
	return newEnumerator(goiter.Filter(e.Iter(), predicate))
}
