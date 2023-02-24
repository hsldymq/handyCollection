package handyCollection

func Map[S any, D any](srcCollection *GeneralCollection[S], mapper func(item S) D) *GeneralCollection[D] {
	distCollection := NewGeneralCollection[D]()
	srcCollection.ForEach(func(each *ItemInfo[S]) {
		if srcCollection.isAutoGenKey(each.Key) {
			distCollection.Add(mapper(each.Item))
		} else {
			distCollection.AddWithKey(mapper(each.Item), each.Key)
		}
	})

	return distCollection
}

func Reduce[S any, D any](collection *GeneralCollection[S], reducer func(item S, carry D) D, init D) D {
	carry := init
	collection.ForEach(func(each *ItemInfo[S]) {
		carry = reducer(each.Item, carry)
	})

	return carry
}
