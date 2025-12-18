package twopointers

import "math"

// You are given an integer array heights where heights[i] represents the height of the
// ith bar.

// You may choose any two bars to form a container.
// Return the maximum amount of water a container can store.

// Example 1:
// Input: height = [1,7,2,5,4,7,3,6]
// Output: 36

// Example 2:
// Input: height = [2,2,2]
// Output: 4

func MaxArea(heights []int) int {
	max := 0
	l := 0
	r := len(heights) - 1

	for l < r {
		var newMax int
		if heights[l] < heights[r] {
			newMax = (r - l) * int(math.Max(float64(heights[l]), float64(heights[l])))
			l++
		} else {
			newMax = (r - l) * int(math.Max(float64(heights[r]), float64(heights[r])))
			r--
		}

		if newMax > max {
			max = newMax
		}
	}

	return max
}
