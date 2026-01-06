package backtracking

// Given an array nums of unique integers, return all the possible permutations.
// You may return the answer in any order.

// Example 1:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// Example 2:
// Input: nums = [7]
// Output: [[7]]

// decision tree:
//        1               2               3
//     2     3        1       3       1      2
//     3     2        3       2       2      1

// from each starting index, to explore all possiblities you have to
// jump to any other index and then include from there
// - forwards and backwards

func permute(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	chosen := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := append([]int{}, path...)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !chosen[i] {
				path = append(path, nums[i])
				chosen[i] = true
				backtrack()
				path = path[:len(path)-1]
				chosen[i] = false
			}
		}
	}

	backtrack()

	return res
}
