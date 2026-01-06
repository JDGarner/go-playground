package slidingwindow

import "math"

// You are given a sorted integer array arr, two integers k and x,
// return the k closest integers to x in the array.
// The result should also be sorted in ascending order.

// An integer a is closer to x than an integer b if:
// |a - x| < |b - x|, or
// |a - x| == |b - x| and a < b

// (if the distance is the same, choose the smaller one)

// Example 1:
// Input: arr = [2,3,4,5,8], k = 2, x = 6
// Output: [4,5]

// Example 2:
// Input: arr = [2,3,4], k = 3, x = 1
// Output: [2,3,4]

// keep a left starting at 0
// keep a right starting at end
// loop through until the gap between left and right is k
// - if the gap from arr[left] to x is less than arr[right] to x
//  - decrement right (since it's further away we don't want it in the final window)
// - else:
//   - increment left

// return arr[left:right+1]

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1

	for right-left >= k {
		if math.Abs(float64(arr[left]-x)) <= math.Abs(float64(arr[right]-x)) {
			right--
		} else {
			left++
		}
	}

	return arr[left:right+1]
}

