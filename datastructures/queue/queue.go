package queue

import (
	"fmt"
	"strings"
)

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// Puts the value at the end of slice (back of the queue)
func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

// Takes the value from front of slice (front of the queue)
func (q *Queue[T]) Dequeue() T {
	first := q.data[0]
	q.data = q.data[1:]

	return first
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) String() string {
	var sb strings.Builder

	for _, v := range q.data {
		sb.WriteString(fmt.Sprintf("%v, ", v))
	}

	return sb.String()
}
