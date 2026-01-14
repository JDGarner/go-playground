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

	// res := change(5, []int{1, 2, 5})
	// fmt.Println(">>> res: ", res)

	// res := rob([]int{5, 1, 1, 6})
	// fmt.Println(">>> res: ", res)

	// res := rob2([]int{5, 1, 2, 6, 12, 7, 9, 3, 4, 10})
	// fmt.Println(">>> res: ", res)

	// res := longestCommonSubsequence("hellox", "helloz")
	// res := longestCommonSubsequence("xpkhello", "zlmhello")
	// res := longestCommonSubsequence("xpkhello", "zlmhgfhelloqqqqqqqq")
	// fmt.Println(">>> res: ", res)

	res := lengthOfLIS([]int{9, 1, 4, 2, 3, 3, 7, 1, 10, 0})
	fmt.Println(">>> res: ", res)
}
