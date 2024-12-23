package handy

import (
    "github.com/hsldymq/goiter"
    "slices"
    "sort"
    "strings"
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
    list := NewListFromIter(goiter.Range(1, 3))

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
        &person{"Bob", 50},
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
    type person struct {
        Name string
        Age  int
    }
    list := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
        &person{"Bob", 50},
    )

    p := list.FindOrDefault(func(p *person) bool {
        return p.Name == "Bob"
    }, &person{"Default", 40})
    expect := person{"Bob", 30}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }

    p = list.FindOrDefault(func(p *person) bool {
        return p.Name == "Alex"
    }, &person{"Default", 40})
    expect = person{"Default", 40}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }
}

func TestList_FindLast(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
        &person{"Bob", 50},
    )

    p, ok := list.FindLast(func(p *person) bool {
        return p.Name == "Bob"
    })
    if !ok {
        t.Fatalf("test List.Find, returned ok should be true")
    }
    expect := person{"Bob", 50}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }

    _, ok = list.FindLast(func(p *person) bool {
        return p.Name == "Alex"
    })
    if ok {
        t.Fatalf("test List.Find, returned ok should be false")
    }
}

func TestList_FindLastOrDefault(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
        &person{"Bob", 50},
    )

    p := list.FindLastOrDefault(func(p *person) bool {
        return p.Name == "Bob"
    }, &person{"Default", 40})
    expect := person{"Bob", 50}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }

    p = list.FindLastOrDefault(func(p *person) bool {
        return p.Name == "Alex"
    }, &person{"Default", 80})
    expect = person{"Default", 80}
    if expect != *p {
        t.Fatalf("test List.Find expect: %v, actual: %v", expect, *p)
    }
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

func TestList_Contains(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    if !list.Contains(3) {
        t.Fatalf("test List.Contains, 3 should be in list")
    }

    if list.Contains(6) {
        t.Fatalf("test List.Contains, 6 should not be in list")
    }
}

func TestList_Sort(t *testing.T) {
    list := NewList[int](3, 1, 2, 5, 4)
    list.Sort(func(a, b int) int {
        return a - b
    })
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Sort expect: %v, actual: %v", expect, actual)
    }
}

func TestList_StableSort(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }

    list := NewList(
        person{"Alice", 20},
        person{"Bob", 30},
        person{"Eve", 40},
        person{"Bob", 50},
    )
    list.StableSort(func(a, b person) int {
        return strings.Compare(a.Name, b.Name)
    })
    actual := []person{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Alice", 20},
        {"Bob", 30},
        {"Bob", 50},
        {"Eve", 40},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.StableSort expect: %v, actual: %v", expect, actual)
    }
}

func TestList_FilterSelf(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    list.FilterSelf(func(v int) bool {
        return v%2 == 0
    })
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{2, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.FilterSelf expect: %v, actual: %v", expect, actual)
    }
}

func TestList_FilterTo(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    list2 := list.FilterTo(func(v int) bool {
        return v%2 == 0
    })
    if list == list2 {
        t.Fatalf("test List.FilterTo, should return a new list")
    }
    actual := []int{}
    for v := range list2.Iter() {
        actual = append(actual, v)
    }
    expect := []int{2, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.FilterTo expect: %v, actual: %v", expect, actual)
    }
}

func TestList_ShuffleTo(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    list2 := list.ShuffleTo()
    if list == list2 {
        t.Fatalf("test List.ShuffleTo, should return a new list")
    }
    actual := []int{}
    for v := range list2.Iter() {
        actual = append(actual, v)
    }
    sort.Ints(actual)
    if !slices.Equal([]int{1, 2, 3, 4, 5}, actual) {
        t.Fatalf("test List.ShuffleSelf, should be shuffled")
    }
}

func TestList_Clear(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    list.Clear()
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    if len(actual) != 0 {
        t.Fatalf("test List.Clear, list should be empty")
    }
}

func TestList_IterBackward(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    actual := []int{}
    for v := range list.IterBackward() {
        actual = append(actual, v)
    }
    expect := []int{5, 4, 3, 2, 1}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.IterBackward expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Any(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    result := list.Any(func(v int) bool {
        return v%2 == 0
    })
    if !result {
        t.Fatalf("test List.Any, should return true")
    }

    result = list.Any(func(v int) bool {
        return v == 6
    })
    if result {
        t.Fatalf("test List.Any, should return false")
    }
}

func TestList_All(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    result := list.All(func(v int) bool {
        return v%2 == 0
    })
    if result {
        t.Fatalf("test List.All, should return false")
    }

    result = list.All(func(v int) bool {
        return v < 6
    })
    if !result {
        t.Fatalf("test List.All, should return true")
    }
}

func TestList_Filter(t *testing.T) {
    list := NewList[int](1, 2, 3, 4, 5)
    e := list.Filter(func(v int) bool {
        return v%2 == 0
    })
    actual := []int{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []int{2, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Filter expect: %v, actual: %v", expect, actual)
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

func TestList_Distinct(t *testing.T) {
    // case 1
    list := NewList(1, 1, 2, 2, 3, 3, 4, 4, 5, 5)
    actual := []int{}
    for v := range list.Distinct().Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Distinct expect: %v, actual: %v", expect, actual)
    }

    // case 2
    f := func() {}
    list2 := NewList(f, f, f)
    actual2 := []func(){}
    for v := range list2.Distinct().Iter() {
        actual2 = append(actual2, v)
    }
    expect2 := 3
    if expect2 != len(actual2) {
        t.Fatalf("test List.Distinct, should return 3 elements")
    }

    // case 3
    list3 := NewList(
        &personWithID{ID: "1", Name: "Alice"},
        &personWithID{ID: "2", Name: "Bob"},
        &personWithID{ID: "3", Name: "Eve"},
        &personWithID{ID: "2", Name: "Robert"},
        &personWithID{ID: "5", Name: "Ellie"},
        &personWithID{ID: "3", Name: "Eva"},
    )
    actual3 := []personWithID{}
    for v := range list3.Distinct().Iter() {
        actual3 = append(actual3, *v)
        if v.Name == "Ellie" {
            break
        }
    }
    expect3 := []personWithID{
        {ID: "1", Name: "Alice"},
        {ID: "2", Name: "Bob"},
        {ID: "3", Name: "Eve"},
        {ID: "5", Name: "Ellie"},
    }
    if !slices.Equal(expect3, actual3) {
        t.Fatalf("test List.Distinct expect: %v, actual: %v", expect3, actual3)
    }
}

func TestList_DistinctBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }

    list := NewList(
        person{"Alice", 20},
        person{"Bob", 30},
        person{"Eve", 40},
        person{"Bob", 50},
        person{"Robert", 60},
    )
    e := list.DistinctBy(func(p person) any {
        return p.Name
    })
    actual := []person{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Alice", 20},
        {"Bob", 30},
        {"Eve", 40},
        {"Robert", 60},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.DistinctBy expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Union(t *testing.T) {
    list1 := NewList(1, 2, 3, 4, 5)
    list2 := NewList(3, 4, 5, 6, 7)
    actual := []int{}
    for v := range list1.Union(list2).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6, 7}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Union expect: %v, actual: %v", expect, actual)
    }

    p1 := NewList(
        &personWithID{"1", "Alice"},
        &personWithID{"2", "Bob"},
        &personWithID{"3", "Eve"},
        &personWithID{"4", "Charlie"},
    )
    p2 := NewList(
        &personWithID{"3", "David"},
        &personWithID{"4", "Frank"},
        &personWithID{"5", "George"},
        &personWithID{"6", "Helen"},
    )
    actual2 := []personWithID{}
    for v := range p1.Union(p2).Iter() {
        actual2 = append(actual2, *v)
    }
    expect2 := []personWithID{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Eve"},
        {"4", "Charlie"},
        {"5", "George"},
        {"6", "Helen"},
    }
    if !slices.Equal(expect2, actual2) {
        t.Fatalf("test List.Union expect: %v, actual: %v", expect2, actual2)
    }
}

func TestList_UnionBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }

    list1 := NewList(
        person{"Alice", 20},
        person{"Bob", 30},
        person{"Eve", 40},
    )
    list2 := NewList(
        person{"Bob", 60},
        person{"Eve", 70},
        person{"Frank", 50},
        person{"George", 60},
    )
    e := list1.UnionBy(list2, func(p person) any {
        return p.Name
    })
    actual := []person{}
    for v := range e.Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Alice", 20},
        {"Bob", 30},
        {"Eve", 40},
        {"Frank", 50},
        {"George", 60},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.UnionBy expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Intersect(t *testing.T) {
    list1 := NewList(1, 2, 3, 4, 5)
    list2 := NewList(3, 4, 5, 6, 7)
    actual := []int{}
    for v := range list1.Intersect(list2).Iter() {
        actual = append(actual, v)
    }
    expect := []int{3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Intersect expect: %v, actual: %v", expect, actual)
    }
}

func TestList_IntersectBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list1 := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
    )
    list2 := NewList(
        &person{"Bob", 60},
        &person{"Eve", 70},
        &person{"Frank", 50},
    )
    e := list1.IntersectBy(list2, func(p *person) any {
        return p.Name
    })
    actual := []person{}
    for v := range e.Iter() {
        actual = append(actual, *v)
    }
    expect := []person{
        {"Bob", 30},
        {"Eve", 40},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.IntersectBy expect: %v, actual: %v", expect, actual)
    }
}

func TestList_Except(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list1 := NewList(
        &person{"Alice", 20},
        &person{"Bob", 30},
        &person{"Eve", 40},
    )
    list2 := NewList(
        &person{"Bob", 60},
        &person{"Eve", 70},
        &person{"Frank", 50},
    )
    e := list1.ExceptBy(list2, func(p *person) any {
        return p.Name
    })
    actual := []person{}
    for v := range e.Iter() {
        actual = append(actual, *v)
    }
    expect := []person{
        {"Alice", 20},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.ExceptBy expect: %v, actual: %v", expect, actual)
    }
}

func TestList_ExceptBy(t *testing.T) {
    list1 := NewList(1, 2, 3, 4, 5)
    list2 := NewList(3, 4, 5, 6, 7)
    actual := []int{}
    for v := range list1.Except(list2).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Intersect expect: %v, actual: %v", expect, actual)
    }
}

func TestList_SequenceEqual(t *testing.T) {
    list1 := NewList(1, 2, 3, 4, 5)
    list2 := NewList(1, 2, 3, 4, 5)
    if !list1.SequenceEqual(list2) {
        t.Fatalf("test List.SequenceEqual failed, not true")
    }
}

func TestList_SequenceEqualBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }
    list1 := NewList(
        person{"Alice", 20},
        person{"Bob", 30},
        person{"Eve", 40},
    )
    list2 := NewList(
        person{"Alice", 50},
        person{"Bob", 60},
        person{"Eve", 70},
    )
    if !list1.SequenceEqualBy(list2, func(p person) any { return p.Name }) {
        t.Fatalf("test List.SequenceEqual failed, not true")
    }
}

func TestList_Concat(t *testing.T) {
    list := NewList(1, 2, 3)

    actual := []int{}
    for v := range list.Concat(NewList(4, 5, 6)).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5, 6}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.Concat expect: %v, actual: %v", expect, actual)
    }
}

func TestList_FirstOrDefault(t *testing.T) {
    list := NewList(1, 2, 3)
    expect := 1
    actual := list.FirstOrDefault(-1)
    if expect != actual {
        t.Fatalf("test List.FirstOrDefault expect: %d, actual: %d", expect, actual)
    }

    list = NewList[int]()
    expect = -1
    actual = list.FirstOrDefault(-1)
    if expect != actual {
        t.Fatalf("test List.FirstOrDefault expect: %d, actual: %d", expect, actual)
    }
}

func TestList_LastOrDefault(t *testing.T) {
    list := NewList(1, 2, 3)
    expect := 3
    actual := list.LastOrDefault(-1)
    if expect != actual {
        t.Fatalf("test List.LastOrDefault expect: %d, actual: %d", expect, actual)
    }

    list = NewList[int]()
    expect = -1
    actual = list.LastOrDefault(-1)
    if expect != actual {
        t.Fatalf("test List.LastOrDefault expect: %d, actual: %d", expect, actual)
    }
}

func TestList_OrderBy(t *testing.T) {
    type person struct {
        Name string
        Age  int
    }

    list := NewList(
        person{"Bob", 30},
        person{"Alice", 20},
        person{"Eve", 40},
    )
    actual := []person{}
    for v := range list.OrderBy(func(a, b person) int {
        return strings.Compare(a.Name, b.Name)
    }).Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Alice", 20},
        {"Bob", 30},
        {"Eve", 40},
    }
    if !slices.Equal(expect, actual) {
        t.Fatalf("test List.OrderBy expect: %v, actual: %v", expect, actual)
    }
}

type personWithID struct {
    ID   string
    Name string
}

func (p *personWithID) ComparingKey() any {
    return p.ID
}
