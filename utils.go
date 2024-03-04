package handyCollection

func shallowEqual(v1, v2 any) bool {
	return v1 == v2
}

func zeroVal[T any]() T {
	var v T
	return v
}
