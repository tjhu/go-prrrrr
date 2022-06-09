package benchmark

import (
	"fmt"
	"math"
	"testing"

	jucardi "github.com/jucardi/go-streams/streams"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/tjhu/go-prrrrr/stream"
)

const SIZE int = 10e6

func BenchmarkBatching(b *testing.B) {
	less_than_ten := func(x int) bool { return x < 10 }

	for _, num_threads := range []int{1, 4, 16, 32, 64} {
		b.Run(fmt.Sprint(num_threads), func(b *testing.B) {
			// jucardi
			b.Run("jucardi", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.Skip()
					b.StopTimer()
					slice := lo.Range(SIZE)
					b.StartTimer()
					less_than_ten := func(x interface{}) bool { return x.(int) < 10 }
					jucardi.FromArray(slice).Filter(less_than_ten, num_threads).Count()
				}
			})

			// Unoptimized
			b.Run("Unoptimized", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					slice := lo.Range(SIZE)
					b.StartTimer()
					stream.OfSlice(slice).Filter(less_than_ten).SetWorkers(num_threads).Count()
				}
			})

			// Batching
			for _, batch_size := range []int{32, 64, 512, 1024} {
				stream.BATCH_SIZE = batch_size
				b.Run(fmt.Sprint("Batching", batch_size), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						slice := lo.Range(SIZE)
						b.StartTimer()
						stream.OfSlice(slice).Filter(less_than_ten).SetWorkers(num_threads).Count(stream.OptimizeKindBatching)
					}
				})
			}
		})
	}
}

func BenchmarkMerging(b *testing.B) {
	add := func(x int) int { return x + 1 }

	for _, depth := range []int{1, 4, 16, 32, 64} {
		b.Run(fmt.Sprint(depth), func(b *testing.B) {
			b.Run("jucardi", func(b *testing.B) {
				b.Skip()
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					slice := lo.Range(SIZE)
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
				slice := lo.Range(SIZE)
				b.StartTimer()
				stream := stream.OfSlice(slice)
				for i := 0; i < depth; i++ {
					stream = stream.Map(add)
				}
				stream.Count()
			})

			b.Run("Optimized", func(b *testing.B) {
				b.StopTimer()
				slice := lo.Range(SIZE)
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

func BenchmarkAllOptimization(b *testing.B) {
	add := func(x int) int { return x + 1 }
	factorial := func(x int) int {
		result := 1
		for i := 1; i <= x; i++ {
			result *= i
		}
		return result
	}
	sin := func(x int) int {
		const DEPTH int = 10
		negative := 1
		result := 0

		for i := 0; i < DEPTH; i++ {
			j := i*2 + 1
			result += negative * int(math.Pow(float64(x), float64(j))) / factorial(j)
			negative *= -1
		}

		return result
	}
	workloads := []struct {
		name string
		fn   stream.MapFn[int]
	}{
		{"small", add},
		{"big", sin},
	}

	for _, workload := range workloads {
		b.Run(workload.name, func(b *testing.B) {
			for _, depth := range []int{1, 4, 16} {
				b.Run(fmt.Sprint(depth), func(b *testing.B) {
					b.Run("jucardi", func(b *testing.B) {
						b.Skip()
						for i := 0; i < b.N; i++ {
							b.StopTimer()
							slice := lo.Range(SIZE)
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
						slice := lo.Range(SIZE)
						b.StartTimer()
						stream := stream.OfSlice(slice)
						for i := 0; i < depth; i++ {
							stream = stream.Map(workload.fn)
						}
						stream.Count()
					})

					b.Run("Optimized", func(b *testing.B) {
						b.StopTimer()
						slice := lo.Range(SIZE)
						b.StartTimer()
						s := stream.OfSlice(slice)
						for i := 0; i < depth; i++ {
							s = s.Map(workload.fn)
						}
						s.Count(stream.OptimizeKindOperatorMerging)
					})
				})
			}
		})
	}
}

func init() {
	log.SetLevel(log.WarnLevel)
}
