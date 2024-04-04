package handy

import (
    "cmp"
    "github.com/hsldymq/goiter"
    "slices"
    "testing"
)

func TestOrder(t *testing.T) {
    e := NewEnumerator(goiter.SliceElem([]int{3, 2, 1}))
    actual := []int{}
    for v := range Order[int](e).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Order, expect: %v, actual: %v", expect, actual)
    }

    e = NewEnumerator(goiter.SliceElem([]int{1, 2, 3}))
    actual = []int{}
    for v := range Order[int](e, true).Iter() {
        actual = append(actual, v)
    }
    expect = []int{3, 2, 1}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Order, expect: %v, actual: %v", expect, actual)
    }
}

func TestOrderBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    e := NewEnumerator(goiter.SliceElem([]person{
        {"Alice", 20},
        {"Charlie", 22},
        {"Bob", 18},
    }))
    actual := []person{}
    for v := range orderBy(e, func(a, b person) int { return cmp.Compare(a.Age, b.Age) }).Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Bob", 18},
        {"Alice", 20},
        {"Charlie", 22},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test orderBy, expect: %v, actual: %v", expect, actual)
    }
}
