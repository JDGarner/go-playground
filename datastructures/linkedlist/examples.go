package linkedlist

import "fmt"

func Example() {
	l := New[int]()
	l.Add(7)
	l.Add(42)
	l.Add(18)
	l.Add(81)
	l.Add(2)
	fmt.Println("linked list:")
	fmt.Println(l)

	fmt.Println("reversed:")
	l.Reverse()
	fmt.Println(l)

	fmt.Println("remove 2 items:")
	l.Remove()
	l.Remove()
	fmt.Println(l)

	fmt.Println("remove 1 item:")
	l.Remove()
	fmt.Println(l)

	fmt.Println("reversed:")
	l.Reverse()
	fmt.Println(l)

	fmt.Println("remove 1 item:")
	l.Remove()
	fmt.Println(l)

	fmt.Println("reversed:")
	l.Reverse()
	fmt.Println(l)
}
