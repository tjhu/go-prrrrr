package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	// Test with empty
	assert.ElementsMatch(t, []int{}, Of[int]().ToSlice())
	assert.ElementsMatch(t, []int{}, OfSlice([]int{}).ToSlice())

	// Test with non-empty
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, Of(1, 2, 3, 4).ToSlice())
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice([]int{1, 2, 3, 4}).ToSlice())
}
