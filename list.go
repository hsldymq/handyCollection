package handy

import (
    "github.com/hsldymq/goiter"
    "iter"
    "math/rand"
    "slices"
    "time"
)

func NewList[T any](elems ...T) *List[T] {
    l := &List[T]{
        elems:      make([]T, 0, len(elems)),
        comparable: isTypeComparable[T](),
    }
    if len(elems) > 0 {
        l.Add(elems...)
    }
    return l
}

func NewListFromIter[T any](it iter.Seq[T]) *List[T] {
    l := NewList[T]()
    for each := range it {
        l.Add(each)
    }
    return l
}

type List[T any] struct {
    elems      []T
    comparable bool
}

func (l *List[T]) Add(items ...T) {
    l.elems = slices.Concat(l.elems, items)
}

func (l *List[T]) Merge(tl ...*List[T]) {
    if len(tl) == 0 {
        return
    }

    s := make([][]T, 0, len(tl)+1)
    s = append(s, l.elems)
    for _, each := range tl {
        s = append(s, each.elems)
    }
    l.elems = slices.Concat(s...)
}

func (l *List[T]) MergeSlices(sl ...[]T) {
    if len(sl) == 0 {
        return
    }

    s := make([][]T, 0, len(sl)+1)
    s = append(s, l.elems)
    s = append(s, sl...)
    l.elems = slices.Concat(s...)
}

func (l *List[T]) Remove(item T) {
    idx := l.IndexOf(item)
    if idx >= 0 {
        l.RemoveAt(idx)
    }
}

func (l *List[T]) RemoveAt(i int) (T, bool) {
    idx, valid := l.actualIndex(i)
    if !valid {
        return zVal[T](), false
    }

    lastIdx := len(l.elems) - 1
    elem := l.elems[idx]
    if idx == 0 {
        l.elems = l.elems[1:]
    } else if idx == lastIdx {
        l.elems = l.elems[:lastIdx]
    } else {
        copy(l.elems[idx:], l.elems[idx+1:])
        l.elems = l.elems[:lastIdx]
    }
    return elem, true
}

func (l *List[T]) Pop() (T, bool) {
    return l.RemoveAt(-1)
}

func (l *List[T]) Shift() (T, bool) {
    return l.RemoveAt(0)
}

func (l *List[T]) Get(idx int) (T, bool) {
    if idx >= len(l.elems) || idx < 0 {
        return zVal[T](), false
    }
    return l.elems[idx], true
}

func (l *List[T]) Find(predicate func(T) bool) (T, bool) {
    for _, each := range l.elems {
        if predicate(each) {
            return each, true
        }
    }
    return zVal[T](), false
}

func (l *List[T]) FindOrDefault(predicate func(T) bool, def T) T {
    v, found := l.Find(predicate)
    if !found {
        return def
    }
    return v
}

func (l *List[T]) FindLast(predicate func(T) bool) (T, bool) {
    for i := len(l.elems) - 1; i >= 0; i-- {
        each := l.elems[i]
        if predicate(each) {
            return each, true
        }
    }
    return zVal[T](), false
}

func (l *List[T]) FindLastOrDefault(predicate func(T) bool, def T) T {
    v, found := l.FindLast(predicate)
    if !found {
        return def
    }
    return v
}

func (l *List[T]) IndexOf(item T) int {
    return l.indexOf(item, 0)
}

func (l *List[T]) Contains(item T) bool {
    return l.IndexOf(item) >= 0
}

func (l *List[T]) Sort(cmp func(T, T) int) {
    slices.SortFunc(l.elems, cmp)
}

func (l *List[T]) StableSort(cmp func(T, T) int) {
    slices.SortStableFunc(l.elems, cmp)
}

func (l *List[T]) FilterSelf(predicate func(T) bool) {
    newElems := make([]T, 0, len(l.elems))
    for _, each := range l.elems {
        if predicate(each) {
            newElems = append(newElems, each)
        }
    }
    l.elems = newElems
}

func (l *List[T]) FilterTo(predicate func(T) bool) *List[T] {
    newElems := make([]T, 0, len(l.elems))
    for _, each := range l.elems {
        if predicate(each) {
            newElems = append(newElems, each)
        }
    }

    return NewList(newElems...)
}

func (l *List[T]) ShuffleSelf() {
    count := len(l.elems)
    if count < 2 {
        return
    }

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    // Fisher–Yates shuffle
    for i := count - 1; i > 0; i-- {
        idx := r.Intn(i + 1)
        l.elems[idx], l.elems[i] = l.elems[i], l.elems[idx]
    }
}

func (l *List[T]) ShuffleTo() *List[T] {
    newList := l.Clone()
    newList.ShuffleSelf()
    return newList
}

func (l *List[T]) Clone() *List[T] {
    clonedList := NewList[T]()
    clonedList.elems = make([]T, l.Count())
    copy(clonedList.elems, l.elems)
    return clonedList
}

func (l *List[T]) Clear() {
    l.elems = l.elems[:0]
}

func (l *List[T]) IterBackward() iter.Seq[T] {
    return goiter.SliceSourceElem(func() []T { return l.elems }, true).Seq()
}

func (l *List[T]) Iter() iter.Seq[T] {
    return goiter.SliceSourceElem(func() []T { return l.elems }).Seq()
}

func (l *List[T]) Count() int {
    return len(l.elems)
}

func (l *List[T]) Any(predicate func(T) bool) bool {
    for _, each := range l.elems {
        if predicate(each) {
            return true
        }
    }
    return false
}

func (l *List[T]) All(predicate func(T) bool) bool {
    for _, each := range l.elems {
        if !predicate(each) {
            return false
        }
    }
    return true
}

func (l *List[T]) Filter(predicate func(T) bool) Enumerable[T] {
    return filter(l, predicate)
}

func (l *List[T]) Take(n int) Enumerable[T] {
    return newEnumerator(goiter.Take(l.Iter(), n))
}

func (l *List[T]) TakeLast(n int) Enumerable[T] {
    iterator := func(yield func(T) bool) {
        if n <= 0 {
            return
        }

        elems := l.elems
        startIdx := max(len(elems)-n, 0)
        for i := 0; i < n && startIdx+i < len(elems); i++ {
            if !yield(elems[startIdx+i]) {
                return
            }
        }
    }

    return newEnumerator(iterator)
}

func (l *List[T]) Skip(n int) Enumerable[T] {
    return newEnumerator(goiter.Skip(l.Iter(), n))
}

func (l *List[T]) SkipLast(n int) Enumerable[T] {
    iterator := func(yield func(T) bool) {
        elems := l.elems
        endIdx := len(elems) - 1
        if n > 0 {
            endIdx = len(elems) - n - 1
        }

        for i := 0; i <= endIdx; i++ {
            if !yield(elems[i]) {
                return
            }
        }
    }

    return newEnumerator(iterator)
}

func (l *List[T]) Distinct() Enumerable[T] {
    return distinct[T](l)
}

func (l *List[T]) DistinctBy(keySelector func(T) any) Enumerable[T] {
    return distinctBy(l, keySelector)
}

func (l *List[T]) Union(target Enumerable[T]) Enumerable[T] {
    return union(l, target)
}

func (l *List[T]) UnionBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
    return unionBy(l, target, keySelector)
}

func (l *List[T]) Intersect(target Enumerable[T]) Enumerable[T] {
    return intersect(l, target)
}

func (l *List[T]) IntersectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
    return intersectBy(l, target, keySelector)
}

func (l *List[T]) Expect(target Enumerable[T]) Enumerable[T] {
    return except(l, target)
}

func (l *List[T]) ExpectBy(target Enumerable[T], keySelector func(T) any) Enumerable[T] {
    return exceptBy(l, target, keySelector)
}

func (l *List[T]) SequenceEqual(target Enumerable[T]) bool {
    return sequenceEqual[T](l, target)
}

func (l *List[T]) SequenceEqualBy(target Enumerable[T], keySelector func(T) any) bool {
    return sequenceEqualBy(l, target, keySelector)
}

func (l *List[T]) Concat(iterables ...Iterable[T]) Enumerable[T] {
    return concat(l, iterables...)
}

func (l *List[T]) OrderBy(cmp func(T, T) int) Enumerable[T] {
    return orderBy(l, cmp)
}

func (l *List[T]) actualIndex(idx int) (int, bool) {
    if idx < 0 {
        idx = len(l.elems) + idx
        if idx < 0 {
            return -1, false
        }
    } else if idx >= len(l.elems) {
        return -1, false
    }
    return idx, true
}

func (l *List[T]) indexOf(item T, startsFrom int) int {
    if !l.comparable {
        return -1
    }

    for idx := startsFrom; idx < len(l.elems); idx++ {
        elem := l.elems[idx]
        if any(elem) == any(item) {
            return idx
        }
    }

    return -1
}
