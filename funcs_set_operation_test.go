//go:build goexperiment.rangefunc

package handy

import (
    "fmt"
    "github.com/hsldymq/goiter"
    "slices"
    "testing"
)

func TestUnion(t *testing.T) {
    // case 1
    case1S1 := goiter.SliceElem([]int{1, 2, 3, 5, 3})
    case1S2 := goiter.SliceElem([]int{3, 4, 5})
    case1E := union[int](newEnumerator(case1S1), newEnumerator(case1S2))
    case1Actual := make([]int, 0, 5)
    for v := range case1E.Iter() {
        case1Actual = append(case1Actual, v)
    }
    case1Expect := []int{1, 2, 3, 5, 4}
    if !slices.Equal(case1Expect, case1Actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", case1Expect, case1Actual))
    }

    // case 2
    type person struct {
        name string
    }
    case2S1 := goiter.SliceElem([]person{
        {"Alice"},
        {"Bob"},
    })
    case2S2 := goiter.SliceElem([]person{
        {"Bob"},
        {"Charlie"},
    })
    case2E := union[person](newEnumerator(case2S1), newEnumerator(case2S2))
    case2Actual := make([]person, 0, 5)
    for v := range case2E.Iter() {
        case2Actual = append(case2Actual, v)
    }
    case2Expect := []person{
        {"Alice"},
        {"Bob"},
        {"Charlie"},
    }
    if !slices.Equal(case2Expect, case2Actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", case2Expect, case2Actual))
    }

    // case 3: func() int is not comparable, so union will just concat the enumerators
    funcs := []func() int{
        func() int { return 1 },
        func() int { return 2 },
    }
    case3S1 := goiter.SliceElem([]func() int{funcs[0], funcs[1]})
    case3S2 := goiter.SliceElem([]func() int{funcs[1], funcs[0]})
    case3E := union[func() int](newEnumerator(case3S1), newEnumerator(case3S2))
    case3Actual := make([]func() int, 0, 2)
    for v := range case3E.Iter() {
        case3Actual = append(case3Actual, v)
    }
    case3Expect := []func() int{funcs[0], funcs[1], funcs[1], funcs[0]}
    if len(case3Actual) != 4 || &case3Actual[0] == &funcs[0] || &case3Actual[1] == &case3Expect[1] || &case3Actual[2] == &case3Expect[2] || &case3Actual[3] == &case3Expect[3] {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", case3Expect, case3Actual))
    }
}

func TestUnionBy(t *testing.T) {
    type person struct {
        name string
        job  string
    }
    s1 := goiter.SliceElem([]person{
        {"Alice", "Teacher"},
        {"Bob", "Doctor"},
    })
    s2 := goiter.SliceElem([]person{
        {"Bob", "Computer programmer"},
        {"Charlie", "Uber driver"},
    })
    e1 := unionBy[person](newEnumerator(s1), newEnumerator(s2), func(p person) any {
        return p.name
    })
    actual := make([]person, 0, 5)
    for v := range e1.Iter() {
        actual = append(actual, v)
    }
    expect := []person{
        {"Alice", "Teacher"},
        {"Bob", "Doctor"},
        {"Charlie", "Uber driver"},
    }
    if !slices.Equal(expect, actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
    }
}

func TestIntersect(t *testing.T) {
    // case 1
    case1S1 := goiter.SliceElem([]int{1, 2, 3, 5, 3})
    case1S2 := goiter.SliceElem([]int{5, 4, 3})
    case1E := intersect[int](newEnumerator(case1S1), newEnumerator(case1S2))
    case1Actual := make([]int, 0, 2)
    for v := range case1E.Iter() {
        case1Actual = append(case1Actual, v)
    }
    case1Expect := []int{3, 5}
    if !slices.Equal(case1Expect, case1Actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", case1Expect, case1Actual))
    }

    // case 2
    case2S1 := goiter.SliceElem([]*personWithID{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Eve"},
        {"4", "Charlie"},
        {"5", "Helen"},
    })
    case2S2 := goiter.SliceElem([]*personWithID{
        {"2", "David"},
        {"3", "Frank"},
        {"4", "George"},
        {"5", "Ivy"},
        {"6", "Lily"},
    })
    case2E := intersect[*personWithID](newEnumerator(case2S1), newEnumerator(case2S2))
    case2Actual := make([]personWithID, 0, 3)
    for v := range case2E.Iter() {
        case2Actual = append(case2Actual, *v)
        if v.ID == "4" {
            break
        }
    }
    case2Expect := []personWithID{
        {"2", "Bob"},
        {"3", "Eve"},
        {"4", "Charlie"},
    }
    if !slices.Equal(case2Expect, case2Actual) {
        t.Fatal(fmt.Sprintf("test intersect, expect: %v, actual: %v", case2Expect, case2Actual))
    }

    // case 3: func() int is not comparable, so intersect will just return enumerable that contains no elements
    funcs := []func() int{
        func() int { return 1 },
        func() int { return 2 },
    }
    case3S1 := goiter.SliceElem([]func() int{funcs[0], funcs[1]})
    case3S2 := goiter.SliceElem([]func() int{funcs[1], funcs[0]})
    case3E := intersect[func() int](newEnumerator(case3S1), newEnumerator(case3S2))
    case3Actual := make([]func() int, 0, 2)
    for v := range case3E.Iter() {
        case3Actual = append(case3Actual, v)
    }
    if len(case3Actual) != 0 {
        t.Fatal(fmt.Sprintf("expect no elements, got: %v", len(case3Actual)))
    }
}

func TestExcept(t *testing.T) {
    // case 1
    case1S1 := goiter.SliceElem([]int{1, 2, 3, 4, 5})
    case1S2 := goiter.SliceElem([]int{4, 5, 6, 7, 8})
    case1E := except[int](newEnumerator(case1S1), newEnumerator(case1S2))
    case1Actual := make([]int, 0, 3)
    for v := range case1E.Iter() {
        case1Actual = append(case1Actual, v)
    }
    case1Expect := []int{1, 2, 3}
    if !slices.Equal(case1Expect, case1Actual) {
        t.Fatal(fmt.Sprintf("expect: %v, actual: %v", case1Expect, case1Actual))
    }

    // case 2
    case2S1 := goiter.SliceElem([]*personWithID{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Eve"},
        {"4", "Charlie"},
        {"5", "Helen"},
    })
    case2S2 := goiter.SliceElem([]*personWithID{
        {"5", "Ivy"},
        {"6", "Lily"},
    })
    case2E := except[*personWithID](newEnumerator(case2S1), newEnumerator(case2S2))
    case2Actual := make([]personWithID, 0, 3)
    for v := range case2E.Iter() {
        case2Actual = append(case2Actual, *v)
        if v.ID == "3" {
            break
        }
    }
    case2Expect := []personWithID{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Eve"},
    }
    if !slices.Equal(case2Expect, case2Actual) {
        t.Fatal(fmt.Sprintf("test intersect, expect: %v, actual: %v", case2Expect, case2Actual))
    }

    // case 3: func() int is not comparable
    funcs := []func() int{
        func() int { return 1 },
        func() int { return 2 },
    }
    case3S1 := goiter.SliceElem([]func() int{funcs[0], funcs[1]})
    case3S2 := goiter.SliceElem([]func() int{funcs[0], funcs[1]})
    case3E := except[func() int](newEnumerator(case3S1), newEnumerator(case3S2))
    case3Actual := make([]func() int, 0, 2)
    for v := range case3E.Iter() {
        case3Actual = append(case3Actual, v)
    }
    if len(case3Actual) != 2 {
        t.Fatal(fmt.Sprintf("expect 2 elements, got: %v", len(case3Actual)))
    }
}
