package handy

import (
	"slices"
	"testing"
)

func TestList_Take(t *testing.T) {
	list := NewListWithElems(1, 2, 3, 4, 5, 6, 7, 8)

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
	list := NewListWithElems(1, 2, 3, 4, 5, 6, 7, 8)

	actual := []int{}
	for v := range list.TakeLast(5).Iter() {
		actual = append(actual, v)
	}
	expect := []int{4, 5, 6, 7, 8}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
	}

	actual = []int{}
	for v := range list.TakeLast(10).Iter() {
		actual = append(actual, v)
	}
	expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test List.TakeLast expect: %v, actual: %v", expect, actual)
	}

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
	list := NewListWithElems(1, 2, 3, 4, 5, 6, 7, 8)

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
	list := NewListWithElems(1, 2, 3, 4, 5, 6, 7, 8)

	actual := []int{}
	for v := range list.SkipLast(5).Iter() {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
	}

	actual = []int{}
	for v := range list.SkipLast(8).Iter() {
		actual = append(actual, v)
	}
	expect = []int{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
	}

	actual = []int{}
	for v := range list.SkipLast(-1).Iter() {
		actual = append(actual, v)
	}
	expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test List.SkipLast expect: %v, actual: %v", expect, actual)
	}
}
