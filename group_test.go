package handyCollection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGroupCollectionBy(t *testing.T) {
	c := NewGeneralCollection[int]().
		Add(1, 2, 3).
		AddWithKey(11, "11").
		AddWithKey(12, "12").
		AddWithKey(13, "13").
		Add(21, 22, 23)
	group := GroupCollectionBy(c, func(item int, _ int, _ string) string {
		return fmt.Sprintf("%d", item/10)
	})

	c1, _ := group.Find("0")
	assert.Equal(t, []int{1, 2, 3}, c1.AsSlice())

	c2, _ := group.Find("1")
	assert.Equal(t, []int{11, 12, 13}, c2.AsSlice())
	actual, _ := c2.FindByKey("11")
	assert.Equal(t, 11, actual)
	actual, _ = c2.FindByKey("12")
	assert.Equal(t, 12, actual)
	actual, _ = c2.FindByKey("13")
	assert.Equal(t, 13, actual)

	c3, _ := group.Find("2")
	assert.Equal(t, []int{21, 22, 23}, c3.AsSlice())

	_, found := group.Find("3")
	assert.False(t, found)
}

func TestGroup_SelfSortBy(t *testing.T) {
	g := NewGroup[int]().
		Set("2022", 22).
		Set("2010", 10).
		Set("2019", 19).
		Set("2018", 18).
		Set("2020", 20)

	g.SelfSortBy(func(iKey string, _ int, jKey string, _ int) bool {
		t1, _ := time.Parse("2006", iKey)
		t2, _ := time.Parse("2006", jKey)
		return t1.Year() <= t2.Year()
	})

	s := g.AsSlice()
	assert.Equal(t, []int{10, 18, 19, 20, 22}, s)
	for idx, val := range s {
		key, _ := g.KeyByIndex(idx)
		assert.Equal(t, fmt.Sprintf("20%d", val), key)
	}
}
