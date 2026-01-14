package dynamic

// Given an integer array nums, return the length of the longest strictly increasing
// subsequence.

// A subsequence is a sequence that can be derived from the given sequence by
// deleting some or no elements without changing the relative order of the
// remaining characters.

// Example 1:
// Input: nums = [9,1,4,2,3,3,7]
// Output: 4

// Explanation: The longest increasing subsequence is [1,2,3,7], which has a length of 4.

func lengthOfLIS(nums []int) int {
	n := len(nums)
	memo := make([][]int, n)
	for i := range n {
		memo[i] = make([]int, n)
		for j := range n {
			memo[i][j] = -1
		}
	}

	var dfs func(currentIndex, prevIndex int) int

	dfs = func(currentIndex, prevIndex int) int {
		if currentIndex == n {
			return 0
		}
		if prevIndex != -1 && memo[currentIndex][prevIndex] != -1 {
			return memo[currentIndex][prevIndex]
		}

		// don't include it:
		LIS := dfs(currentIndex+1, prevIndex)

		// include it:
		if prevIndex == -1 || nums[currentIndex] > nums[prevIndex] {
			LIS = max(LIS, 1+dfs(currentIndex+1, currentIndex))
		}

		if prevIndex != -1 {
			memo[currentIndex][prevIndex] = LIS
		}

		return LIS
	}

	return dfs(0, -1)
}

func lengthOfLISBottomUp(nums []int) int {
	n := len(nums)
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	// loop backwards from n-1 to 0
	for i := n - 1; i >= 0; i-- {

		// loop backwards from i-1 to -1
		for j := i - 1; j >= -1; j-- {

			// 
			LIS := memo[i+1][j+1] // Not including nums[i]

			if j == -1 || nums[j] < nums[i] {
				LIS = max(LIS, 1+memo[i+1][i+1]) // Including nums[i]
			}

			memo[i][j+1] = LIS
		}
	}

	return memo[0][0]
}
