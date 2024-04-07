package handy

import (
    "reflect"
)

func zVal[T any]() (v T) {
    return
}

func isTypeComparable[T any]() bool {
    return reflect.ValueOf(zVal[T]()).Comparable()
}
