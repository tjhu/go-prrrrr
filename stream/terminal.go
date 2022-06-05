package stream

func RunDAG[T any](stream IStream[T]) {
	for {
		stream.Exec()
	}
}
