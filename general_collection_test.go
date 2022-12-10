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

func TestGeneralCollection_SelfSortBy(t *testing.T) {
	c := NewGeneralCollection[int](100, 200, 300, 400, 500, 600, 700, 800, 900)
	c.SelfSortBy(func(i, j int) bool {
		return i > j
	})
	assert.Equal(t, []int{900, 800, 700, 600, 500, 400, 300, 200, 100}, c.AsSlice())
}

func TestGeneralCollection_actualIndex(t *testing.T) {
	c := NewGeneralCollection[int](1, 2)

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
