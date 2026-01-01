package intervals

// You are given a sorted list of disjoint intervals intervals [a, b]
// You are also given another interval toBeRemoved.

// Return the set of real numbers with the interval toBeRemoved
// removed from intervals.
// In other words, return the set of real numbers such that every x
// in the set is in intervals but not in toBeRemoved.
// Your answer should be a sorted list of disjoint intervals as described above.

// Example 1:
// Input: intervals = [[0,2],[3,4],[5,7]], toBeRemoved = [1,6]
// Output: [[0,1],[6,7]]

// Example 2:
// Input: intervals = [[0,5]], toBeRemoved = [2,3]
// Output: [[0,2],[3,5]]

// Example 3:
// Input: intervals = [[-5,-4],[-3,-2],[1,2],[3,5],[8,9]], toBeRemoved = [-1,4]
// Output: [[-5,-4],[-3,-2],[4,5],[8,9]]

// Basically any interval that overlaps completely:
// => remove it
// For any interval that overlaps partially:
// => include the bit that doesn't overlap with toBeRemoved
// For any interval that doesn't overlap at all:
// => include it

// ----i|---i|---i|---i|---i|---i|-i|---i|----
// ------r|-------------------------r|---

func remove(intervals [][]int, toBeRemoved []int) [][]int {
	res := [][]int{}
	removeStart, removeEnd := toBeRemoved[0], toBeRemoved[1]

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// No overlap - interval is completely before removal
		if end < removeStart {
			res = append(res, interval)
			continue
		}

		// No overlap - interval is completely after removal
		if start > removeEnd {
			res = append(res, interval)
			continue
		}

		// Partial overlap - keep left part if it exists
		// --i|-----i|-------
		// ------r|-----r|---
		// output:
		// --i|--r|----------
		if start < removeStart {
			res = append(res, []int{start, removeStart})
		}

		// Partial overlap - keep right part if it exists
		// ----------i|-----i|---
		// ------r|-----r|-------
		// output:
		// -------------r|--i|---
		if end > removeEnd {
			res = append(res, []int{removeEnd, end})
		}

		// This case will have been handled by the two if blocks above:
		// ---i|-----------i|---
		// -------r|--r|--------
		// output:
		// ---i|--r|--r|---i|---

		// This case will have been ignored:
		// ------i|---i|-------
		// ---r|----------r|---
		// output:
		// - nothing

	}

	return res
}

func removeFirstImpl(intervals [][]int, toBeRemoved []int) [][]int {

	i, n := 0, len(intervals)
	res := [][]int{}

	// loop through any intervals that are completely BEFORE toBeRemoved
	for i < n && intervals[i][1] < toBeRemoved[0] { // end before toBeRemoved.start
		res = append(res, intervals[i])
		i++
	}

	// Overlap case 1:
	// --i|-----i|-------
	// ------r|-----r|---
	// output:
	// --i|--r|----------

	// Overlap case 2:
	// ------i|---i|-------
	// ---r|----------r|---
	// output:
	// - nothing

	// Overlap case 3:
	// ----------i|-----i|---
	// ------r|-----r|-------
	// output:
	// -------------r|--i|---

	// Overlap case 4:
	// ---i|-----------i|---
	// -------r|--r|--------
	// output:
	// ---i|--r|--r|---i|---

	// loop through ALL overlapping intervals
	// keep going while end of toBeRemoved is after or equal to start of i
	for i < n && toBeRemoved[1] >= intervals[i][0] {
		// if toBeRemoved completely contains i (case 2):
		if toBeRemoved[0] <= intervals[i][0] && toBeRemoved[1] >= intervals[i][1] {
			i++
			continue
		}

		// if i completely contains toBeRemoved (case 4):
		if intervals[i][0] < toBeRemoved[0] && intervals[i][1] > toBeRemoved[1] {
			res = append(res, []int{intervals[i][0], toBeRemoved[0]})
			res = append(res, []int{toBeRemoved[1], intervals[i][1]})
			i++
			continue
		}

		// toBeRemoved overlaps on right side (case 1):
		if toBeRemoved[0] <= intervals[i][1] && toBeRemoved[1] >= intervals[i][1] { // it wraps the end of i
			res = append(res, []int{intervals[i][0], toBeRemoved[0]})
			i++
			continue
		}

		// toBeRemoved overlaps on left side (case 3):
		res = append(res, []int{toBeRemoved[1], intervals[i][1]})
		i++
	}

	// add remaining intervals
	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}
