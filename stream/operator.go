package stream

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type OperatorWorkerFn[T any] func(in <-chan T, out chan<- T)

// The concrete implementation of a stream operator.
type Operator[T any] struct {
	num_workers int
	parent      IStream[T]
	out         chan T
	worker_fn   OperatorWorkerFn[T]
	name        string
	ty          StreamType
}

func makeOperator[T any](num_workers int, parent IStream[T], worker_fn OperatorWorkerFn[T], name string, ty StreamType) Operator[T] {
	return Operator[T]{
		num_workers: num_workers,
		parent:      parent,
		worker_fn:   worker_fn,
		out:         make(chan T),
		name:        name,
		ty:          ty,
	}
}

func (op *Operator[T]) Out() <-chan T {
	return op.out
}

func (op *Operator[T]) Workers(num_workers int) {
	op.num_workers = num_workers
}

func (op *Operator[T]) Parent() IStream[T] {
	return op.parent
}

func (op *Operator[T]) Type() StreamType {
	return op.ty
}

func (op *Operator[T]) Exec() {
	log.Info("Running stage: ", op.name)
	var wg sync.WaitGroup
	var in <-chan T
	if op.parent != nil {
		in = op.parent.Out()
	}

	for i := 0; i < op.num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			op.worker_fn(in, op.out)
		}()
	}

	wg.Wait()
	close(op.out)
	log.Info("Finished running stage: ", op.name)
}

func (op *Operator[T]) Filter(filter_fn FilterFn[T]) IStream[T] {
	filter := makeFilterOperator[T](op.num_workers, op, filter_fn, op.name)
	return &filter
}

func (op *Operator[T]) ToSlice() []T {
	RunDAG[T](op)
	return toSlice(op.out)
}

func toSlice[T any](in <-chan T) []T {
	slice := make([]T, 0)
	for element := range in {
		slice = append(slice, element)
	}
	return slice
}
