package handy

import (
    "slices"
    "testing"
)

func TestToList(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5)
    newList := ToList[int](list.Filter(func(each int) bool {
        return each%2 == 0
    }))
    actual := []int{}
    for v := range newList.Iter() {
        actual = append(actual, v)
    }
    expect := []int{2, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test ToList, expect: %v, actual: %v", expect, actual)
    }
}
