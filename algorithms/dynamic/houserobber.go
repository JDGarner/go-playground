package dynamic

// You are given an integer array nums where nums[i] represents the amount of
// money the ith house has.
// The houses are arranged in a straight line, i.e. the ith house is the
// neighbor of the (i-1)th and (i+1)th house.

// You are planning to rob money from the houses, but you cannot rob two
// adjacent houses because the security system will automatically alert
// the police if two adjacent houses were both broken into.

// Return the maximum amount of money you can rob without alerting the police.

// Example 1:
// Input: nums = [1,1,3,3]
// Output: 4

// Explanation: nums[0] + nums[2] = 1 + 3 = 4.

// Example 2:
// Input: nums = [2,9,8,3,6]
// Output: 16

// Explanation: nums[0] + nums[2] + nums[4] = 2 + 8 + 6 = 16.

// [5,1,1,5]
// 0: 5
// 1: max of [1] and sol[0]
// 2: max of sol[0] + [2] or sol[1]
// 3: max of sol[1] + [3] or sol[2]

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// map of best solution in the array so far up that that index
	// e.g. memo[1] returns the best solution for the subarray to index 1
	memo := map[int]int{
		0: nums[0],
		1: max(nums[0], nums[1]),
	}

	for i := 2; i < len(nums); i++ {
		memo[i] = max(memo[i-2]+nums[i], memo[i-1])
	}

	return memo[len(nums)-1]
}
