package handyCollection

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	JSONMarshalFunc   = json.Marshal
	JSONUnmarshalFunc = json.Unmarshal
)

func JSONMarshalAsMap[T any](c *GeneralCollection[T]) ([]byte, error) {
	return JSONMarshalFunc(c.items)
}

func JSONMarshalAsArray[T any](c *GeneralCollection[T]) ([]byte, error) {
	return JSONMarshalFunc(c.AsSlice())
}

func JSONUnmarshal[T any](data []byte, c *GeneralCollection[T], keepOriginalItems bool) error {
	d := strings.TrimSpace(string(data))
	if d[0] == '[' {
		l := make([]T, 0)
		if err := JSONUnmarshalFunc([]byte(d), &l); err != nil {
			return err
		}
		if !keepOriginalItems {
			c.Clear()
		}
		c.Add(l...)
	} else if d[0] == '{' {
		m := make(map[string]T)
		if err := JSONUnmarshalFunc([]byte(d), &m); err != nil {
			return err
		}
		if !keepOriginalItems {
			c.Clear()
		}
		c.MergeMaps(true, m)
	} else {
		return errors.New("data should be array or map")
	}

	return nil
}
