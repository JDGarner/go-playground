package array

// You are given an integer array nums of size n,
// find all elements that appear more than ⌊ n/3 ⌋ times.
// You can return the result in any order.

// Example 1:
// Input: nums = [5,2,3,2,2,2,2,5,5,5]
// Output: [2,5]

// Example 2:
// Input: nums = [4,4,4,4,4]
// Output: [4]

// Example 3:
// Input: nums = [1,2,3]
// Output: []

// do candidate selection algo twice,
// don't include the previous winner on next run
// - include second winner if it's count is more than n/3

// there can only be max of 2 elements appearing more than n/3 times?

// [1, 1, 2, 2, 3, 3]

func majorityElement(nums []int) []int {
	n := len(nums)
	num1, num2 := -1, -1
	cnt1, cnt2 := 0, 0

	// Phase 1: Find candidates
	for _, num := range nums {
		if num == num1 {
			cnt1++ // Support current candidate 1
		} else if num == num2 {
			cnt2++ // Support current candidate 2
		} else if cnt1 == 0 {
			cnt1 = 1 // Slot 1 is empty, claim it
			num1 = num
		} else if cnt2 == 0 {
			cnt2 = 1 // Slot 2 is empty, claim it
			num2 = num
		} else {
			cnt1-- // Cancel votes from both candidates
			cnt2--
		}
	}

	// Phase 2: Verify candidates
	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			cnt1++
		} else if num == num2 {
			cnt2++
		}
	}

	res := []int{}
	if cnt1 > n/3 {
		res = append(res, num1)
	}
	if cnt2 > n/3 {
		res = append(res, num2)
	}

	return res
}

func majorityElement1(nums []int) int {
	res, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			res = num
		}

		if num == res {
			count++
		} else {
			count--
		}
	}

	return res
}
