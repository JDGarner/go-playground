package backtracking

// You are given two integers n and k, return all possible combinations
// of k numbers chosen from the range [1, n].

// You may return the answer in any order.

// Example 1:
// Input: n = 3, k = 2
// Output: [
//     [1,2],
//     [1,3],
//     [2,3]
// ]

// n = [1, 2, 3]
// but return when len(path) = k

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(start int)

	backtrack = func(start int) {
		if len(path) == k {
			temp := append([]int{}, path...)
			res = append(res, temp)
			return
		}

		// just go forwards, start loop from a fixed start
		// explore one path where you add i and one path where you dont
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)

	return res
}
