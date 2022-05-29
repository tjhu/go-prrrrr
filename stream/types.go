package stream

type MapFn[T any, R any] func(T) R
type FilterFn[T any] func(T) bool
