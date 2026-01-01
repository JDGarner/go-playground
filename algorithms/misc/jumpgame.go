package misc

// You are given an array of integers nums, where nums[i] represents the
// maximum length of a jump towards the right from index i.
// For example, if you are at nums[i], you can jump to any index i + j where:

// j <= nums[i]
// i + j < nums.length
// You are initially positioned at nums[0].

// Return the minimum number of jumps to reach the last position in the
// array (index nums.length - 1).
// You may assume there is always a valid answer.

// Example 1:
// Input: nums = [2,4,1,1,1,1]
// Output: 2

// Explanation: Jump from index 0 to index 1, then jump from index 1 to the last index.

// Example 2:
// Input: nums = [2,1,2,1,0]
// Output: 2

// Maintain two pointers, l and r, initially set to 0, representing the range
// of reachable indices.
// At each step, we iterate through the indices in the range l to r and
// determine the farthest index that can be reached from the current range.

// Example 3:
// [2,4,1,1,1,1,1,1]
// l=0, r=2, steps=1

// l=2, r=5, steps=2
// l=6, r=6, steps=3
// l=7, r=7, steps=4

func jump(nums []int) int {
	left := 0
	right := 0
	steps := 0

	for right < len(nums)-1 {
		newRight := right
		for i := left; i <= right; i++ {
			if i+nums[i] > newRight {
				newRight = i + nums[i]
			}
		}
		left = right + 1
		steps++
		right = newRight
	}

	return steps
}

// [2, 7, 2, 1, 0, 0, 0, 0, 0]
// i=0, farthest=2, previousFarthest=2, jumps=1
// i=1, farthest=8, previousFarthest=2, jumps=1
// i=2, farthest=8, previousFarthest=8, jumps=2

// Take the first jump
// Set farthest to 0+num[0] (e.g. 2)
// incr jump and store farthest in a temp variable (previousFarthest)
// - Loop from 0 to previousFarthest and see how far you can jump to from any
//   of those indices, store that in farthest
//   - once you get to previousFarthest (end of previous farthest jump), increment
//     jumps and set new previousFarthest to farthest
//     - do the same again until you hit previousFarthest again

func jump2(nums []int) int {
	jumps := 0
	farthest := 0
	previousFarthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == previousFarthest {
			jumps++
			previousFarthest = farthest
		}
	}

	return jumps
}
