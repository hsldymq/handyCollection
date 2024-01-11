package handyCollection

type Sortable[T any] interface {
}

type ICollection[T any] interface {
	Count() int
	Contains(T) bool
	Add(T)
	Remove(T) bool
}

type IList[T any] interface {
	ICollection[T]

	Get(index int) (T, bool)
	IndexOf(item T) int
	Insert(item T) error
	RemoveAt(index int) error
	Clear()

	Sort(comparer func(a T, b T) bool) IList[T]
	SelfSort(comparer func(a T, b T) bool)
	Shuffle() IList[T]
	SelfShuffle() IList[T]
}
