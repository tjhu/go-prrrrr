package stream

func ChanToSlice[T any](in <-chan T) []T {
	slice := make([]T, 0)
	for element := range in {
		slice = append(slice, element)
	}
	return slice
}
