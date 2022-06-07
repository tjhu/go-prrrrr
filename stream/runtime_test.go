package stream

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSmallBatch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 10, 100}
	batch_sizes := []int{1, 2, 3, 4, 5, 6}
	var batch_size int

	run_test := func(t *testing.T) {
		logrus.Info("Testing batch size ", batch_size)

		// Original
		assert.ElementsMatch(t, slice, OfSlice(slice).ToSlice(OptimizeKindBatching))

		// Greater than 10
		greater_than_ten := func(x int) bool { return x > 10 }
		assert.ElementsMatch(t, []int{100}, OfSlice(slice).Filter(greater_than_ten).ToSlice(OptimizeKindBatching))

		// Less than 10
		less_than_ten := func(x int) bool { return x < 10 }
		assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice(slice).Filter(less_than_ten).ToSlice(OptimizeKindBatching))
	}

	for _, batch_size = range batch_sizes {
		t.Run(fmt.Sprint("batch_size=", batch_size), run_test)
	}
}

func TestBatchLargeSlice(t *testing.T) {
	slice := lo.Range(10000)
	batch_sizes := []int{1, 2, 3, 4, 5, 6, 100, 999}

	for batch_size := range batch_sizes {
		logrus.Info("Testing batch size ", batch_size)

		// Original
		assert.ElementsMatch(t, slice, OfSlice(slice).ToSlice(OptimizeKindBatching))

		// Greater than 10
		greater_than_ten := func(x int) bool { return x > 10 }
		assert.Equal(t, 99990, len(OfSlice(slice).Filter(greater_than_ten).ToSlice(OptimizeKindBatching)))

		// Less than 10
		less_than_ten := func(x int) bool { return x < 10 }
		assert.Equal(t, 10, len(OfSlice(slice).Filter(less_than_ten).ToSlice(OptimizeKindBatching)))
	}
}
