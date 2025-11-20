package queue

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Add(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Pop() T {
	first := q.data[0]
	q.data = q.data[1:]

	return first
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) ForEach(f func(i int, t T)) {
	for i, v := range q.data {
		f(i, v)
	}
}
