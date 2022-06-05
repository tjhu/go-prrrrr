package stream

type SliceInputOperator[T any] struct {
	Source[T]
}

func makeSliceInputOperator[T any](slice []T) SliceInputOperator[T] {
	worker_fn := func(out chan<- T) {
		for _, element := range slice {
			out <- element
		}
	}

	return SliceInputOperator[T]{
		makeSource(worker_fn, "Slice"),
	}
}

func Of[T any](elements ...T) IStream[T] {
	return OfSlice(elements)
}

func OfSlice[T any](slice []T) IStream[T] {
	op := makeSliceInputOperator(slice)
	return &op
}
