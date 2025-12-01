package binaryheap

import "fmt"

func PushAndPopExample() {
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

func HeapifyExample() {
	data := []int{60, 50, 80, 40, 30, 10, 70, 20, 90}
	fmt.Println("data: ", data)

	heap := Heapify(data)

	fmt.Println("heapified:")
	fmt.Println(heap)

	data = []int{100, 101, 116, 107, 111, 115, 110, 106, 103, 116, 104}
	fmt.Println("data: ", data)

	heap = Heapify(data)

	fmt.Println("heapified:")
	fmt.Println(heap)
}
