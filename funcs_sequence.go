package handy

import (
	"github.com/hsldymq/goiter"
	"iter"
)

func sequenceEqual[T any](e1, e2 Enumerable[T]) bool {
	typeComparable := isTypeComparable[T]()
	_, comparableImpl := any(zeroVal[T]()).(Comparable)
	if comparableImpl {
		return sequenceEqualBy(e1, e2, func(v T) any { return any(v).(Comparable).ComparableKey() })
	} else if typeComparable {
		return sequenceEqualBy(e1, e2, func(v T) any { return v })
	}
	return false
}

func sequenceEqualBy[T any](e1, e2 Enumerable[T], keySelector func(T) any) bool {
	next1, stop1 := iter.Pull(e1.Iter())
	defer stop1()
	next2, stop2 := iter.Pull(e2.Iter())
	defer stop2()
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if keySelector(v1) != keySelector(v2) {
			return false
		}
	}
}

func concat[T any](e Enumerable[T], iterables ...Iterable[T]) Enumerable[T] {
	seqs := []iter.Seq[T]{e.Iter()}
	for _, each := range iterables {
		seqs = append(seqs, each.Iter())
	}

	return newEnumerator(goiter.Concat(seqs...))
}
