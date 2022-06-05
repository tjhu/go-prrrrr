package stream

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type FilterFn[T any] func(T) bool

type FilterOperator[T any] struct {
	Operator[T]
}

func makeFilterOperator[T any](num_workers int, parent IStream[T], filter_fn FilterFn[T], name string) FilterOperator[T] {
	worker_fn := func(in <-chan T, out chan<- T) {
		log.Info("Start filter worker: ", name)
		for element := range in {
			if filter_fn(element) {
				out <- element
			}
		}
		log.Info("Done filter worker: ", name)
	}

	return FilterOperator[T]{
		makeOperator(num_workers, parent, worker_fn, fmt.Sprintf("Filter-%s", name), IntermediateType),
	}
}
