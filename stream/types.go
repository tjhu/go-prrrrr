package stream

type MapFn[T any, R any] func(T) R

type Void struct{}
