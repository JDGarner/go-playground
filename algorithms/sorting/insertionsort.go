package sorting

// Algorithm overview:
// Loop through data starting from 2nd index (1st element is already a sorted sub-array)
// For each element:
// - Iterate backwards, swapping items unless current item is more than previous item

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] > data[j-1] {
				break
			}
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}
