package sorting

// Algorithm overview:
// Splits array in half, recursively calls MergeSort on each half then merges the two sorted
// arrays together
// Base case: if zero/one element => return array

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	merged := merge(MergeSort(left), MergeSort(right))

	return merged
}

func merge(sortedI, sortedJ []int) []int {
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
