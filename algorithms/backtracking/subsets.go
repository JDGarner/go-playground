package backtracking

// Given an array nums of unique integers, return all possible subsets of nums.

// The solution set must not contain duplicate subsets.
// You may return the solution in any order.

// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Decision tree:
//                 root
//        1          2        3
//     2    3

//                    root
//          1                       no1
//     2         no2            2       no2
//  3    no3   3   no3             etc

//                                root
//                [1]                               []
//       [1,2]             [1]                 [2]       []
//  [1,2,3]    [1,2]  [1,3]   [1]        [2,3]       [2]       []

// bottom row with the search finished is what we want to add

func subsets(nums []int) [][]int {
	output := [][]int{}

	var dfs func(path []int, start int)

	dfs = func(path []int, start int) {
		// we've reached the end, insert what we've got
		if start > len(nums)-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			output = append(output, temp)

			return
		}

		// do first dfs where i include the nums[start]
		path = append(path, nums[start])
		dfs(path, start+1)
		path = path[:len(path)-1]

		// and second dfs where i don't include it
		dfs(path, start+1)
	}

	dfs([]int{}, 0)

	return output
}

func subsets2(nums []int) [][]int {
	output := [][]int{}
	subset := []int{}

	var dfs func(start int)

	dfs = func(start int) {
		// we've reached the end, insert what we've got
		if start > len(nums)-1 {
			temp := make([]int, len(subset))
			copy(temp, subset)
			output = append(output, temp)

			return
		}

		// do first dfs where i include the nums[start]
		subset = append(subset, nums[start])
		dfs(start + 1)

		// and second dfs where i don't include it
		subset = subset[:len(subset)-1] // backtrack step
		dfs(start + 1)
	}

	dfs(0)

	return output
}
