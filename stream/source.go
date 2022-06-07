package stream

import (
	"fmt"
)

type Generator[T any] func() (T, bool)

// Data source stream
type Source[T any] struct {
	Operator[T]
}

func makeSource[T any](generator Generator[T], name string) Source[T] {
	// An extra copy of argument here. Shouldn't impact the overall performance much.
	worker_fn := func(T) (T, bool) {
		return generator()
	}

	return Source[T]{
		makeOperator(1, nil, worker_fn, fmt.Sprintf("Source<%s>", name), StreamTypeSource),
	}
}
