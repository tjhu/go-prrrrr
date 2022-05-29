package stream

type IStream[T any] interface {
	Filter(FilterFn[T]) IStream[T]
}

type Stream[T any] struct {
}

// func (stream Stream[T]) Map[T any, R any](fn MapFn[T, R]) Stream[R] {}

func Map[T any, R any](stream Stream[T], fn MapFn[T, R]) Stream[R] {
	panic("asd")
}
