package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 10, 100}

	// Greater than 10
	greater_than_ten := func(x int) bool { return x > 10 }
	assert.ElementsMatch(t, []int{100}, OfSlice(slice).Filter(greater_than_ten).ToSlice())

	// Less than 10
	less_than_ten := func(x int) bool { return x < 10 }
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice(slice).Filter(less_than_ten).ToSlice())

	// Even
	even := func(x int) bool { return x%2 == 0 }
	assert.ElementsMatch(t, []int{2, 4, 10, 100}, OfSlice(slice).Filter(even).ToSlice())

	// Odd
	odd := func(x int) bool { return x%2 != 0 }
	assert.ElementsMatch(t, []int{1, 3}, OfSlice(slice).Filter(odd).ToSlice())
}
