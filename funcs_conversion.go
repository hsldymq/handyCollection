package handy

func ToList[T any](e Iterable[T]) *List[T] {
    l := NewList[T]()
    for each := range e.Iter() {
        l.Add(each)
    }
    return l
}
