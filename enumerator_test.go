package handy

import (
    "cmp"
    "github.com/hsldymq/goiter"
    "slices"
    "testing"
)

func TestEnumerator_Iter(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    actual := []int{}
    for v := range list.Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5}
    if !slices.Equal(actual, expect) {
        t.Fatalf("test Enumerator.Iter, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Count(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    actual := list.Count()
    expect := 5
    if actual != expect {
        t.Fatalf("test Enumerator.Count, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Any(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    actual := list.Any(func(v int) bool {
        return v == 3
    })
    if !actual {
        t.Fatalf("test Enumerator.Any, expect: %v, actual: %v", true, false)
    }

    list = NewEnumerator(goiter.SliceElems([]int{1, 2, 9, 4, 5}))
    actual = list.Any(func(v int) bool {
        return v == 3
    })
    if actual {
        t.Fatalf("test Enumerator.Any, expect: %v, actual: %v", false, true)
    }
}

func TestEnumerator_All(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    actual := list.All(func(v int) bool {
        return v > 0
    })
    if !actual {
        t.Fatalf("test Enumerator.All, expect: %v, actual: %v", true, false)
    }

    list = NewEnumerator(goiter.SliceElems([]int{1, -1, 3, 4, 5}))
    actual = list.All(func(v int) bool {
        return v > 0
    })
    if actual {
        t.Fatalf("test Enumerator.All, expect: %v, actual: %v", false, true)
    }
}

func TestEnumerator_Filter(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e := list.Filter(func(v int) bool {
        return v%2 == 0
    }).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{2, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Filter, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Take(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e := list.Take(3).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Take, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_TakeLast(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e := list.TakeLast(3).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.TakeLast, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Skip(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e := list.Skip(3).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Skip, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_SkipLast(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 4, 5}))
    e := list.SkipLast(3).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.SkipLast, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Distinct(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 2, 1}))
    e := list.Distinct().Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Distinct, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_DistinctBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3, 2, 1}))
    e := list.DistinctBy(func(v int) any {
        return v % 2
    }).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.DistinctBy, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Union(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.Union(target).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Union, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_UnionBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.UnionBy(target, func(v int) any {
        return v % 2
    }).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.UnionBy, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Intersect(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.Intersect(target).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Intersect, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_IntersectBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.IntersectBy(target, func(v int) any {
        return v % 2
    }).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1, 2}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.IntersectBy, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_Except(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.Except(target).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{1}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Except, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_ExceptBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{2, 3, 4}))
    e := list.ExceptBy(target, func(v int) any {
        return v % 2
    }).Iter()
    actual := []int{}
    for v := range e {
        actual = append(actual, v)
    }
    expect := []int{}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.ExceptBy, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_SequenceEqual(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    actual := list.SequenceEqual(target)
    if !actual {
        t.Fatalf("test Enumerator.SequenceEqual, expect: %v, actual: %v", true, false)
    }
}

func TestEnumerator_SequenceEqualBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    actual := list.SequenceEqualBy(target, func(v int) any {
        return v
    })
    if !actual {
        t.Fatalf("test Enumerator.SequenceEqualBy, expect: %v, actual: %v", true, false)
    }
}

func TestEnumerator_Concat(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{1, 2, 3}))
    target := NewEnumerator(goiter.SliceElems([]int{4, 5}))
    actual := []int{}
    for v := range list.Concat(target).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3, 4, 5}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.Concat, expect: %v, actual: %v", expect, actual)
    }
}

func TestEnumerator_OrderBy(t *testing.T) {
    list := NewEnumerator(goiter.SliceElems([]int{3, 2, 1}))
    actual := []int{}
    for v := range list.OrderBy(func(a, b int) int {
        return cmp.Compare(a, b)
    }).Iter() {
        actual = append(actual, v)
    }
    expect := []int{1, 2, 3}
    if !slices.Equal(expect, actual) {
        t.Fatalf("test Enumerator.OrderBy, expect: %v, actual: %v", expect, actual)
    }
}
