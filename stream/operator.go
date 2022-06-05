package stream

import (
	"sync"
)

type OperatorWorkerFn[In any, Out any] func(in <-chan In, out chan<- Out)

// The concrete implementation of a stream operator.
type Operator[In any, Out any] struct {
	num_workers int
	in          <-chan In
	out         chan Out
	worker_fn   OperatorWorkerFn[In, Out]
	name        string
}

func makeOperator[In any, Out any](num_workers int, in <-chan In, worker_fn OperatorWorkerFn[In, Out], name string) Operator[In, Out] {
	return Operator[In, Out]{
		num_workers: num_workers,
		in:          in,
		worker_fn:   worker_fn,
		out:         make(chan Out),
		name:        name,
	}
}

func (op *Operator[In, Out]) Iter() <-chan Out {
	return op.out
}

func (op *Operator[In, Out]) Workers(num_workers int) {
	op.num_workers = num_workers
}

func (op *Operator[In, Out]) Exec() {
	println("Running ", op.name)
	var wg sync.WaitGroup

	for i := 0; i < op.num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			op.worker_fn(op.in, op.out)
		}()
	}

	wg.Wait()
	close(op.out)
	println("Finished running ", op.name)
}

func (op *Operator[In, Out]) Filter(filter_fn FilterFn[Out]) IStream[Out] {
	filter := makeFilterOperator(op.num_workers, op.out, filter_fn, op.name)
	return &filter
}

func (op *Operator[In, Out]) ToSlice() []Out {
	RunDAG[Out](op)
	return toSlice(op.out)
}

func toSlice[T any](in <-chan T) []T {
	slice := make([]T, 0)
	for element := range in {
		slice = append(slice, element)
	}
	return slice
}
