package backtracking

// You are given an array of distinct integers nums and a target integer target.
// Your task is to return a list of all unique combinations of nums where the
// chosen numbers sum to target.

// The same number may be chosen from nums an unlimited number of times.
// Two combinations are the same if the frequency of each of the chosen numbers
// is the same, otherwise they are different.

// You may return the combinations in any order and the order of the numbers
// in each combination can be in any order.

// Example 1:
// Input:
// nums = [2,5,6,9]
// target = 9
// Output: [[2,2,5],[9]]

// Example 1:
// Input:
// nums = [3,4,5]
// target = 16
// Output: [[3,3,3,3,4],[3,3,5,5],[4,4,4,4],[3,4,4,5]]

// Example 3:
// Input:
// nums = [3]
// target = 5
// Output: []

func combinationSum(nums []int, target int) [][]int {
	output := [][]int{}

	var dfs func(path []int, startIndex, currentSum int)

	dfs = func(path []int, startIndex, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			output = append(output, temp)
			return
		}

		if currentSum > target {
			return
		}

		// explore all 'decision paths' from this point
		// start from startIndex so we don't go backwards to explore earlier indexes again
		// (otherwise we could get duplicates like [2, 2, 5] and [2, 5, 2])
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(path, i, currentSum+nums[i])
			path = path[:len(path)-1]
		}
	}

	dfs([]int{}, 0, 0)

	return output
}

// nums = [2,5,6,9]
// target = 9
//               2
//    2      5      6     9
//  2569   2569   2569   2569
