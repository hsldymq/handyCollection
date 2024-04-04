package handy

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "slices"
    "testing"
)

func TestTransform(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5)
    e := Transform(list, func(each int) string {
        return fmt.Sprintf("%d", each*2)
    })
    actual := []string{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []string{"2", "4", "6", "8", "10"}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Transform, expect: %v, actual: %v", expect, actual)
    }
}

func TestTransformExpand(t *testing.T) {
    list := NewList([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
    e := TransformExpand(list, func(each []int) Iterable[int] {
        return newEnumerator(goiter.SliceElem(each))
    })
    actual := []int{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test TransformExpand, expect: %v, actual: %v", expect, actual)
    }

    for _ = range e.Iter() {
        break
    }
}

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
