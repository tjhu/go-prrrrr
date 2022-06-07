package stream

import (
	"golang.org/x/exp/constraints"
)

func ChanToSlice[T any](in <-chan T) []T {
	slice := make([]T, 0)
	for element := range in {
		slice = append(slice, element)
	}
	return slice
}

func BatchChanToSlice[T any](in <-chan []T) []T {
	slice := make([]T, 0)
	for elements := range in {
		slice = append(slice, elements...)
	}
	return slice
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	} else {
		return a
	}
}
