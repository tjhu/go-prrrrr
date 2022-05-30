package stream

import "sync"

// The concrete implementation of a stream operator.
type Operator[In any, Out any] struct {
	num_workers int
	in          *chan In
	out         *chan Out
	worker_fn   func(in *chan In, out *chan Out)
}

func makeOperator[In any, Out any](in *chan In, worker_fn func(in *chan In, out *chan Out)) Operator[In, Out] {
	out := make(chan Out)
	return Operator[In, Out]{
		num_workers: 1,
		in:          in,
		worker_fn:   worker_fn,
		out:         &out,
	}
}

func (op *Operator[In, Out]) Iter() *chan Out {
	return op.out
}

func (op *Operator[In, Out]) Workers(num_workers int) {
	op.num_workers = num_workers
}

func (op *Operator[In, Out]) Exec() {
	var wg sync.WaitGroup

	for i := 0; i < op.num_workers; i++ {
		go func() {
			defer wg.Done()
			op.worker_fn(op.in, op.out)
		}()
	}

	wg.Wait()
}

func (op *Operator[In, Out]) Filter(FilterFn[Out]) IStream[Out] {
	return &FilterOperator[Out]{}
}
