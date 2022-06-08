package benchmark

import (
	"testing"

	jucardi "github.com/jucardi/go-streams/streams"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/tjhu/go-parallelstream/stream"
)

func BenchmarkRandInt(b *testing.B) {
	size := int(10e6)

	// jucardi
	b.Run("jucardi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			slice := lo.Range(size)
			b.StartTimer()
			less_than_ten := func(x interface{}) bool { return x.(int) < 10 }
			jucardi.FromArray(slice).Filter(less_than_ten).ToArray()
		}
	})

	// Unoptimized
	b.Run("Unoptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			slice := lo.Range(size)
			b.StartTimer()
			less_than_ten := func(x int) bool { return x < 10 }
			stream.OfSlice(slice).Filter(less_than_ten).ToSlice()
		}
	})

	// Batching
	b.Run("Batching", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			slice := lo.Range(size)
			b.StartTimer()
			less_than_ten := func(x int) bool { return x < 10 }
			stream.OfSlice(slice).Filter(less_than_ten).ToSlice(stream.OptimizeKindBatching)
		}
	})
}

func init() {
	log.SetLevel(log.WarnLevel)
}
