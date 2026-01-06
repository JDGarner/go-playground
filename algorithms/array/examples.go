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

func SpiralMatrixExample() {
	// matrix := [][]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// }

	// expected output: [1,2,3,4,8,12,11,10,9,5,6,7]

	// res := spiralOrder(matrix)
	// fmt.Println(">>> res: ", res)

	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
		{21, 22, 23, 24},
	}

	// expected output: [1,2,3,4,8,12,16,20,24,23,22,21,17,13,9,5,6,7,11,15,19,18,14,10]

	res2 := spiralOrder(matrix2)
	fmt.Println(">>> res2: ", res2)
}

func MajorityElementExample() {
	// res := majorityElement1([]int{1, 1, 2, 2, 3, 3, 1, 2})
	res := majorityElement([]int{1, 1, 2, 2, 3, 3, 1, 2})
	fmt.Println(">>> res: ", res)
}
