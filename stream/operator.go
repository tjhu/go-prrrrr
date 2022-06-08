package stream

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type OptionalMapFn[T any] func(data T) (T, bool)
type BatchGeneratorFn[T any] func(batch_size int) ([]T, bool)

// The concrete implementation of a stream operator.
type Operator[T any] struct {
	IStream[T]

	num_workers int
	parent      IStream[T]
	out         chan T
	batched_out chan []T
	worker_fn   OptionalMapFn[T]
	batch_gen   BatchGeneratorFn[T] // This doesn't belong here. It belongs to Source.
	name        string
	ty          StreamType
}

func makeOperator[T any](num_workers int, parent IStream[T], worker_fn OptionalMapFn[T], batch_gen BatchGeneratorFn[T], name string, ty StreamType) Operator[T] {
	return Operator[T]{
		num_workers: num_workers,
		parent:      parent,
		worker_fn:   worker_fn,
		batch_gen:   batch_gen,
		out:         make(chan T),
		batched_out: make(chan []T),
		name:        name,
		ty:          ty,
	}
}

func (op *Operator[T]) Out() <-chan T {
	return op.out
}

func (op *Operator[T]) BatchedOut() <-chan []T {
	return op.batched_out
}

func (op *Operator[T]) SetWorkers(num_workers int) IStream[T] {
	op.num_workers = num_workers
	return op
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

	for i := 0; i < op.num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			switch op.ty {
			case StreamTypeSource:
				// Dump the generator into the channel.
				var empty T
				for element, more := op.worker_fn(empty); more; element, more = op.worker_fn(empty) {
					op.out <- element
				}
			case StreamTypeIntermediate:
				for element := range op.parent.Out() {
					if new_element, more := op.worker_fn(element); more {
						op.out <- new_element
					}
				}
			}
		}()
	}

	wg.Wait()
	close(op.out)
	log.Info("Finished running stage: ", op.name)
}

func (op *Operator[T]) BatchExec(batch_size int) {
	log.Info("Running stage: ", op.name)
	var wg sync.WaitGroup

	for i := 0; i < op.num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			switch op.ty {
			case StreamTypeSource:
				// Dump the generator into the channel.
				for elements, more := op.batch_gen(batch_size); more; elements, more = op.batch_gen(batch_size) {
					op.batched_out <- elements
				}
			case StreamTypeIntermediate:
				var buffer []T
				var buffer2 []T // So like when `buffer` is filled up but we have more to process.
				buffer_size := 0
				for elements := range op.parent.BatchedOut() {
					if buffer == nil {
						buffer = elements
					} else {
						buffer2 = elements
					}

					for _, element := range elements {
						if element, more := op.worker_fn(element); more {
							// Append element.
							buffer[buffer_size] = element
							buffer_size++

							// Flush buffer if full.
							if buffer_size >= batch_size {
								op.batched_out <- buffer
								buffer_size = 0

								// Swap the second buffer in in case there're more stuff in `elements`
								buffer = buffer2
							}
						}
					}

					// Flush buffer if there's any remaining stuff.
					if buffer_size > 0 {
						op.batched_out <- buffer[:buffer_size]
						buffer_size = 0
					}
				}
			}
		}()
	}

	wg.Wait()
	close(op.out)
	close(op.batched_out)
	log.Info("Finished running stage: ", op.name)
}

func (op *Operator[T]) GetWorkerFn() OptionalMapFn[T] {
	return op.worker_fn
}

func (op *Operator[T]) Name() string {
	return op.name
}

func (op *Operator[T]) Filter(filter_fn FilterFn[T]) IStream[T] {
	filter := makeFilterOperator[T](op.num_workers, op, filter_fn, op.name)
	return &filter
}

func (op *Operator[T]) Map(map_fn MapFn[T]) IStream[T] {
	m := makeMapOperator[T](op.num_workers, op, map_fn, op.name)
	return &m
}

func (op *Operator[T]) ToSlice(optimizations ...OptimizationKind) []T {
	optimization := OptimizeKindUnoptimized
	for _, opt := range optimizations {
		optimization |= opt
	}

	// Run the operators
	RunDAG[T](op, optimization)

	// Ouput to slice
	if optimization&OptimizeKindBatching == 0 {
		return ChanToSlice(op.out)
	} else {
		return BatchChanToSlice(op.batched_out)
	}
}
