package handyCollection

func Map[S any, D any](srcCollection *GeneralCollection[S], mapper func(item S) D) *GeneralCollection[D] {
	distCollection := NewGeneralCollection[D]()
	srcCollection.ForEach(func(item S, idx int, key string) {
		if srcCollection.isAutoGenKey(key) {
			distCollection.Add(mapper(item))
		} else {
			distCollection.AddWithKey(mapper(item), key)
		}
	})

	return distCollection
}

func Reduce[S any, D any](collection *GeneralCollection[S], reducer func(item S, carry D) D, init D) D {
	carry := init
	collection.ForEach(func(item S, _ int, _ string) {
		carry = reducer(item, carry)
	})

	return carry
}
