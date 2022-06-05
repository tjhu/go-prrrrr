package stream

func RunDAG[T any](stream IStream[T]) {
	for ; stream.Type() == IntermediateType; stream = stream.Parent() {
		go stream.Exec()
	}
	go stream.Exec()
}
