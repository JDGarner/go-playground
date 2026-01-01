package intervals

// You are given an array of non-overlapping intervals intervals
// where intervals[i] = [start_i, end_i] represents the start and the
// end time of the ith interval. intervals is initially sorted in ascending
// order by start_i.

// You are given another interval newInterval = [start, end].

// Insert newInterval into intervals such that intervals is still sorted
// in ascending order by start_i and also intervals still does not have
// any overlapping intervals. You may merge the overlapping intervals if needed.

// Return intervals after adding newInterval.

// Note: Intervals are non-overlapping if they have no common point.
// For example, [1,2] and [3,4] are non-overlapping, but [1,2] and [2,3] are overlapping.

// Example 1:
// Input: intervals = [[1,3],[4,6]], newInterval = [2,5]
// Output: [[1,6]]

// Example 2:
// Input: intervals = [[1,2],[3,5],[9,10]], newInterval = [6,7]
// Output: [[1,2],[3,5],[6,7],[9,10]]

// -|--|-|--|
// --|--|----

// what if newInterval overlaps many intervals, e.g.:
// ----|-|-|-|-----
// -|------------|-
// [1, 10]

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := [][]int{}

	// Add every interval before an overlap
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// We've just hit an overlap at i
	// Expand the newInterval as long as:
	// - newInterval.end >= interval[i].start
	// e.g:
	// ---i|---------i|
	// -----------n|----n|
	// or:
	// ---i|---------i|
	// ------n|----n|
	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		newInterval = getMergedIntervals(newInterval, intervals[i])
		i++
	}

	res = append(res, newInterval)

	// No more overlaps from now
	// Add all the remaining intervals
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func getMergedIntervals(a, b []int) []int {
	min := a[0]
	if b[0] < min {
		min = b[0]
	}

	return []int{min, max(a[1], b[1])}
}

func insertFirstImpl(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	newStart, newEnd := newInterval[0], newInterval[1]

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	for i, interval := range intervals {
		start, end := interval[0], interval[1]

		latestEnd := 0
		if len(res) > 0 {
			latestEnd = res[len(res)-1][1]
		}

		switch {
		// - newInterval is entirely before interval
		//   => need to insert newInterval if not already in and return because there
		//      will be no more overlaps (also insert any remaining intervals)
		case newEnd < start:
			// if res is empty or latestEnd < newEnd then we haven't already merged newInterval in
			if len(res) == 0 || latestEnd < newEnd {
				res = append(res, newInterval)
			}

			// insert in the rest and return
			res = append(res, intervals[i:]...)

			return res

		// - newInterval is entirely after interval
		//   => insert interval[i]
		case newStart > end:
			res = append(res, interval)

			// if it is the last interval i, insert newInterval
			if i == len(intervals)-1 {
				res = append(res, newInterval)
			}

		// - newInterval is overlapping from the left/right (or one is fully containing the other)
		//   => insert merged intervals
		case newStart <= end || newEnd >= start:
			// if interval[i] is overlapping with previous res, update the latestEnd
			// otherwise merge them
			if len(res) > 0 && latestEnd >= start {
				res[len(res)-1][1] = max(latestEnd, end)
			} else {
				res = append(res, getMergedIntervals(interval, newInterval))
			}
		}
	}

	return res
}

func areOverlapping(a, b []int) bool {
	// if it's entirely before return false
	// - end and start of A is before start of B
	// ---a|---a|
	// -----------b|----b|
	if a[0] < b[0] && a[1] < b[0] {
		return false
	}

	// if it's entirely after return false
	// - start and end of A is after end of B
	// -----------a|----a|
	// ---b|---b|
	if a[1] > b[1] && a[0] > b[1] {
		return false
	}

	// Example contained case:
	// ---a|------------------a|
	// -----------b|----b|

	return true
}