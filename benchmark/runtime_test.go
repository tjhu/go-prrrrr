package benchmark

import (
	"fmt"
	"testing"

	jucardi "github.com/jucardi/go-streams/streams"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/tjhu/go-parallelstream/stream"
)

func BenchmarkBatching(b *testing.B) {
	size := int(10e6)

	for _, num_threads := range []int{1, 2, 4, 8} {
		b.Run(fmt.Sprint(num_threads), func(b *testing.B) {
			// jucardi
			b.Run("jucardi", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					slice := lo.Range(size)
					b.StartTimer()
					less_than_ten := func(x interface{}) bool { return x.(int) < 10 }
					jucardi.FromArray(slice).Filter(less_than_ten, num_threads).ToArray()
				}
			})

			// Unoptimized
			b.Run("Unoptimized", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					slice := lo.Range(size)
					b.StartTimer()
					less_than_ten := func(x int) bool { return x < 10 }
					stream.OfSlice(slice).Filter(less_than_ten).SetWorkers(num_threads).ToSlice()
				}
			})

			// Batching
			for _, batch_size := range []int{128, 512, 1024, 2048, 8192} {
				stream.BATCH_SIZE = batch_size
				b.Run(fmt.Sprint("Batching", batch_size), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						slice := lo.Range(size)
						b.StartTimer()
						less_than_ten := func(x int) bool { return x < 10 }
						stream.OfSlice(slice).Filter(less_than_ten).SetWorkers(num_threads).ToSlice(stream.OptimizeKindBatching)
					}
				})
			}
		})
	}
}

func BenchmarkMerging(b *testing.B) {
	size := int(10e6)
	add := func(x int) int { return x + 1 }

	for _, depth := range []int{1, 4, 16} {
		b.Run(fmt.Sprint("depth=", depth), func(b *testing.B) {
			b.Run("jucardi", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					slice := lo.Range(size)
					b.StartTimer()
					add := func(x interface{}) interface{} { return x.(int) + 1 }
					stream := jucardi.FromArray(slice)
					for i := 0; i < depth; i++ {
						stream = stream.Map(add)
					}
					stream.Count()
				}
			})

			b.Run("Unoptimized", func(b *testing.B) {
				b.StopTimer()
				slice := lo.Range(size)
				b.StartTimer()
				stream := stream.OfSlice(slice)
				for i := 0; i < depth; i++ {
					stream = stream.Map(add)
				}
				stream.Count()
			})

			b.Run("Optimized", func(b *testing.B) {
				b.StopTimer()
				slice := lo.Range(size)
				b.StartTimer()
				s := stream.OfSlice(slice)
				for i := 0; i < depth; i++ {
					s = s.Map(add)
				}
				s.Count(stream.OptimizeKindOperatorMerging)
			})
		})
	}
}

func init() {
	log.SetLevel(log.WarnLevel)
}
