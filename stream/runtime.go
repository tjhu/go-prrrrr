// This file defines the runtime for running the pipeline.

package stream

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// OptimizationKind
type OptimizationKind int

const (
	OptimizeKindUnoptimized     OptimizationKind = 0
	OptimizeKindOperatorMerging OptimizationKind = 1 << iota
	OptimizeKindBatching
	OptimizeKindAll OptimizationKind = OptimizeKindOperatorMerging & OptimizeKindBatching
)

// A global variable to control the batch size.
// It's ugly but I dont' have time for this.
var BATCH_SIZE int = 1024

func OptimizeOperatorMerging[T any](stream IStream[T]) IStream[T] {
	parent := stream.Parent()
	if parent == nil {
		return stream
	}

	for ; parent.Type() != StreamTypeSource && parent.Type() == stream.Type(); parent = stream.Parent() {
		// Merge `parent` and `stream` into one.
		name := fmt.Sprintf("<%s+%s>", stream.Name(), parent.Name())
		log.Info("Merging ", parent.Name(), " and ", stream.Name(), " into ", name)

		// Merge two worker_fn into one.
		old_worker_fn := stream.GetWorkerFn()
		paren_worker_fn := parent.GetWorkerFn()
		worker_fn := func(data T) (T, bool) {
			rtn1, more1 := paren_worker_fn(data)
			if !more1 {
				return rtn1, more1
			}

			rtn2, more2 := old_worker_fn(rtn1)
			return rtn2, more2
		}

		new_op := makeOperator(stream.GetWorkers(), parent.Parent(), worker_fn, nil, name, stream.Type())
		stream = &new_op
	}

	return stream
}

func RunDAG[T any](stream IStream[T], optimizations OptimizationKind) IStream[T] {
	if optimizations&OptimizeKindOperatorMerging != 0 {
		stream = OptimizeOperatorMerging(stream)
	}

	if optimizations&OptimizeKindBatching != 0 {
		for s := stream; s != nil; s = s.Parent() {
			go s.BatchExec(BATCH_SIZE)
		}
	} else {
		for s := stream; s != nil; s = s.Parent() {
			go s.Exec()
		}
	}

	return stream
}
