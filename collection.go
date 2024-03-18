package handy

type Collection[T any] interface {
	Add(items ...T) Collection[T]
	AddWithKey(item T, key string) Collection[T]
	Merge(collections ...Collection[T]) Collection[T]
	MergeSlices(slices ...[]T) Collection[T]
	MergeMaps(keepKeys bool, maps ...map[string]T) Collection[T]
	Clear() Collection[T]
	FindByIndex(idx int) (T, bool)
	FindByKey(key string) (T, bool)
	HasKey(key string) bool
	KeyByIndex(idx int) (string, bool)
	IndexByKey(key string) (int, bool)
	Count() int
	RemoveByIndex(idx int) (T, bool)
	RemoveByKey(key string) (T, bool)
	Pop() (T, bool)
	Shift() (T, bool)
	AsSlice() []T
	AsMap() map[string]T
	ForEach(iteratee func(each *ItemInfo[T]))
	Every(tester func(item T) bool) bool
	Some(tester func(item T) bool) bool
	GroupCount(grouper func(each *ItemInfo[T]) string) *Group[int]
	FilterCount(filter func(each *ItemInfo[T]) bool) int
	Filter(filter func(each *ItemInfo[T]) bool) Collection[T]
	SelfFilter(filter func(each *ItemInfo[T]) bool) Collection[T]
	Sort(less func(a *ItemInfo[T], b *ItemInfo[T]) bool) Collection[T]
	SelfSort(less func(a *ItemInfo[T], b *ItemInfo[T]) bool) Collection[T]
	Shuffle() Collection[T]
	SelfShuffle() Collection[T]
}

// ItemInfo contains the information of an item in the collection
type ItemInfo[T any] struct {
	Item         T
	Index        int
	Key          string
	IsAutoGenKey bool
}
