package handy

func first[T any](e Iterable[T]) (T, bool) {
    for each := range e.Iter() {
        return each, true
    }
    return zVal[T](), false
}

func firstOrDefault[T any](e Iterable[T], def T) T {
    if v, hasVal := first(e); hasVal {
        return v
    }
    return def
}

func last[T any](e Iterable[T]) (T, bool) {
    last, hasVal := zVal[T](), false
    for each := range e.Iter() {
        last, hasVal = each, true
    }
    return last, hasVal
}

func lastOrDefault[T any](e Iterable[T], def T) T {
    if v, hasVal := last(e); hasVal {
        return v
    }
    return def
}
