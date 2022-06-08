package stream

type IStream[T any] interface {
	// Returns the output channel.
	Out() <-chan T
	// Returns a batching output channel.
	BatchedOut() <-chan []T
	// Execute the operator and wait for it to finish.
	Exec()
	// Execute the batch version of the operator.
	BatchExec(batch_size int)
	// Set the number of max workers.
	Workers(num_workers int) IStream[T]
	// Returns the parent stream.
	Parent() IStream[T]
	// Return the type of the stream
	Type() StreamType

	// Return a mapped stream.
	Map(MapFn[T]) IStream[T]
	// Returns a filtered stream.
	Filter(FilterFn[T]) IStream[T]
	// Materialize the stream to a slice.
	ToSlice(optimizations ...OptimizationKind) []T
}
