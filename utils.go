package handy

import (
	"reflect"
)

func shallowEqual(v1, v2 any) bool {
	return v1 == v2
}

func zVal[T any]() T {
	var v T
	return v
}

func isTypeComparable[T any]() bool {
	var v T
	return reflect.ValueOf(v).Comparable()
}
