package stream

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	batch_sizes := []int{1, 2, 3, 4, 5, 6, 100, 999}

	run_test := func(t *testing.T) {
		for _, num_workers := range []int{2, 3, 5, 6} {
			// Greater than 10
			t.Run("Greater10", func(t *testing.T) {
				slice := lo.Range(10000)
				greater_than_ten := func(x int) bool { return x > 10 }
				assert.Equal(t, 9989, len(OfSlice(slice).Filter(greater_than_ten).Workers(num_workers).ToSlice(OptimizeKindBatching)))
			})

			// Less than 10
			t.Run("Less10", func(t *testing.T) {
				slice := lo.Range(10000)
				less_than_ten := func(x int) bool { return x < 10 }
				assert.Equal(t, 10, len(OfSlice(slice).Filter(less_than_ten).Workers(num_workers).ToSlice(OptimizeKindBatching)))
			})

			// Small greater than 10
			t.Run("SmallGreater10", func(t *testing.T) {
				slice := []int{1, 2, 3, 4, 10, 100}
				greater_than_ten := func(x int) bool { return x > 10 }
				assert.ElementsMatch(t, []int{100}, OfSlice(slice).Filter(greater_than_ten).Workers(num_workers).ToSlice(OptimizeKindBatching))
			})

			// Small less than 10
			t.Run("SmallLess10", func(t *testing.T) {
				slice := []int{1, 2, 3, 4, 10, 100}
				less_than_ten := func(x int) bool { return x < 10 }
				assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice(slice).Filter(less_than_ten).Workers(num_workers).ToSlice(OptimizeKindBatching))
			})
		}
	}

	for _, BATCH_SIZE = range batch_sizes {
		t.Run(fmt.Sprint("batch_size=", BATCH_SIZE), run_test)
	}
}

func TestSmallBatch(t *testing.T) {
	batch_sizes := []int{1, 2, 3, 4, 5, 6}

	run_test := func(t *testing.T) {
		// Unmodified
		t.Run("unmodified", func(t *testing.T) {
			slice := []int{1, 2, 3, 4, 10, 100}
			assert.ElementsMatch(t, slice, OfSlice(slice).ToSlice(OptimizeKindBatching))
		})

		// Greater than 10
		t.Run("Greater10", func(t *testing.T) {
			slice := []int{1, 2, 3, 4, 10, 100}
			greater_than_ten := func(x int) bool { return x > 10 }
			assert.ElementsMatch(t, []int{100}, OfSlice(slice).Filter(greater_than_ten).ToSlice(OptimizeKindBatching))
		})

		// Less than 10
		t.Run("Less10", func(t *testing.T) {
			slice := []int{1, 2, 3, 4, 10, 100}
			less_than_ten := func(x int) bool { return x < 10 }
			assert.ElementsMatch(t, []int{1, 2, 3, 4}, OfSlice(slice).Filter(less_than_ten).ToSlice(OptimizeKindBatching))
		})

	}

	for _, BATCH_SIZE = range batch_sizes {
		t.Run(fmt.Sprint("batch_size=", BATCH_SIZE), run_test)
	}
}

func TestBatchLargeSlice(t *testing.T) {
	batch_sizes := []int{1, 2, 3, 4, 5, 6, 100, 999}

	run_test := func(t *testing.T) {
		// Unmodified
		t.Run("unmodified", func(t *testing.T) {
			slice := lo.Range(10000)
			assert.ElementsMatch(t, slice, OfSlice(slice).ToSlice(OptimizeKindBatching))
		})

		// Greater than 10
		t.Run("Greater10", func(t *testing.T) {
			slice := lo.Range(10000)
			greater_than_ten := func(x int) bool { return x > 10 }
			assert.Equal(t, 9989, len(OfSlice(slice).Filter(greater_than_ten).ToSlice(OptimizeKindBatching)))
		})

		// Less than 10
		t.Run("Less10", func(t *testing.T) {
			slice := lo.Range(10000)
			less_than_ten := func(x int) bool { return x < 10 }
			assert.Equal(t, 10, len(OfSlice(slice).Filter(less_than_ten).ToSlice(OptimizeKindBatching)))
		})
	}

	for _, BATCH_SIZE = range batch_sizes {
		t.Run(fmt.Sprint("batch_size=", BATCH_SIZE), run_test)
	}
}
