package handy

import (
	"github.com/hsldymq/goiter"
	"iter"
)

func Transform[In any, Out any](e Enumerable[In], transformer func(In) Out) Enumerable[Out] {
	return newEnumerator(goiter.Transform(e.Iter(), transformer))
}

func TransformExpand[In any, Out any](e Enumerable[In], transformer func(In) Iterable[Out]) Enumerable[Out] {
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

func ToList[T any](e Enumerable[T]) *List[T] {
	l := NewList[T]()
	for each := range e.Iter() {
		l.Add(each)
	}
	return l
}

func ToListBy[T, R any](e Enumerable[T], transformer func(T) R) *List[R] {
	l := NewList[R]()
	for each := range e.Iter() {
		l.Add(transformer(each))
	}
	return l
}

func ToDictBy[T any, K comparable, V any](e Enumerable[T], transformer func(T) (K, V)) *Dict[K, V] {
	l := NewDict[K, V]()
	for each := range e.Iter() {
		l.Set(transformer(each))
	}
	return l
}
