package binarysearch

import "fmt"

// You are given a sorted array consisting of only integers where
// every element appears exactly twice, except for one element which
// appears exactly once.

// Return the single element that appears only once.

// Your solution must run in O(log n) time and O(1) space.

// Example 1:
// Input: nums = [1,1,2,3,3,4,4,8,8]
// Output: 2

// Example 2:
// Input: nums = [3,3,7,7,10,11,11]
// Output: 10

// [2, 2, 4, 4, 6, 6]
//  0,    2,    4,

// pairs always start on an even index
// until a single element is inserted:

// [2, 2, 4, 4, 5, 6, 6]
//  0,    2,       5,

// [1, 2, 2, 6]

// if a pair is starting on an even index it must be BEFORE the single element

func singleNonDuplicate(nums []int) int {
	fmt.Println(">>> nums: ", nums)
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		fmt.Println(">>> left, mid, right: ", left, mid, right)

		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		// if mid is on an even index, then it should be the start of the pair, if
		// the singleElement is later on
		isMidEven := mid%2 == 0
		compareTo := nums[mid-1]
		if isMidEven {
			compareTo = nums[mid+1]
		}

		if nums[mid] == compareTo {
			fmt.Println(">>> num on the right side")
			left = mid + 1
		} else {
			fmt.Println(">>> num on the left side")
			right = mid - 1
		}
	}

	fmt.Println(">>> done! left, right: ", left, right)

	return nums[right]
}
