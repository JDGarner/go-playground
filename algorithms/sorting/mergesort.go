package sorting

// Algorithm overview:
// Splits array in half, recursively calls MergeSort on each half then merges the two sorted
// arrays together
// Base case: if zero/one element => return array

// (In place version)
func MergeSort(data []int) {
	mergeSortHelper(data, 0, len(data))
}

func mergeSortHelper(data []int, start, end int) {
	if end-start <= 1 {
		return
	}

	mid := (start + end) / 2

	mergeSortHelper(data, start, mid)
	mergeSortHelper(data, mid, end)

	merge(data, start, mid, end)
}

// from start:mid is a sorted sub section of data
// from mid:end is a sorted sub section of data
// merge func merges these two sorted sub sections together so that start:end is sorted
func merge(data []int, start, mid, end int) {
	// Create a temporary slice to hold merged result
	temp := make([]int, end-start)

	i := start // pointer for left subarray
	j := mid   // pointer for right subarray
	k := 0     // pointer for temp array

	for i < mid && j < end {
		if data[i] <= data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
		k++
	}

	// Copy remaining elements from left subarray, if any
	for i < mid {
		temp[k] = data[i]
		i++
		k++
	}

	// Copy remaining elements from right subarray, if any
	for j < end {
		temp[k] = data[j]
		j++
		k++
	}

	// Copy merged elements back to original array
	for m, v := range temp {
		data[start+m] = v
	}
}

func MergeSortNotInPlace(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	merged := mergeNotInPlace(MergeSortNotInPlace(left), MergeSortNotInPlace(right))

	return merged
}

func mergeNotInPlace(sortedI, sortedJ []int) []int {
	i, j, k := 0, 0, 0
	result := make([]int, len(sortedI)+len(sortedJ))

	for i < len(sortedI) && j < len(sortedJ) {
		if sortedI[i] < sortedJ[j] {
			result[k] = sortedI[i]
			i++
		} else {
			result[k] = sortedJ[j]
			j++
		}
		k++
	}

	if i == len(sortedI) {
		copy(result[k:], sortedJ[j:])
	} else {
		copy(result[k:], sortedI[i:])
	}

	return result
}
