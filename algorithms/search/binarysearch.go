package search

// Algorithm overview:
// Input must be in sorted order already.
// Find the value at midpoint, then:
// - if target equals the value => return index of this value
// - if the target is less than the value => repeat process on left half
// - if the target is more than the value => repeat process on right half

func BinarySearch(data []int, target int) (index int, found bool) {
	start, end := 0, len(data)-1

	for start <= end {
		midpoint := start + (end-start)/2

		if data[midpoint] == target {
			return midpoint, true
		}

		if target < data[midpoint] {
			end = midpoint - 1
		} else {
			start = midpoint + 1
		}
	}

	return -1, false
}

func BinarySearchFirstImpl(data []int, target int) (index int, ok bool) {
	if len(data) == 0 {
		return -1, false
	}

	start, end, midpoint := 0, len(data)-1, len(data)/2

	if data[midpoint] == target {
		return midpoint, true
	}

	for data[midpoint] != target {
		if start > end || start >= len(data) || end < 0 {
			return -1, false
		}

		if target < data[midpoint] {
			end = midpoint - 1
		} else {
			start = midpoint + 1
		}

		midpoint = (start + end) / 2
	}

	return midpoint, true
}
