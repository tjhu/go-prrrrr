package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	greater_than_ten := func(x int) bool { return x > 10 }
	less_than_ten := func(x int) bool { return x < 10 }
	even := func(x int) bool { return x%2 == 0 }
	odd := func(x int) bool { return x%2 != 0 }

	// Greater than 10
	t.Run("Greater", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{100}, OfSlice(slice).Filter(greater_than_ten).ToSlice())
	})

	// Less than 10
	t.Run("Less", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice(slice).Filter(less_than_ten).ToSlice())
	})

	// Even
	t.Run("Even", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{2, 4, 10, 100}, OfSlice(slice).Filter(even).ToSlice())
	})

	// Odd
	t.Run("Odd", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{1, 3}, OfSlice(slice).Filter(odd).ToSlice())
	})

	// Greater and less than 10
	t.Run("GreaterLess", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{}, OfSlice(slice).Filter(greater_than_ten).Filter(less_than_ten).ToSlice())
	})

	// Even and odd
	t.Run("EvenOdd", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{}, OfSlice(slice).Filter(even).Filter(odd).ToSlice())
	})

	// Even less than 10
	t.Run("EvenLess", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{2, 4}, OfSlice(slice).Filter(even).Filter(less_than_ten).ToSlice())
	})

	// Greater than 10 even
	t.Run("GreaterEven", func(t *testing.T) {
		slice := []int{12, 1, 2, 3, 4, 10, 100}
		assert.ElementsMatch(t, []int{12, 100}, OfSlice(slice).Filter(greater_than_ten).Filter(even).ToSlice())
	})
}
