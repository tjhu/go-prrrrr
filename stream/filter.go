package stream

import "fmt"

type FilterFn[T any] func(T) bool

type FilterOperator[T any] struct {
	Operator[T, T]
}

func makeFilterOperator[T any](num_workers int, in <-chan T, filter_fn FilterFn[T], name string) FilterOperator[T] {
	println("yo")
	worker_fn := func(in <-chan T, out chan<- T) {
		println("bo")
		for element := range in {
			println(element)
			if filter_fn(element) {
				out <- element
			}
		}
	}

	return FilterOperator[T]{
		makeOperator(num_workers, in, worker_fn, fmt.Sprintf("Filter-%s", name)),
	}
}
