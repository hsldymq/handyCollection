package handy

import (
    "slices"
    "testing"
)

func TestReduce(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5)
    actual := Reduce(list, 0, func(acc, each int) int {
        return acc + each
    })
    expect := 15
    if actual != expect {
        t.Fatalf("test Reduce, expect: %v, actual: %v", expect, actual)
    }
}

func TestScan(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5)
    e := Scan(list, 0, func(acc, each int) int {
        return acc + each
    })
    actual := []int{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 3, 6, 10, 15}
    if !slices.Equal(actual, expect) {
        t.Fatalf("test Scan, expect: %v, actual: %v", expect, actual)
    }
}
