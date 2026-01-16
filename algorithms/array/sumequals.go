package array

// You are given an array of integers nums and an integer k, return the total
// number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,-1,1,2], k = 2
// Output: 4

// Explanation: [2], [2,-1,1], [-1,1,2], [2] are the subarrays whose sum is equals to k.

// Example 2:
// Input: nums = [4,4,4,4,4,4], k = 4
// Output: 6

// can decide to include or not include at each step
// if it passes k i still need to keep it because there could be negative numbers

func subarraySum(nums []int, k int) int {
	res, curSum := 0, 0
	prefixSums := map[int]int{0: 1}

	// prefixSums map = "A counter that remembers: for each possible total,
	// how many times have I seen that total as I walked through the array?"

	for _, num := range nums {
		curSum += num
		diff := curSum - k
		res += prefixSums[diff]

		// prefixSums stores how many times each cumulative sum has been seen
		prefixSums[curSum]++
	}

	return res
}

// brute force:
// - find all subarrays (double for loop), keeping track of total as we go
// - if the total == k => increment counter

func subarraySumBruteForce(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		sum := 0

		for j := i; j < len(nums); j++ {
			sum += nums[k]
			if sum == k {
				res++
			}
		}
	}

	return res
}
