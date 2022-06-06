package stream

import (
	"fmt"
)

type FilterFn[T any] func(T) bool

type FilterOperator[T any] struct {
	Operator[T]
}

func makeFilterOperator[T any](num_workers int, parent IStream[T], filter_fn FilterFn[T], name string) FilterOperator[T] {
	map_fn := func(data T) (T, bool) {
		return data, filter_fn(data)
	}

	return FilterOperator[T]{
		makeOperator(num_workers, parent, map_fn, fmt.Sprintf("Filter-%s", name), IntermediateType),
	}
}
