package intervals

// You are given two lists of closed intervals, firstList and secondList,
// where firstList[i] = [start[i], end[i]] and secondList[j] = [start[j], end[j]].
// Each list of intervals is pairwise disjoint and in sorted order.

// Return the intersection of these two interval lists.

// A closed interval [a, b] (with a <= b) denotes the set of real numbers
// x with a <= x <= b.

// The intersection of two closed intervals is a set of real numbers that are
// either empty or represented as a closed interval.
// For example, the intersection of [1, 3] and [2, 4] is [2, 3].

// Example 1:
// Input: firstList = [[0,2],[5,10],[13,23],[24,25]],
// 				secondList = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

// At any point where both first and second are present => create interval

// Case 1 - second overlaps first later:
// |-----|----
// ---|-----|-
// output:
// ---|--|----

// Case 2 - no overlap:
// |-----|----
// ---------|-----|-
// output:
// skip

// Case 3 - first overlaps second later:
// ---|-----|-
// |-----|----
// output:
// ---|--|----

// Case 4:
// ---|--|----
// -|------|----

// Case 5:
// -|------|----
// ---|--|----

// after checking case:
// increment intervals pointer on the one that ends earlier

// iterate through until one pointer is at the end of their list

// Example 2:
// Input: firstList = [[1,3],[5,9]], secondList = []
// Output: []

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	firstP, secondP := 0, 0
	res := [][]int{}

	for firstP < len(firstList) && secondP < len(secondList) {
		firstStart, firstEnd := firstList[firstP][0], firstList[firstP][1]
		secondStart, secondEnd := secondList[secondP][0], secondList[secondP][1]

		// if there is any overlap:
		// - take the higher of the starts and smaller of the ends
		if areOverlapping2(firstList[firstP], secondList[secondP]) {
			start := max(firstStart, secondStart)
			end := min(firstEnd, secondEnd)

			res = append(res, []int{start, end})
		}

		// increment intervals pointer on the one that ends earlier
		if firstEnd < secondEnd {
			firstP++
		} else {
			secondP++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func areOverlapping2(a, b []int) bool {
	// if a is entirely ahead of b
	if a[0] > b[1] {
		return false
	}

	// if b is entirely ahead of a
	if b[0] > a[1] {
		return false
	}

	return true
}

func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	firstP, secondP := 0, 0
	res := [][]int{}

	for firstP < len(firstList) && secondP < len(secondList) {
		firstStart, firstEnd := firstList[firstP][0], firstList[firstP][1]
		secondStart, secondEnd := secondList[secondP][0], secondList[secondP][1]

		// if there is any overlap:
		// - take the higher of the starts and smaller of the ends
		start := max(firstStart, secondStart)
		end := min(firstEnd, secondEnd)

		if start <= end {
			res = append(res, []int{start, end})
		}

		// increment intervals pointer on the one that ends earlier
		if firstEnd < secondEnd {
			firstP++
		} else {
			secondP++
		}
	}

	return res
}


// Case 1 - second overlaps first later:
// |-----|----
// ---|-----|-
// output:
// ---|--|----

// Case 2 - no overlap:
// |-----|----
// ---------|-----|-
// output:
// skip

// Case 3 - first overlaps second later:
// ---|-----|-
// |-----|----
// output:
// ---|--|----

// Case 4:
// ---|--|----
// -|------|----

// Case 5:
// -|------|----
// ---|--|----