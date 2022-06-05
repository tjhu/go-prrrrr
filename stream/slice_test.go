package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.ElementsMatch(t, Of(1, 2, 3, 4).ToSlice(), []int{1, 2, 3, 4})
}
