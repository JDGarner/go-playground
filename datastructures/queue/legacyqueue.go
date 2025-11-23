package queue

type LegacyQueue[T any] struct {
	data []T
}

func NewLegacy[T any]() *LegacyQueue[T] {
	return &LegacyQueue[T]{
		data: make([]T, 0),
	}
}

func (q *LegacyQueue[T]) Add(value T) {
	q.data = append(q.data, value)
}

func (q *LegacyQueue[T]) Pop() T {
	first := q.data[0]
	q.data = q.data[1:]

	return first
}

func (q *LegacyQueue[T]) Len() int {
	return len(q.data)
}
