package stream

import (
	"fmt"
)

type MapFn[T any] func(T) T

type MapOperator[T any] struct {
	Operator[T]
}

func makeMapOperator[T any](num_workers int, parent IStream[T], map_fn MapFn[T], name string) FilterOperator[T] {
	map_fn1 := func(data T) (T, bool) {
		return map_fn(data), true
	}

	return FilterOperator[T]{
		makeOperator(num_workers, parent, map_fn1, nil, fmt.Sprintf("Filter-%s", name), StreamTypeIntermediate),
	}
}
