package handy

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONMarshalAsArray(t *testing.T) {
	c := NewGeneralCollection[string]().Add("1", "2", "3").AddWithKey("4", "44")
	data, err := JSONMarshalAsArray(c)
	assert.NoError(t, err)
	assert.Equal(t, `["1","2","3","4"]`, string(data))
}

func TestJSONMarshalAsObject(t *testing.T) {
	c := NewGeneralCollection[string]().
		AddWithKey("11", "1").
		AddWithKey("22", "2").
		AddWithKey("33", "3").
		AddWithKey("44", "4")

	data, err := JSONMarshalAsObject(c)
	assert.NoError(t, err)

	m := map[string]string{}
	_ = json.Unmarshal(data, &m)
	assert.Equal(t, "11", m["1"])
	assert.Equal(t, "22", m["2"])
	assert.Equal(t, "33", m["3"])
	assert.Equal(t, "44", m["4"])
}

func TestJSONUnmarshal(t *testing.T) {
	s1 := "   [1,2,3,4]   "
	c1 := NewGeneralCollection[int]()
	c1.Add(9, 8, 7, 6)

	err := JSONUnmarshal[int]([]byte(s1), c1, true)
	assert.NoError(t, err)
	assert.Equal(t, []int{9, 8, 7, 6, 1, 2, 3, 4}, c1.AsSlice())

	_ = JSONUnmarshal[int]([]byte(s1), c1, false)
	assert.Equal(t, []int{1, 2, 3, 4}, c1.AsSlice())

	s2 := `{"a":1,"b":2,"c":3}`
	c2 := NewGeneralCollection[int]()
	c2.AddWithKey(9, "x").
		AddWithKey(8, "y").
		AddWithKey(7, "z")
	err = JSONUnmarshal[int]([]byte(s2), c2, true)
	assert.NoError(t, err)
	actual, _ := c2.FindByKey("a")
	assert.Equal(t, 1, actual)
	actual, _ = c2.FindByKey("b")
	assert.Equal(t, 2, actual)
	actual, _ = c2.FindByKey("c")
	assert.Equal(t, 3, actual)
	actual, _ = c2.FindByKey("x")
	assert.Equal(t, 9, actual)
	actual, _ = c2.FindByKey("y")
	assert.Equal(t, 8, actual)
	actual, _ = c2.FindByKey("z")
	assert.Equal(t, 7, actual)

	_ = JSONUnmarshal[int]([]byte(s2), c2, false)
	actual, _ = c2.FindByKey("a")
	assert.Equal(t, 1, actual)
	actual, _ = c2.FindByKey("b")
	assert.Equal(t, 2, actual)
	actual, _ = c2.FindByKey("c")
	assert.Equal(t, 3, actual)
	_, found := c2.FindByKey("x")
	assert.False(t, found)

	s3 := "1"
	c3 := NewGeneralCollection[int]()
	err = JSONUnmarshal[int]([]byte(s3), c3, true)
	assert.Error(t, err)
}
