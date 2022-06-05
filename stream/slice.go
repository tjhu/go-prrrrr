package stream

import log "github.com/sirupsen/logrus"

type SliceInputOperator[T any] struct {
	Source[T]
}

func makeSliceInputOperator[T any](slice []T) SliceInputOperator[T] {
	worker_fn := func(out chan<- T) {
		log.Info("Generating from slice")
		for _, element := range slice {
			out <- element
		}
		log.Info("Done generating from slice")
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
