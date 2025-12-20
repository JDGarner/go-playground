package array

import "fmt"

func ProductExceptSelfExample() {
	// res := productExceptSelf([]int{-1, 0, 1, 2, 3})
	res := productExceptSelfIdeal([]int{1, 2, 4, 8, 16})
	fmt.Println(">>> res: ", res)
}

func RotatedSortedExample() {
	// res := findMin([]int{1, 2})
	// fmt.Println(">>> res: ", res)
	res := findMin([]int{4, 5, 6, 7})
	fmt.Println(">>> res: ", res)
	// res := findMin([]int{3, 4, 5, 6, 1, 2})
	// fmt.Println(">>> res: ", res)
	// res = findMin([]int{6, 1, 2, 3, 4, 5})
	// fmt.Println(">>> res: ", res)
	// res = findMin([]int{1, 2, 3, 4, 5, 6})
	// fmt.Println(">>> res: ", res)
	// res = findMin([]int{2, 3, 4, 5, 6, 1})
	// fmt.Println(">>> res: ", res)
}
