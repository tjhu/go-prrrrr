package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 10, 100}

	// Greater than 10
	greater_than_ten := func(x int) bool { return x > 10 }
	assert.ElementsMatch(t, OfSlice(slice).Filter(greater_than_ten).ToSlice(), []int{100})

	// Less than 10
	less_than_ten := func(x int) bool { return x < 10 }
	assert.ElementsMatch(t, OfSlice(slice).Filter(less_than_ten).ToSlice(), []int{1, 2, 3, 4})

	// Even
	even := func(x int) bool { return x%2 == 0 }
	assert.ElementsMatch(t, OfSlice(slice).Filter(even).ToSlice(), []int{2, 4, 10, 100})

	// Odd
	odd := func(x int) bool { return x%2 != 0 }
	assert.ElementsMatch(t, OfSlice(slice).Filter(odd).ToSlice(), []int{1, 3})
}
