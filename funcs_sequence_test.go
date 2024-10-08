package handy

import (
    "github.com/hsldymq/goiter"
    "testing"
)

func TestSequenceEqual(t *testing.T) {
    // case 1
    e1 := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e2 := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    if !sequenceEqual[int](e1, e2) {
        t.Fatalf("test sequenceEqual failed, expect true")
    }

    // case 2
    e1 = NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5, 6}))
    e2 = NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    if sequenceEqual[int](e1, e2) {
        t.Fatalf("test sequenceEqual failed, expect false")
    }

    // case 3
    e1 = NewEnumerator(goiter.SliceElems([]int{1, 2, 0, 4, 5}))
    e2 = NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    if sequenceEqual[int](e1, e2) {
        t.Fatalf("test sequenceEqual failed, expect false")
    }

    // case 4
    ep1 := NewEnumerator(goiter.SliceElems([]*personWithID{
        {ID: "1", Name: "Alice"},
        {ID: "2", Name: "Bob"},
        {ID: "3", Name: "Eve"},
    }))
    ep2 := NewEnumerator(goiter.SliceElems([]*personWithID{
        {ID: "1", Name: "Eve"},
        {ID: "2", Name: "Alice"},
        {ID: "3", Name: "Bob"},
    }))
    if !sequenceEqual[*personWithID](ep1, ep2) {
        t.Fatalf("test sequenceEqual failed, expect true")
    }

    // case 5
    f := func() {}
    ef1 := NewEnumerator(goiter.SliceElems([]func(){f}))
    ef2 := NewEnumerator(goiter.SliceElems([]func(){f}))
    if sequenceEqual[func()](ef1, ef2) {
        t.Fatalf("test sequenceEqual failed, expect false")
    }
}
