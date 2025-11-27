package binaryheap

import "fmt"

func Example() {
	heap := New()
	heap.Push(7)
	heap.Push(42)
	heap.Push(14)
	heap.Push(27)
	heap.Push(2)
	heap.Push(1)
	heap.Push(31)
	heap.Push(2)

	fmt.Println(heap)

	pop, _ := heap.Pop()
	fmt.Println("popped off: ", pop)
	fmt.Println(heap)

	pop, _ = heap.Pop()
	fmt.Println("popped off: ", pop)
	fmt.Println(heap)

	pop, _ = heap.Pop()
	fmt.Println("popped off: ", pop)
	fmt.Println(heap)
}
