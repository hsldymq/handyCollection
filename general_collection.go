package handyCollection

import (
	"github.com/google/uuid"
	"math/rand"
	"sort"
	"time"
)

// GeneralCollection 通用集合
type GeneralCollection[T any] struct {
	items             map[string]T
	orderedKeys       []string        // 用于维持数据的有序性
	autoGeneratedKeys map[string]bool // 用于标记哪些item的key是自动生成的，这些key不会在覆盖到新产生的集合中(例如 FilterBy, SortBy 这类方法所生成的新集合)
	keysIndex         map[string]int  // 用于通过key搜索其下标
	sliceCache        []T
	defaultItem       T
}

// NewGeneralCollection 创建GeneralCollection
func NewGeneralCollection[T any]() *GeneralCollection[T] {
	c := &GeneralCollection[T]{}
	c.defaultItem = c.items[""]
	c.init()

	return c
}

func (c *GeneralCollection[T]) init() {
	c.items = make(map[string]T)
	c.orderedKeys = make([]string, 0)
	c.autoGeneratedKeys = make(map[string]bool)
	c.keysIndex = nil
	c.sliceCache = nil
}

// Add 向集合中新增数据
func (c *GeneralCollection[T]) Add(items ...T) *GeneralCollection[T] {
	for _, each := range items {
		key := c.genKey()
		c.items[key] = each
		c.orderedKeys = append(c.orderedKeys, key)
		c.autoGeneratedKeys[key] = true
		c.setKeyIndex(key, len(c.orderedKeys)-1)
		c.tryAppendToSliceCache(each)
	}
	return c
}

// AddWithKey 向集合中新增数据，并与一个key关联,后续的操作中可以通过这个key获取到数据本身
// 如果集合中已经有数据关联了相同的key, 则会覆盖原有数据, 并保持数据所在的次序
func (c *GeneralCollection[T]) AddWithKey(item T, key string) *GeneralCollection[T] {
	_, hasKey := c.items[key]
	c.items[key] = item
	if !hasKey {
		c.orderedKeys = append(c.orderedKeys, key)
		c.setKeyIndex(key, len(c.orderedKeys)-1)
		c.tryAppendToSliceCache(item)
	} else {
		if idx, found := c.IndexByKey(key); found {
			c.trySetSliceCacheItem(idx, item)
		}
	}
	return c
}

// Merge 合并其他集合的数据到该集合中
func (c *GeneralCollection[T]) Merge(collections ...*GeneralCollection[T]) *GeneralCollection[T] {
	for _, each := range collections {
		for _, key := range each.orderedKeys {
			item := each.items[key]
			if each.isAutoGenKey(key) {
				c.Add(item)
			} else {
				c.AddWithKey(item, key)
			}
		}
	}
	return c
}

// MergeSlices 合并slice中的数据到该集合中
func (c *GeneralCollection[T]) MergeSlices(slices ...[]T) *GeneralCollection[T] {
	for _, each := range slices {
		for _, item := range each {
			c.Add(item)
		}
	}
	return c
}

// MergeMaps 合并map中的数据到该集合中
// keepKeys为true时, map的key会作为集合搜索键的key
func (c *GeneralCollection[T]) MergeMaps(keepKeys bool, maps ...map[string]T) *GeneralCollection[T] {
	for _, each := range maps {
		for key, item := range each {
			if keepKeys {
				c.AddWithKey(item, key)
			} else {
				c.Add(item)
			}
		}
	}
	return c
}

func (c *GeneralCollection[T]) Clear() *GeneralCollection[T] {
	c.init()
	return c
}

// FindByIndex return the item by given index
// if idx < 0, the index start counting from the end of the collection.
// if the second return value is true, it means the item is found, otherwise not found.
func (c *GeneralCollection[T]) FindByIndex(idx int) (T, bool) {
	idx, valid := c.actualIndex(idx)
	if !valid {
		return c.defaultItem, false
	}
	return c.FindByKey(c.orderedKeys[idx])
}

func (c *GeneralCollection[T]) FindByKey(key string) (T, bool) {
	d, ok := c.items[key]
	return d, ok
}

func (c *GeneralCollection[T]) HasKey(key string) bool {
	_, ok := c.items[key]
	return ok
}

func (c *GeneralCollection[T]) KeyByIndex(idx int) (string, bool) {
	idx, valid := c.actualIndex(idx)
	if !valid {
		return "", false
	}
	return c.orderedKeys[idx], true
}

func (c *GeneralCollection[T]) IndexByKey(key string) (int, bool) {
	if c.keysIndex == nil {
		c.keysIndex = make(map[string]int)
		for idx, k := range c.orderedKeys {
			c.keysIndex[k] = idx
		}
	}
	idx, ok := c.keysIndex[key]
	return idx, ok
}

// Count return the number of items in the collection
func (c *GeneralCollection[T]) Count() int {
	return len(c.orderedKeys)
}

// RemoveByIndex 根据索引移除数据项
// 允许idx为负,即从末尾回数
// 返回true即移除成功，false时代表给定的idx超出范围.
func (c *GeneralCollection[T]) RemoveByIndex(idx int) (T, bool) {
	idx, valid := c.actualIndex(idx)
	if !valid {
		return c.defaultItem, false
	}

	lastIdx := len(c.orderedKeys) - 1
	key := c.orderedKeys[idx]
	item := c.items[key]
	delete(c.items, key)
	delete(c.autoGeneratedKeys, key)
	if idx == 0 {
		c.orderedKeys = c.orderedKeys[1:]
		c.deleteKeyIndex(key)
		if c.sliceCache != nil {
			c.sliceCache = c.sliceCache[1:]
		}
	} else if idx == lastIdx {
		c.orderedKeys = c.orderedKeys[:lastIdx]
		c.deleteKeyIndex(key)
		if c.sliceCache != nil {
			c.sliceCache = c.sliceCache[:lastIdx]
		}
	} else {
		copy(c.orderedKeys[idx:], c.orderedKeys[idx+1:])
		c.orderedKeys = c.orderedKeys[:lastIdx]
		c.clearKeysIndex()
		c.clearSliceCache()
	}
	return item, true
}

// RemoveByKey 根据搜索key移除数据项
// 返回true即移除成功; false代表key不存在
func (c *GeneralCollection[T]) RemoveByKey(key string) (T, bool) {
	idx, found := c.IndexByKey(key)
	if !found {
		return c.defaultItem, false
	}
	return c.RemoveByIndex(idx)
}

// Pop remove and return the last item in the collection
// This is a shorthand method of RemoveByIndex(-1)
func (c *GeneralCollection[T]) Pop() (T, bool) {
	return c.RemoveByIndex(-1)
}

// Shift remove and return the first item in the collection
// This is a shorthand method of RemoveByIndex(0)
func (c *GeneralCollection[T]) Shift() (T, bool) {
	return c.RemoveByIndex(0)
}

func (c *GeneralCollection[T]) AsSlice() []T {
	if c.sliceCache == nil {
		c.sliceCache = make([]T, 0, len(c.orderedKeys))
		for _, key := range c.orderedKeys {
			c.tryAppendToSliceCache(c.items[key])
		}
	}
	return c.sliceCache
}

func (c *GeneralCollection[T]) AsMap() map[string]T {
	m := map[string]T{}
	for key, val := range c.items {
		m[key] = val
	}
	return m
}

// ForEach iterates over items and invokes iteratee for each item
func (c *GeneralCollection[T]) ForEach(iteratee func(item T, idx int, key string)) {
	for idx, key := range c.orderedKeys {
		iteratee(c.items[key], idx, key)
	}
}

// Every test every item in the collection by given testing function
// returns true if every test result is true or the collection is empty, otherwise false
func (c *GeneralCollection[T]) Every(tester func(item T) bool) bool {
	for _, item := range c.items {
		if !tester(item) {
			return false
		}
	}
	return true
}

// Some test items in the collection by given testing function
// if any item test result is true, it returns true, otherwise false
// if collection is empty, it returns false
func (c *GeneralCollection[T]) Some(tester func(item T) bool) bool {
	for _, item := range c.items {
		if tester(item) {
			return true
		}
	}
	return false
}

// GroupCount 根据给定的分组逻辑，计算出每个分组中数据项的数量, 拆分后的组的键由参数中的分组逻辑提供
func (c *GeneralCollection[T]) GroupCount(grouper func(item T, idx int, key string) string) *Group[int] {
	g := NewGroup[int]()
	for idx, key := range c.orderedKeys {
		item := c.items[key]
		groupKey := grouper(item, idx, key)
		count, _ := g.Find(groupKey)
		g.Set(groupKey, count+1)
	}
	return g
}

// FilterCount 根据指定的过滤器, 返回符合条件的数据项数量
func (c *GeneralCollection[T]) FilterCount(filter func(item T, idx int, key string) bool) int {
	count := 0
	for idx, key := range c.orderedKeys {
		item := c.items[key]
		if filter(item, idx, key) {
			count += 1
		}
	}
	return count
}

func (c *GeneralCollection[T]) FilterBy(matcher func(item T, idx int, key string) bool) *GeneralCollection[T] {
	newCollection := NewGeneralCollection[T]()
	for idx, key := range c.orderedKeys {
		item := c.items[key]
		if matcher(item, idx, key) {
			if c.isAutoGenKey(key) {
				newCollection.Add(item)
			} else {
				newCollection.AddWithKey(item, key)
			}
		}
	}
	return newCollection
}

func (c *GeneralCollection[T]) SelfFilterBy(matcher func(item T, idx int, key string) bool) *GeneralCollection[T] {
	newColl := c.FilterBy(matcher)
	c.init()
	c.items = newColl.items
	c.orderedKeys = newColl.orderedKeys
	c.autoGeneratedKeys = newColl.autoGeneratedKeys

	return c
}

// SortBy return a cloned collection with its items order sorted
func (c *GeneralCollection[T]) SortBy(less func(a T, b T) bool) *GeneralCollection[T] {
	collection := c.clone()
	collection.SelfSortBy(less)
	return collection
}

// SelfSortBy do in-place sort
func (c *GeneralCollection[T]) SelfSortBy(less func(a T, b T) bool) {
	sort.Slice(c.orderedKeys, func(i, j int) bool {
		return less(c.items[c.orderedKeys[i]], c.items[c.orderedKeys[j]])
	})
	c.clearKeysIndex()
	c.clearSliceCache()
}

// Shuffle return a cloned collection with its items order shuffled
func (c *GeneralCollection[T]) Shuffle() *GeneralCollection[T] {
	collection := c.clone()
	collection.SelfShuffle()
	return collection
}

// SelfShuffle do in-place shuffle
func (c *GeneralCollection[T]) SelfShuffle() {
	count := c.Count()
	if count < 2 {
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Fisher–Yates shuffle
	shuffleSlice := c.orderedKeys
	for i := count - 1; i > 0; i-- {
		idx := r.Intn(i + 1)
		shuffleSlice[idx], shuffleSlice[i] = shuffleSlice[i], shuffleSlice[idx]
	}
	c.clearSliceCache()
}

func (c *GeneralCollection[T]) actualIndex(idx int) (int, bool) {
	if idx < 0 {
		idx = len(c.orderedKeys) + idx
		if idx < 0 {
			return 0, false
		}
	} else if idx > len(c.orderedKeys)-1 {
		return 0, false
	}
	return idx, true
}

func (c *GeneralCollection[T]) clone() *GeneralCollection[T] {
	newCollection := NewGeneralCollection[T]()
	for _, key := range c.orderedKeys {
		item := c.items[key]
		if c.isAutoGenKey(key) {
			newCollection.Add(item)
		} else {
			newCollection.AddWithKey(item, key)
		}
	}
	return newCollection
}

func (c *GeneralCollection[T]) setKeyIndex(key string, idx int) {
	if c.keysIndex != nil {
		c.keysIndex[key] = idx
	}
}

func (c *GeneralCollection[T]) deleteKeyIndex(key string) {
	if c.keysIndex != nil {
		delete(c.keysIndex, key)
	}
}

func (c *GeneralCollection[T]) clearKeysIndex() {
	c.keysIndex = nil
}

func (c *GeneralCollection[T]) trySetSliceCacheItem(idx int, item T) {
	if c.sliceCache != nil && idx >= 0 && idx < len(c.sliceCache) {
		c.sliceCache[idx] = item
	}
}

func (c *GeneralCollection[T]) tryAppendToSliceCache(item T) {
	if c.sliceCache != nil {
		c.sliceCache = append(c.sliceCache, item)
	}
}

func (c *GeneralCollection[T]) clearSliceCache() {
	c.sliceCache = nil
}

func (c *GeneralCollection[T]) isAutoGenKey(key string) bool {
	_, ok := c.autoGeneratedKeys[key]
	return ok
}

func (c *GeneralCollection[T]) genKey() string {
	for {
		key := "_autogen_" + uuid.New().String()
		if _, ok := c.items[key]; !ok {
			return key
		}
	}
}
