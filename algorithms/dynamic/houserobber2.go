package dynamic

// You are given an integer array nums where nums[i] represents the amount of
// money the ith house has.
// The houses are arranged in a circle, i.e. the first house and the last house
// are neighbors.

// You are planning to rob money from the houses, but you cannot rob two
// adjacent houses because the security system will automatically alert the
// police if two adjacent houses were both broken into.

// Return the maximum amount of money you can rob without alerting the police.

// Example 1:
// Input: nums = [3,4,3]
// Output: 4

// Explanation: You cannot rob nums[0] + nums[2] = 6 because nums[0]
// and nums[2] are adjacent houses. The maximum you can rob is nums[1] = 4.

// Example 2:
// Input: nums = [2,9,8,3,6]
// Output: 15

// Explanation: You cannot rob nums[0] + nums[2] + nums[4] = 16 because
// nums[0] and nums[4] are adjacent houses.
// The maximum you can rob is nums[1] + nums[4] = 15.

// [5,1,1,5]
// 0: max of [0] + sol[len(nums)-2] OR sol[len(nums)-1]
// 1: max of [1] + sol[len(nums)-1] OR sol[0]
// 2: max of [2] + sol[0] OR sol[1]

// 5, 1, 1, 1, 1, 5

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// 5, 1, 1, 6
	// 0, 1, 1, 6  => first pass set first house to 0
	// 5, 1, 1, 0. => second pass set last house to 0
	//             => then return max sol1[len(nums)-1], sol2[len(nums)-1]

	// 2, 8, 1, 7, 10
	// 0, 8, 1, 7, 10
	// 2, 8, 1, 7, 0

	// first calculate the solutions where we don't consider the first house
	// set it to 0
	solutions1 := map[int]int{
		0: 0,
		1: nums[1], // cant use nums[0]
	}

	solutions2 := map[int]int{
		0: nums[0], // now include the first house
		1: max(nums[0], nums[1]),
	}

	for i := 2; i < len(nums); i++ {
		solutions1[i] = max(solutions1[i-2]+nums[i], solutions1[i-1])

		if i == len(nums)-1 {
			solutions2[i] = max(solutions2[i-2], solutions2[i-1]) // don't include last house in calculation
		} else {
			solutions2[i] = max(solutions2[i-2]+nums[i], solutions2[i-1])
		}
	}

	return max(solutions1[len(nums)-1], solutions2[len(nums)-1])
}

func rob2WithArrays(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)

	// first calculate the solutions where we don't consider the first house
	// set it to 0
	solutions1 := make([]int, n)
	solutions1[0] = 0
	solutions1[1] = nums[1]

	solutions2 := make([]int, n)
	solutions2[0] = nums[0]
	solutions2[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		solutions1[i] = max(solutions1[i-2]+nums[i], solutions1[i-1])

		if i == len(nums)-1 {
			solutions2[i] = max(solutions2[i-2], solutions2[i-1]) // don't include last house in calculation
		} else {
			solutions2[i] = max(solutions2[i-2]+nums[i], solutions2[i-1])
		}
	}

	return max(solutions1[len(nums)-1], solutions2[len(nums)-1])
}
