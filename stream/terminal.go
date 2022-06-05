package stream

func RunDAG[T any](stream IStream[T]) {
	go stream.Exec()
}
