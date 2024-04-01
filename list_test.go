package handy

import (
    "github.com/hsldymq/goiter"
    "slices"
    "testing"
)

func TestNewList(t *testing.T) {
    list := NewList[int]()
    if len(list.elems) != 0 {
        t.Fatalf("test NewList, create an empty list, len should be 0")
    }

    list = NewList(1, 2, 3)
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test NewList expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Merge(t *testing.T) {
    list := NewList[int]()
    l2 := NewList(1, 2, 3)
    l3 := NewList(4, 5, 6)
    l4 := NewList(7, 8, 9)
    list.Merge(l2, l3, l4)
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Merge expect: %v, actual: %v", expect, actual)
    }

    list = NewList(1, 2, 3)
    list.Merge()
    actual = []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Merge expect: %v, actual: %v", expect, actual)
    }
}

func TestList_MergeSlices(t *testing.T) {
    list := NewList[int]()
    list.MergeSlices([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.MergeSlices expect: %v, actual: %v", expect, actual)
    }

    list = NewList[int](1, 2, 3)
    list.MergeSlices()
    actual = []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.MergeSlices expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Remove(t *testing.T) {
    list := NewList[int]()
    list.Merge(NewList(1, 2, 3, 4, 5))
    list.Remove(3)
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Remove expect: %v, actual: %v", expect, actual)
    }

    type person struct {
        Name string
        Age  int
    }
    alice := &person{"Alice", 20}
    bob := &person{"Bob", 30}
    eve := &person{"Eve", 40}
    list2 := NewList[*person](alice, bob, eve)
    list2.Remove(bob)
    actual2 := []person{}
    for v := range list2.Iter() {
        actual2 = append(actual2, *v)
    }
    expect2 := []person{
        {"Alice", 20},
        {"Eve", 40},
    }
    if !slices.Equal(expect2, actual2) {
        t.Fatalf("test List.Remove expect: %v, actual: %v", expect, actual)
    }

    f1 := func() {}
    list3 := NewList[func()](f1, func() {}, func() {})
    list3.Remove(f1)
    if list3.Count() != 3 {
        t.Fatalf("test List.Remove, func is not comparable, Remove should has no effect")
    }
}

func TestList_RemoveAt(t *testing.T) {

}

func TestList_Pop(t *testing.T) {
    list := NewListFromIter[int](goiter.Range(1, 3).Seq())

    actual, ok := list.Pop()
    if !ok {
        t.Fatalf("test List.Pop, returned ok is not true")
    }
    expect := 3
    if expect != actual {
        t.Fatalf("test List.Pop expect: %v, actual: %v", expect, actual)
    }

    actual, ok = list.Pop()
    if !ok {
        t.Fatalf("test List.Pop, returned ok is not true")
    }
    expect = 2
    if expect != actual {
        t.Fatalf("test List.Pop expect: %v, actual: %v", expect, actual)
    }

    actual, ok = list.Pop()
    if !ok {
        t.Fatalf("test List.Pop, returned ok is not true")
    }
    expect = 1
    if expect != actual {
        t.Fatalf("test List.Pop expect: %v, actual: %v", expect, actual)
    }

    _, ok = list.Pop()
    if ok {
        t.Fatalf("test List.Pop, returned ok is not false")
    }
}

func TestList_Shift(t *testing.T) {
    list := NewList[int]()
    list.Add(1, 2, 3)

    actual, ok := list.Shift()
    if !ok {
        t.Fatalf("test List.Shift, returned ok is not true")
    }
    expect := 1
    if expect != actual {
        t.Fatalf("test List.Shift expect: %v, actual: %v", expect, actual)
    }

    actual, ok = list.Shift()
    if !ok {
        t.Fatalf("test List.Shift, returned ok is not true")
    }
    expect = 2
    if expect != actual {
        t.Fatalf("test List.Shift expect: %v, actual: %v", expect, actual)
    }

    actual, ok = list.Shift()
    if !ok {
        t.Fatalf("test List.Shift, returned ok is not true")
    }
    expect = 3
    if expect != actual {
        t.Fatalf("test List.Shift expect: %v, actual: %v", expect, actual)
    }

    _, ok = list.Shift()
    if ok {
        t.Fatalf("test List.Shift, returned ok is not false")
    }
}

func TestList_Get(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)

    actual, ok := list.Get(2)
    if !ok {
        t.Fatalf("test List.Get, returned ok is not true")
    }
    expect := 3
    if expect != actual {
        t.Fatalf("test List.Shift expect: %v, actual: %v", expect, actual)
    }

    _, ok = list.Get(-1)
    if ok {
        t.Fatalf("test List.Get, returned ok is not false")
    }

    _, ok = list.Get(5)
    if ok {
        t.Fatalf("test List.Get, returned ok is not false")
    }
}

func TestList_Find(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
    )

    p, ok := list.Find(func(p *person) bool {
        return p.Name == "Bob"
    })
    if !ok {
        t.Fatalf("test List.Find, returned ok should be true")
    }
    expect := person{"Bob", 30}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }

    _, ok = list.Find(func(p *person) bool {
        return p.Name == "Alex"
    })
    if ok {
        t.Fatalf("test List.Find, returned ok should be false")
    }
}

func TestList_FindOrDefault(t *testing.T) {

}

func TestList_FindLast(t *testing.T) {

}

func TestList_FindLastOrDefault(t *testing.T) {

}

func TestList_IndexOf(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    actual := list.IndexOf(3)
    expect := 2
    if expect != actual {
        t.Fatalf("test List.IndexOf expect: %v, actual: %v", expect, actual)
    }

    actual = list.IndexOf(6)
    expect = -1
    if expect != actual {
        t.Fatalf("test List.IndexOf expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Take(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5, 6, 7, 8)

    actual := []int{}
    for v := range list.Take(5).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Take expect: %v, actual: %v", expect, actual)
    }

    actual = []int{}
    for v := range list.Take(10).Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Take expect: %v, actual: %v", expect, actual)
    }

    actual = []int{}
    for v := range list.Take(-1).Iter() {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Take expect: %v, actual: %v", expect, actual)
    }
}

func TestList_TakeLast(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5, 6, 7, 8)

    // case 1
    actual := []int{}
    for v := range list.TakeLast(5).Iter() {
        actual = append(actual, v)
    }
    expect := []int{4, 5, 6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
    }

    // case 2
    actual = []int{}
    for v := range list.TakeLast(5).Iter() {
        actual = append(actual, v)
        if v == 6 {
            break
        }
    }
    expect = []int{4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
    }

    // case 3
    actual = []int{}
    for v := range list.TakeLast(10).Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
    }

    // case 4
    actual = []int{}
    for v := range list.TakeLast(-1).Iter() {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Skip(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5, 6, 7, 8)

    actual := []int{}
    for v := range list.Skip(5).Iter() {
        actual = append(actual, v)
    }
    expect := []int{6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Skip expect: %v, actual: %v", expect, actual)
    }

    actual = []int{}
    for v := range list.Skip(10).Iter() {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Skip expect: %v, actual: %v", expect, actual)
    }

    actual = []int{}
    for v := range list.Skip(-1).Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Skip expect: %v, actual: %v", expect, actual)
    }
}

func TestList_SkipLast(t *testing.T) {
    list := NewList(1, 2, 3, 4, 5, 6, 7, 8)

    // case 1
    actual := []int{}
    for v := range list.SkipLast(5).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
    }

    // case 2
    actual = []int{}
    for v := range list.SkipLast(4).Iter() {
        actual = append(actual, v)
        if v == 3 {
            break
        }
    }
    expect = []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
    }

    // case 3
    actual = []int{}
    for v := range list.SkipLast(8).Iter() {
        actual = append(actual, v)
    }
    expect = []int{}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
    }

    // case 4
    actual = []int{}
    for v := range list.SkipLast(-1).Iter() {
        actual = append(actual, v)
    }
    expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
    }
}
