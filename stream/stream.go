package stream

type IStream[T any] interface {
	// TODO: turn them into lower case. Private methods.
	// Returns the output channel.
	Out() <-chan T
	// Returns a batching output channel.
	BatchedOut() <-chan []T
	// Execute the operator and wait for it to finish.
	Exec()
	// Execute the batch version of the operator.
	BatchExec(batch_size int)
	// Returns the parent stream.
	Parent() IStream[T]
	// Return the underlaying worker_fn for optimization.
	GetWorkerFn() OptionalMapFn[T]

	// Return the type of the stream
	Type() StreamType
	// Get the depth of the stream.
	Depth() int
	// Set the number of max workers.
	SetWorkers(num_workers int) IStream[T]
	// Get the number of max workers.
	GetWorkers() int
	// Return the name of the stream.
	Name() string
	// Return a mapped stream.
	Map(MapFn[T]) IStream[T]
	// Returns a filtered stream.
	Filter(FilterFn[T]) IStream[T]
	// Materialize the stream to a slice.
	ToSlice(optimizations ...OptimizationKind) []T
	// Count the length of the stream.
	Count(optimizations ...OptimizationKind) int
}
