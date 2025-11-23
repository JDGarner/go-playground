package linkedlist

import (
	"fmt"
	"strings"
)

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func New[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Add(t T) {
	node := &Node[T]{
		value: t,
		next:  nil,
	}

	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node

	if l.head == nil {
		l.head = node
	}

	l.size++
}

func (l *SinglyLinkedList[T]) Remove() {
	if l.size == 0 {
		return
	}

	l.head = l.head.next
	l.size--

	if l.size == 0 {
		l.tail = nil
	}
}

func (l *SinglyLinkedList[T]) Reverse() {
	var prev *Node[T]
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head, l.tail = l.tail, l.head
}

func (n *Node[T]) Next() *Node[T] {
	if n == nil {
		return nil
	}

	return n.next
}

func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

func (l *SinglyLinkedList[T]) Traverse(f func(i int, t T)) {
	node := l.head
	i := 0

	for node != nil {
		f(i, node.value)
		node = node.next
		i++
	}
}

func (l *SinglyLinkedList[T]) String() string {
	if l.size == 0 {
		return ""
	}

	sb := strings.Builder{}

	l.Traverse(func(i int, t T) {
		if i == l.size-1 {
			sb.WriteString(fmt.Sprintf("%v", t))
		} else {
			sb.WriteString(fmt.Sprintf("%v -> ", t))
		}
	})

	sb.WriteString(fmt.Sprintf(" (Head is %v, Tail is %v)", l.head.value, l.tail.value))

	return sb.String()
}
