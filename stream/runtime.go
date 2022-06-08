// This file defines the runtime for running the pipeline.

package stream

// OptimizationKind
type OptimizationKind int

const (
	OptimizeKindUnoptimized     OptimizationKind = 0
	OptimizeKindOperatorMerging OptimizationKind = 1 << iota
	OptimizeKindBatching
)

// A global variable to control the batch size.
// It's ugly but I dont' have time for this.
var BATCH_SIZE int = 1024

func OptimizeOperatorMerging[T any](stream IStream[T]) IStream[T] {
	return stream
}

func RunDAG[T any](stream IStream[T], optimizations OptimizationKind) {
	if optimizations&OptimizeKindOperatorMerging != 0 {
		stream = OptimizeOperatorMerging(stream)
	}

	if optimizations&OptimizeKindBatching != 0 {
		for ; stream != nil; stream = stream.Parent() {
			go stream.BatchExec(BATCH_SIZE)
		}
	} else {
		for ; stream != nil; stream = stream.Parent() {
			go stream.Exec()
		}
	}

}
