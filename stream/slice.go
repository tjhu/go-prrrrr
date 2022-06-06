package stream

type SliceInputOperator[T any] struct {
	Source[T]
}

func makeSliceInputOperator[T any](slice []T) SliceInputOperator[T] {
	index := 0
	generator := func() (T, bool) {
		var value T
		if index >= len(slice) {
			return value, false
		}
		value = slice[index]
		index++
		return value, true
	}

	return SliceInputOperator[T]{
		makeSource(generator, "Slice"),
	}
}

func Of[T any](elements ...T) IStream[T] {
	return OfSlice(elements)
}

func OfSlice[T any](slice []T) IStream[T] {
	op := makeSliceInputOperator(slice)
	return &op
}
