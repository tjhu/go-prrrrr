package stream

type FilterOperator[T any] struct {
	Operator[T, T]
}

func makeFilterOperator[T any]() FilterOperator[T] {
	panic("fiter")
}
