package sorting

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	pivotIndex := len(data) - 1
	pivot := data[len(data)-1]
	pivotBoundary := 0

	// Swap anything less than or equal to pivot to the pivotBoundary and increment pivotBoundary
	// Outcome => anything less than pivot will be on left side of pivotBoundary, anything more than will be on right side
	for i := 0; i < len(data)-1; i++ {
		if data[i] <= pivot {
			data[i], data[pivotBoundary] = data[pivotBoundary], data[i]
			pivotBoundary++
		}
	}

	// Swap pivot element into place
	data[pivotIndex], data[pivotBoundary] = data[pivotBoundary], data[pivotIndex]

	left := QuickSort(data[:pivotBoundary])
	right := QuickSort(data[pivotBoundary:])

	return append(left, right...)
}

func QuickSortInPlace(data []int) {
	quickSortHelper(data, 0, len(data)-1)
}

func quickSortHelper(data []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := end
	pivot := data[pivotIndex]
	pivotBoundary := start

	// Swap anything less than or equal to pivot to the pivotBoundary and increment pivotBoundary
	// Outcome => anything less than pivot will be on left side of pivotBoundary, anything more than will be on right side
	for i := start; i < end; i++ {
		if data[i] <= pivot {
			data[i], data[pivotBoundary] = data[pivotBoundary], data[i]
			pivotBoundary++
		}
	}

	// Swap pivot element into place
	data[pivotIndex], data[pivotBoundary] = data[pivotBoundary], data[pivotIndex]

	quickSortHelper(data, start, pivotBoundary-1)
	quickSortHelper(data, pivotBoundary+1, end)
}
