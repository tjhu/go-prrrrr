package stream

import "fmt"

type SourceWorker[T any] func(out chan<- T)

// Data source stream
type Source[T any] struct {
	Operator[Void, T]
}

func makeSource[T any](gen SourceWorker[T], name string) Source[T] {
	worker_fn := func(_ <-chan Void, out chan<- T) {
		gen(out)
	}

	return Source[T]{
		makeOperator(1, nil, worker_fn, fmt.Sprintf("Source-%s", name)),
	}
}
