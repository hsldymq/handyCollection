package handy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	c1 := NewGeneralCollection[string]().Add("a", "aa", "aaa", "aaaa").AddWithKey("aaaaa", "a5")
	c2 := Map(c1, func(item string) int {
		return len(item)
	})

	assert.Equal(t, []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, c1.AsSlice())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, c2.AsSlice())
	actual, _ := c2.FindByKey("a5")
	assert.Equal(t, 5, actual)
}

func TestReduce(t *testing.T) {
	c := NewGeneralCollection[string]().Add("a", "aa", "aaa", "aaaa", "aaaaa")
	actual := Reduce(c, func(item string, carry int) int {
		return len(item) + carry
	}, 0)
	assert.Equal(t, 15, actual)
}
