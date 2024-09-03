package handy

import (
    "fmt"
    "slices"
    "testing"
)

func TestJoin(t *testing.T) {
    type student struct {
        ID   string
        Name string
    }
    type studentMajor struct {
        ID    string
        Major string
    }

    students := newIterableFromSlice([]student{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Cindy"},
    })
    majors := newIterableFromSlice([]studentMajor{
        {"1", "Math"},
        {"1", "Physics"},
        {"2", "English"},
    })
    e := Join(students, majors, func(s student) string {
        return s.ID
    }, func(m studentMajor) string {
        return m.ID
    })
    type res struct {
        ID    string
        Name  string
        Major string
    }
    actual := []res{}
    for v := range e.Iter() {
        actual = append(actual, res{
            ID:    v.Outer.ID,
            Name:  v.Outer.Name,
            Major: v.Inner.Major,
        })
    }
    expect := []res{
        {"1", "Alice", "Math"},
        {"1", "Alice", "Physics"},
        {"2", "Bob", "English"},
    }
    if !slices.Equal(actual, expect) {
        t.Fatalf("test Join, expect: %v, actual: %v", expect, actual)
    }

    students = newIterableFromSlice([]student{})
    e = Join(students, majors, func(s student) string {
        return s.ID
    }, func(m studentMajor) string {
        return m.ID
    })
    if e.Count() != 0 {
        t.Fatalf("test Join, expect: 0, actual: %d", e.Count())
    }
}

func TestJoinAs(t *testing.T) {
    type student struct {
        ID   string
        Name string
    }
    type studentMajor struct {
        ID    string
        Major string
    }

    type res struct {
        ID    string
        Name  string
        Major string
    }
    students := newIterableFromSlice([]student{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Cindy"},
    })
    majors := newIterableFromSlice([]studentMajor{
        {"1", "Math"},
        {"1", "Physics"},
        {"2", "English"},
    })
    e := JoinAs(students, majors, func(s student) string {
        return s.ID
    }, func(m studentMajor) string {
        return m.ID
    }, func(outer student, inner studentMajor) res {
        return res{
            ID:    outer.ID,
            Name:  outer.Name,
            Major: inner.Major,
        }
    })
    actual := []res{}
    for v := range e.Iter() {
        actual = append(actual, v)
        if v.ID == "2" {
            break
        }
    }
    expect := []res{
        {"1", "Alice", "Math"},
        {"1", "Alice", "Physics"},
        {"2", "Bob", "English"},
    }
    if !slices.Equal(actual, expect) {
        t.Fatalf("test JoinAs, expect: %v, actual: %v", expect, actual)
    }
}

func TestGroupJoin(t *testing.T) {
    type student struct {
        ID   string
        Name string
    }
    type studentMajor struct {
        ID    string
        Major string
    }

    students := newIterableFromSlice([]student{
        {"1", "Alice"},
        {"2", "Bob"},
        {"3", "Cindy"},
    })
    majors := newIterableFromSlice([]studentMajor{
        {"1", "Math"},
        {"1", "Physics"},
        {"2", "English"},
    })
    e := GroupJoin(students, majors, func(s student) string {
        return s.ID
    }, func(m studentMajor) string {
        return m.ID
    })
    type res struct {
        ID     string
        Name   string
        Majors string
    }
    actual := []res{}
    for v := range e.Iter() {
        r := res{
            ID:   v.Outer.ID,
            Name: v.Outer.Name,
        }
        for s := range v.Inner.Iter() {
            if r.Majors != "" {
                r.Majors += fmt.Sprintf(",%s", s.Major)
            } else {
                r.Majors = s.Major
            }
        }
        actual = append(actual, r)
        if r.ID == "2" {
            break
        }
    }
    expect := []res{
        {"1", "Alice", "Math,Physics"},
        {"2", "Bob", "English"},
    }
    if !slices.Equal(actual, expect) {
        t.Fatalf("test GroupJoin, expect: %v, actual: %v", expect, actual)
    }

    students = newIterableFromSlice([]student{})
    e = GroupJoin(students, majors, func(s student) string {
        return s.ID
    }, func(m studentMajor) string {
        return m.ID
    })
    if e.Count() != 0 {
        t.Fatalf("test GroupJoin, expect: 0, actual: %d", e.Count())
    }
}
