package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	// Test with empty
	assert.ElementsMatch(t, Of[int]().ToSlice(), []int{})
	assert.ElementsMatch(t, OfSlice([]int{}).ToSlice(), []int{})

	// Test with non-empty
	assert.ElementsMatch(t, Of(1, 2, 3, 4).ToSlice(), []int{1, 2, 3, 4})
	assert.ElementsMatch(t, OfSlice([]int{1, 2, 3, 4}).ToSlice(), []int{1, 2, 3, 4})
}
