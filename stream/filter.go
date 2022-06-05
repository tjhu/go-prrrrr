package stream

type FilterFn[T any] func(T) bool

type FilterOperator[T any] struct {
	Operator[T, T]
}

func makeFilterOperator[T any](num_workers int, in <-chan T, filter_fn FilterFn[T]) FilterOperator[T] {
	worker_fn := func(in <-chan T, out chan<- T) {
		for element := range in {
			if filter_fn(element) {
				out <- element
			}
		}
	}

	return FilterOperator[T]{
		makeOperator(num_workers, in, worker_fn),
	}
}
