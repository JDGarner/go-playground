package queue

type LinkedListQueue[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func NewLinkedListQueue[T comparable]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{}
}

// Adds onto the end (tail) of the queue
func (q *LinkedListQueue[T]) Enqueue(value T) {
	node := &Node[T]{
		value: value,
	}

	if q.tail != nil {
		q.tail.next = node
	}
	q.tail = node

	if q.head == nil {
		q.head = node
	}

	q.size++
}

// Remove from front (head) of the queue
func (q *LinkedListQueue[T]) Dequeue() T {
	if q.size == 0 {
		var zero T
		return zero
	}

	node := q.head
	q.head = q.head.next

	q.size--

	if q.size == 0 {
		q.tail = nil
	}

	return node.value
}

func (q *LinkedListQueue[T]) Len() int {
	return q.size
}
