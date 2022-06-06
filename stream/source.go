package stream

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Generator[T any] func() (T, bool)

// Data source stream
type Source[T any] struct {
	Operator[T]
	generator Generator[T]
}

func makeSource[T any](generator Generator[T], name string) Source[T] {
	return Source[T]{
		makeOperator[T](1, nil, nil, fmt.Sprintf("Source-%s", name), SourceType),
		generator,
	}
}

func (src *Source[T]) Exec() {
	log.Info("Generating source stage: ", src.name)
	var wg sync.WaitGroup

	for i := 0; i < src.num_workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Dump the generator into the channel.
			for element, more := src.generator(); more; element, more = src.generator() {
				src.out <- element
			}
		}()
	}

	wg.Wait()
	close(src.out)
	log.Info("Finished generating source stage: ", src.name)
}
