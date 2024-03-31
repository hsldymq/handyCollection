package handy

import (
    "reflect"
)

func zVal[T any]() T {
    var v T
    return v
}

func isTypeComparable[T any]() bool {
    var v T
    return reflect.ValueOf(v).Comparable()
}
