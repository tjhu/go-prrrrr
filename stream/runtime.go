// This file defines the runtime for running the pipeline.

package stream

// OptimizationKind
type OptimizationKind int

const (
	OptimizeKindUnoptimized     OptimizationKind = 0
	OptimizeKindOperatorMerging OptimizationKind = 1 << iota
	OptimizeKindBatching
)

func OptimizeBatching[T any](stream IStream[T]) IStream[T] {
	return stream
}

func OptimizeOperatorMerging[T any](stream IStream[T]) IStream[T] {
	return stream
}

func RunDAG[T any](stream IStream[T], optimizations OptimizationKind) {
	if optimizations&OptimizeKindOperatorMerging != 0 {
		stream = OptimizeOperatorMerging(stream)
	}

	if optimizations&OptimizeKindBatching != 0 {
		stream = OptimizeBatching(stream)
	}

	for ; stream.Type() == IntermediateType; stream = stream.Parent() {
		go stream.Exec()
	}
	go stream.Exec()
}
