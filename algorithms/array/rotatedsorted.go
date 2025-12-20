package array

// You are given an array of length n which was originally sorted in ascending order.
// It has now been rotated between 1 and n times.
// For example, the array nums = [1,2,3,4,5,6] might become:

// [3,4,5,6,1,2] if it was rotated 4 times.
// [1,2,3,4,5,6] if it was rotated 6 times.
// Notice that rotating the array 4 times moves the last four elements of the
// array to the beginning. Rotating the array 6 times produces the original array.

// Assuming all elements in the rotated sorted array nums are unique,
// return the minimum element of this array.

// A solution that runs in O(n) time is trivial, can you write an algorithm
// that runs in O(log n) time?

// Example 1:
// Input: nums = [3,4,5,6,1,2]
// Output: 1

// Example 2:
// Input: nums = [4,5,0,1,2,3]
// Output: 0

// Example 3:
// Input: nums = [4,5,6,7]
// Output: 4

// - repeatedly find midpoint:
// - while midpoint is not 0 or end of nums
// - if value above is less than mid
//   - return that value
// [3,4,5,6,1,2]
// [6,1,2,3,4,5]

// take a half
// if the start is less than the end
// take the other half (beacuse split is in the other half)

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	mid := len(nums) / 2

	// is split before or after midpoint
	for start < end && mid < len(nums)-1 {
		// is the split just after midpoint?
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// is split just before the midpoint
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[end] > nums[mid] {
			// split must be in first half (start to mid)
			end, mid = mid, (end/2)+start
		} else {
			// split must be in second half (mid to end)
			start, mid = mid, (end/2)+start
		}
	}

	if mid == len(nums)-1 {
		return min(nums[mid], nums[0])
	}

	if mid == 0 {
		return min(nums[0], nums[len(nums)-1])
	}

	return min(min(nums[mid], nums[mid+1]), nums[mid-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
