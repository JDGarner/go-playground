package intervals

import (
	"cmp"
	"slices"
)

// Given an array of intervals intervals where intervals[i] = [start_i, end_i],
// return the minimum number of intervals you need to remove to make the
// rest of the intervals non-overlapping.

// Note: Intervals are non-overlapping even if they have a common point.
// For example, [1, 3] and [2, 4] are overlapping, but [1, 2] and [2, 3]
// are non-overlapping.

// Example 1:
// Input: intervals = [[1,2],[2,4],[1,4]]
// Output: 1

// -|-|--------------
// -|----|-----------
// ---|--|-----------
// output:
// 1

// Explanation: After [1,4] is removed, the rest of the intervals are non-overlapping.

// Example 2:
// Input: intervals = [[1,2],[2,4]]
// Output: 0

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// We then iterate through the sorted intervals from left to right,
	// keeping track of the previous interval’s end value as prevEnd,
	// initially set to the end value of the first interval.

	prevEnd := intervals[0][1]

	// -|-|--------------
	// -|----|-----------
	// ---|--|-----------

	count := 0

	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]

		// If overlap, skip the one with the higher end value
		// we can skip it just by setting the prevEnd to the smaller one so that
		// we just ignore the interval that stretches on for longer
		if start < prevEnd {
			// i should be 'skipped'
			// set prevEnd to the smaller of the two
			prevEnd = min(end, prevEnd)
			count++

			continue
		} else {
			// Else it's not overlapping so update the prevEnd to the end of the new intervl
			prevEnd = end
		}
	}

	return count
}
