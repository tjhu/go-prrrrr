package stream

type IStream[T any] interface {
	// Returns the output channel.
	Out() <-chan T
	// Execute the operator and wait for it to finish.
	Exec()
	// Set the number of max workers.
	Workers(num_workers int)
	// Returns the parent stream.
	Parent() IStream[T]
	// Return the type of the stream
	Type() StreamType

	// Returns a filtered stream.
	Filter(FilterFn[T]) IStream[T]
	// Materialize the stream to a slice.
	ToSlice(optimizations ...OptimizationKind) []T
}

// func (stream Stream[T]) Map[T any, R any](fn MapFn[T, R]) Stream[R] {}

func Map[T any, R any](stream IStream[T], fn MapFn[T, R]) IStream[R] {
	panic("asd")
}
