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

func ChanCount[T any](in <-chan T) int {
	ctr := 0
	for _, e := <-in; e; _, e = <-in {
		ctr++
	}
	return ctr
}

func BatchChanCount[T any](in <-chan []T) int {
	ctr := 0
	for elements := range in {
		ctr += len(elements)
	}
	return ctr
}

func Min[T constraints.Ordered](a, b T) T {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	} else {
		return a
	}
}
