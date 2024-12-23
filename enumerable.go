package handy

type Enumerable[T any] interface {
    Iterable[T]
    Count() int
    Any(func(T) bool) bool
    All(func(T) bool) bool
    Filter(func(T) bool) Enumerable[T]
    Take(int) Enumerable[T]
    TakeLast(int) Enumerable[T]
    Skip(int) Enumerable[T]
    SkipLast(int) Enumerable[T]
    Distinct() Enumerable[T]
    DistinctBy(func(T) any) Enumerable[T]
    Union(Iterable[T]) Enumerable[T]
    UnionBy(Iterable[T], func(T) any) Enumerable[T]
    Intersect(Iterable[T]) Enumerable[T]
    IntersectBy(Iterable[T], func(T) any) Enumerable[T]
    Except(Iterable[T]) Enumerable[T]
    ExceptBy(Iterable[T], func(T) any) Enumerable[T]
    SequenceEqual(Iterable[T]) bool
    SequenceEqualBy(Iterable[T], func(T) any) bool
    Concat(...Iterable[T]) Enumerable[T]
    First() (T, bool)
    FirstOrDefault(T) T
    Last() (T, bool)
    LastOrDefault(T) T
    OrderBy(func(T, T) int) Enumerable[T]
}

type Comparable interface {
    ComparingKey() any
}
