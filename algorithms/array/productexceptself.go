package array

// Given an integer array nums, return an array output where output[i]
// is the product of all the elements of nums except nums[i].

// Each product is guaranteed to fit in a 32-bit integer.

// Follow-up: Could you solve it in
// O(n) time without using the division operation?

// Example 1:

// Input: nums = [1,2,4,6]

// Output: [48,24,12,8]
// Example 2:

// Input: nums = [-1,0,1,2,3]

// Output: [0,-6,0,0,0]

// Input: nums = [1,2,4,6]
// [1-2]

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	productAll := 1
	productWithoutZeros := 0
	zeroCount := 0

	for i := 0; i < len(nums); i++ {
		productAll *= nums[i]
		if nums[i] != 0 {
			if productWithoutZeros == 0 {
				productWithoutZeros = 1
			}
			productWithoutZeros *= nums[i]
		} else {
			zeroCount++
		}
	}

	if zeroCount > 1 {
		return output
	}

	for i := range output {
		if nums[i] == 0 {
			output[i] = productWithoutZeros
		} else {
			output[i] = productAll / nums[i]
		}
	}

	return output
}

// -------------------------------------------------------------------------
// Ideal Solution
// -------------------------------------------------------------------------
// For each index, we need the product of all elements before it and all elements after it.
// Instead of recomputing the product repeatedly, we can pre-compute two helpful arrays:

// Prefix product: pref[i] = product of all elements to the left of i
// Suffix product: suff[i] = product of all elements to the right of i

func productExceptSelfIdeal(nums []int) []int {
	n := len(nums)
	pref := make([]int, n)
	suff := make([]int, n)
	output := make([]int, n)

	pref[0], suff[n-1] = 1, 1

	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		suff[i] = suff[i+1] * nums[i+1]
	}

	for i := range n {
		output[i] = pref[i] * suff[i]
	}

	return output
}
