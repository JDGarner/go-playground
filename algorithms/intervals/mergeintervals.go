package intervals

import (
	"slices"
)

// Given an array of intervals where intervals[i] = [start_i, end_i],
// merge all overlapping intervals, and return an array of the non-overlapping
// intervals that cover all the intervals in the input.

// You may return the answer in any order.

// Note: Intervals are non-overlapping if they have no common point.
// For example, [1, 2] and [3, 4] are non-overlapping,
// but [1, 2] and [2, 3] are overlapping.

// Example 1:
// Input: intervals = [[1,3],[1,5],[6,7]]
// Output: [[1,5],[6,7]]

// Example 2:
// Input: intervals = [[1,2],[2,3]]
// Output: [[1,3]]

// If two intervals are sorted in ascending order by their start values,
// they overlap if the start value of the second interval is less than or
// equal to the end value of the first interval.

func merge(intervals [][]int) [][]int {
	// Sort by starting interval
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})

	res := [][]int{
		intervals[0],
	}

	// init with the first interval
	// go through remaining intervals
	// if the start of interval[i] is more than or equal to the most recent end
	// - put on the max of interval[i][1] and lastEnd
	// else
	// - no overlap so create a new entry in res

	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]
		lastEnd := res[len(res)-1][1]

		if start <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, end)
		} else {
			res = append(res, []int{start, end})
		}
	}

	return res
}
