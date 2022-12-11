package handyCollection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGeneralCollection_AddWithKey(t *testing.T) {
	c := NewGeneralCollection[int]()

	c.AddWithKey(1, "1")
	actual, found := c.FindByKey("1")
	assert.Equal(t, 1, actual)
	assert.True(t, found)

	c.AsSlice()
	c.AddWithKey(11, "1")
	actual, found = c.FindByKey("1")
	assert.Equal(t, 11, actual)
	assert.True(t, found)

	_, found = c.FindByKey("2")
	assert.False(t, found)
}

func TestGeneralCollection_Merge(t *testing.T) {
	c1 := NewGeneralCollection[string]().Add("1", "2").AddWithKey("3", "3")
	c2 := NewGeneralCollection[string]().Add("1", "2").AddWithKey("33", "3")

	c1.Merge(c2)
	assert.Equal(t, []string{"1", "2", "33", "1", "2"}, c1.AsSlice())
}

func TestGeneralCollection_MergeSlices(t *testing.T) {
	c := NewGeneralCollection[string]().
		Add("1", "2", "3").
		MergeSlices([]string{"a", "b", "c"}, []string{"x", "y", "z"})

	assert.Equal(t, []string{"1", "2", "3", "a", "b", "c", "x", "y", "z"}, c.AsSlice())
}

func TestGeneralCollection_MergeMaps(t *testing.T) {
	c := NewGeneralCollection[string]().
		AddWithKey("1", "1").
		AddWithKey("2", "2").
		AddWithKey("3", "3").
		MergeMaps(true, map[string]string{
			"1": "11",
			"2": "22",
			"4": "33",
		})
	assert.Equal(t, []string{"11", "22", "3", "33"}, c.AsSlice())
	actual, found := c.FindByKey("4")
	assert.Equal(t, "33", actual)
	assert.True(t, found)

	c.MergeMaps(false, map[string]string{
		"5": "55",
	})
	assert.Equal(t, []string{"11", "22", "3", "33", "55"}, c.AsSlice())
	_, found = c.FindByKey("5")
	assert.False(t, found)

}

func TestGeneralCollection_FindByIndex(t *testing.T) {
	c := NewGeneralCollection[string]().Add("1", "2", "3", "4")

	actual, found := c.FindByIndex(1)
	assert.Equal(t, "2", actual)
	assert.True(t, found)

	actual, found = c.FindByIndex(-1)
	assert.Equal(t, "4", actual)
	assert.True(t, found)

	actual, found = c.FindByIndex(9)
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_FindByKey(t *testing.T) {
	c := NewGeneralCollection[int]().
		AddWithKey(11, "1").
		AddWithKey(22, "2")

	actual, found := c.FindByKey("1")
	assert.Equal(t, 11, actual)
	assert.True(t, found)

	actual, found = c.FindByKey("5")
	assert.Equal(t, 0, actual)
	assert.False(t, found)
}

func TestGeneralCollection_HasKey(t *testing.T) {
	c := NewGeneralCollection[int]().AddWithKey(11, "1")

	assert.True(t, c.HasKey("1"))
	assert.False(t, c.HasKey("2"))
}

func TestGeneralCollection_KeyByIndex(t *testing.T) {
	c := NewGeneralCollection[int]().AddWithKey(11, "1")

	actual, found := c.KeyByIndex(0)
	assert.Equal(t, "1", actual)
	assert.True(t, found)

	actual, found = c.KeyByIndex(-1)
	assert.Equal(t, "1", actual)
	assert.True(t, found)

	actual, found = c.KeyByIndex(2)
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_IndexByKey(t *testing.T) {
	c := NewGeneralCollection[int]().AddWithKey(11, "1").AddWithKey(22, "2")

	actual, found := c.IndexByKey("2")
	assert.Equal(t, 1, actual)
	assert.True(t, found)

	actual, found = c.IndexByKey("3")
	assert.Equal(t, 0, actual)
	assert.False(t, found)
}

func TestGeneralCollection_RemoveByIndex(t *testing.T) {
	c := NewGeneralCollection[string]().Add("11", "22", "33", "44", "55")

	actual, found := c.RemoveByIndex(4)
	assert.Equal(t, "55", actual)
	assert.True(t, found)

	actual, found = c.RemoveByIndex(0)
	assert.Equal(t, "11", actual)
	assert.True(t, found)

	actual, found = c.RemoveByIndex(1)
	assert.Equal(t, "33", actual)
	assert.True(t, found)

	actual, found = c.RemoveByIndex(8)
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_RemoveByKey(t *testing.T) {
	c := NewGeneralCollection[string]().
		AddWithKey("11", "1").
		AddWithKey("22", "2").
		AddWithKey("33", "3").
		AddWithKey("44", "4")

	actual, found := c.RemoveByKey("2")
	assert.Equal(t, "22", actual)
	assert.True(t, found)

	actual, found = c.RemoveByKey("99")
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_Pop(t *testing.T) {
	c := NewGeneralCollection[string]().Add("11", "22").AddWithKey("33", "3")

	idx, _ := c.IndexByKey("3")
	assert.Equal(t, 2, idx)
	actual, found := c.Pop()
	assert.Equal(t, "33", actual)
	assert.True(t, found)
	assert.Equal(t, []string{"11", "22"}, c.AsSlice())
	_, found = c.IndexByKey("3")
	assert.False(t, found)

	actual, found = c.Pop()
	assert.Equal(t, "22", actual)
	assert.True(t, found)

	actual, found = c.Pop()
	assert.Equal(t, "11", actual)
	assert.True(t, found)

	actual, found = c.Pop()
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_Shift(t *testing.T) {
	c := NewGeneralCollection[string]().Add("11", "22", "33")

	c.AsSlice()
	actual, found := c.Shift()
	assert.Equal(t, "11", actual)
	assert.True(t, found)
	assert.Equal(t, []string{"22", "33"}, c.AsSlice())

	actual, found = c.Shift()
	assert.Equal(t, "22", actual)
	assert.True(t, found)

	actual, found = c.Shift()
	assert.Equal(t, "33", actual)
	assert.True(t, found)

	actual, found = c.Shift()
	assert.Equal(t, "", actual)
	assert.False(t, found)
}

func TestGeneralCollection_Count(t *testing.T) {
	c := NewGeneralCollection[int]().Add(1, 2)
	assert.Equal(t, 2, c.Count())

	c.AddWithKey(3, "3")
	assert.Equal(t, 3, c.Count())

	c.Pop()
	c.Pop()
	c.Pop()
	assert.Equal(t, 0, c.Count())
}

func TestGeneralCollection_FilterCount(t *testing.T) {
	c := NewGeneralCollection[int]().Add(1, 2, 3, 4, 5, 6)
	num := c.FilterCount(func(item int, idx int, key string) bool {
		return item%2 == 0
	})

	assert.Equal(t, 3, num)
}

func TestGeneralCollection_GroupCount(t *testing.T) {
	c := NewGeneralCollection[string]().Add("2022-01-01", "2020-01-01", "2022-12-12", "2019-01-01")
	g := c.GroupCount(func(item string, _ int, _ string) string {
		t, _ := time.Parse("2006-01-02", item)
		return fmt.Sprintf("%d", t.Year())
	}).AsMap()

	assert.Equal(t, 1, g["2019"])
	assert.Equal(t, 1, g["2020"])
	assert.Equal(t, 2, g["2022"])
	assert.Equal(t, 0, g["2023"])
}

func TestGeneralCollection_ForEach(t *testing.T) {
	c := NewGeneralCollection[int]().Add(1, 2, 3, 4, 5)
	sum := 0
	c.ForEach(func(item int, idx int, key string) {
		sum += item
	})
	assert.Equal(t, 15, sum)
}

func TestGeneralCollection_Every(t *testing.T) {
	c1 := NewGeneralCollection[int]().Add(1, 2, 3, 4, 5)
	assert.True(t, c1.Every(func(item int) bool {
		return item > 0
	}))

	c2 := NewGeneralCollection[int]().Add(1, 2, -1, 4, 5)
	assert.False(t, c2.Every(func(item int) bool {
		return item > 0
	}))

	c3 := NewGeneralCollection[int]()
	assert.True(t, c3.Every(func(item int) bool {
		return item > 0
	}))
}

func TestGeneralCollection_Some(t *testing.T) {
	c1 := NewGeneralCollection[int]().Add(-1, 2, -3, -4, -5)
	assert.True(t, c1.Some(func(item int) bool {
		return item > 0
	}))

	c2 := NewGeneralCollection[int]().Add(-1, -2, -3, -4, -5)
	assert.False(t, c2.Some(func(item int) bool {
		return item > 0
	}))

	c3 := NewGeneralCollection[int]()
	assert.False(t, c3.Some(func(item int) bool {
		return item > 0
	}))
}

func TestGeneralCollection_FilterBy(t *testing.T) {
	c1 := NewGeneralCollection[int]().
		Add(1, 2, 3, 4, 5, 6).
		AddWithKey(7, "7").
		AddWithKey(8, "8")
	c2 := c1.FilterBy(func(item int, _ int, _ string) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, c1.AsSlice())
	actual, found := c1.FindByKey("7")
	assert.Equal(t, 7, actual)
	assert.True(t, found)
	actual, found = c1.FindByKey("8")
	assert.Equal(t, 8, actual)
	assert.True(t, found)

	assert.Equal(t, []int{2, 4, 6, 8}, c2.AsSlice())
	actual, found = c2.FindByKey("7")
	assert.Equal(t, 0, actual)
	assert.False(t, found)
	actual, found = c2.FindByKey("8")
	assert.Equal(t, 8, actual)
	assert.True(t, found)
}

func TestGeneralCollection_SelfFilterBy(t *testing.T) {
	c1 := NewGeneralCollection[int]().
		Add(1, 2, 3, 4, 5, 6).
		AddWithKey(7, "7").
		AddWithKey(8, "8")
	c2 := c1.SelfFilterBy(func(item int, _ int, _ string) bool {
		return item%2 == 0
	})

	assert.Equal(t, c1, c2)
	assert.Equal(t, []int{2, 4, 6, 8}, c1.AsSlice())
}

func TestGeneralCollection_SortBy(t *testing.T) {
	c1 := NewGeneralCollection[string]().Add("a", "aa", "aaa", "aaaa").AddWithKey("aaaaa", "5a")
	c2 := c1.SortBy(func(a string, b string) bool {
		return len(a) > len(b)
	})

	assert.NotEqual(t, c1, c2)
	assert.Equal(t, []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, c1.AsSlice())
	assert.Equal(t, []string{"aaaaa", "aaaa", "aaa", "aa", "a"}, c2.AsSlice())

	actual, _ := c2.FindByKey("5a")
	assert.Equal(t, "aaaaa", actual)

	idx, found := c2.IndexByKey("5a")
	assert.Equal(t, 0, idx)
	assert.True(t, found)
}

func TestGeneralCollection_SelfSortBy(t *testing.T) {
	c := NewGeneralCollection[int]().Add(100, 200, 300, 400, 500, 600, 700, 800, 900)
	c.SelfSortBy(func(i, j int) bool {
		return i > j
	})
	assert.Equal(t, []int{900, 800, 700, 600, 500, 400, 300, 200, 100}, c.AsSlice())
}

func TestGeneralCollection_Shuffle(t *testing.T) {
	c1 := NewGeneralCollection[int]().Add(1)
	c2 := c1.Shuffle()

	assert.NotEqual(t, c1, c2)
}

func TestGeneralCollection_actualIndex(t *testing.T) {
	c := NewGeneralCollection[int]().Add(1, 2)

	idx, valid := c.actualIndex(1)
	assert.Equal(t, 1, idx)
	assert.True(t, valid)

	idx, valid = c.actualIndex(-1)
	assert.Equal(t, 1, idx)
	assert.True(t, valid)

	idx, valid = c.actualIndex(2)
	assert.Equal(t, 0, idx)
	assert.False(t, valid)

	idx, valid = c.actualIndex(-3)
	assert.Equal(t, 0, idx)
	assert.False(t, valid)
}
