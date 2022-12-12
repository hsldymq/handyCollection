package handyCollection

import (
	"encoding/json"
	"errors"
	"strings"
)

// 使用方可以替换该方法以获得更高校的marshal 与 unmarshal性能
var (
	JSONMarshalFunc   = json.Marshal
	JSONUnmarshalFunc = json.Unmarshal
)

// JSONMarshalAsObject 将集合序列化为json对象
func JSONMarshalAsObject[T any](c *GeneralCollection[T]) ([]byte, error) {
	return JSONMarshalFunc(c.items)
}

// JSONMarshalAsArray 将集合序列化为json数组
func JSONMarshalAsArray[T any](c *GeneralCollection[T]) ([]byte, error) {
	return JSONMarshalFunc(c.AsSlice())
}

// JSONUnmarshal 将json数据unmarshal后存入集合中
// 支持json数组和对象两种类型，如果是对象，那么对象的key会作为集合中数据项的key
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
