package dynamic

import "fmt"

func FibExample() {
	fmt.Println(Fib(7))
}

func CountPathsExample() {
	fmt.Println(CountPaths(4, 4))
}

func DynamicExample() {
	// res := coinChange([]int{1, 4, 5}, 13)
	// fmt.Println(">>> res: ", res)

	// res := coinChange([]int{4, 5}, 11)
	// fmt.Println(">>> res: ", res)

	res := change(5, []int{1, 2, 5})
	fmt.Println(">>> res: ", res)
}
