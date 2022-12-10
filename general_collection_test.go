package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneralCollection_AddWithKey(t *testing.T) {
	c := NewGeneralCollection[int]()

	c.AddWithKey(1, "1")
	actual, found := c.FindByKey("1")
	assert.Equal(t, 1, actual)
	assert.True(t, found)

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
	c1 := NewGeneralCollection[string]().
		AddWithKey("1", "1").
		AddWithKey("2", "2").
		AddWithKey("3", "3").
		MergeMaps(true, map[string]string{
			"1": "11",
			"2": "22",
			"4": "33",
		})
	assert.Equal(t, []string{"11", "22", "3", "33"}, c1.AsSlice())
}

func TestGeneralCollection_SelfSortBy(t *testing.T) {
	c := NewGeneralCollection[int]().Add(100, 200, 300, 400, 500, 600, 700, 800, 900)
	c.SelfSortBy(func(i, j int) bool {
		return i > j
	})
	assert.Equal(t, []int{900, 800, 700, 600, 500, 400, 300, 200, 100}, c.AsSlice())
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
