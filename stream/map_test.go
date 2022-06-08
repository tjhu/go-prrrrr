package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	multiply := func(x int) int { return x * 2 }
	add := func(x int) int { return x + 1 }

	t.Run("Multiply", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{2, 4, 6, 8, 20, 200}, OfSlice(slice).Map(multiply).ToSlice())
	})

	t.Run("Add", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{2, 3, 4, 5, 11, 101}, OfSlice(slice).Map(add).ToSlice())
	})

	t.Run("AddMultiply", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{4, 6, 8, 10, 22, 202}, OfSlice(slice).Map(add).Map(multiply).ToSlice())
	})

	t.Run("MultiplyAdd", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{3, 5, 7, 9, 21, 201}, OfSlice(slice).Map(multiply).Map(add).ToSlice())
	})
}
