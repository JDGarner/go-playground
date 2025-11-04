package sorting

import (
	"math"
)

// Algorithm Overview
// Put elements into a series of buckets that hold a range of values (e.g. 2 buckets: 1-5, 6-10)
// Sort those buckets separately (e.g. with insertion sort)
// Concatenate those buckets (or replace values of the original array in place)
// Note - numberOfBuckets could be constant or could be calculated as part of the algorithm

func BucketSort(data []int) {
	if len(data) <= 1 {
		return
	}

	min, max := findMinMax(data)

	// Examples:
	// if there are 36 elements, number of buckets is 6
	numberOfBuckets := int(math.Ceil(math.Sqrt(float64(len(data)))))

	// Examples:
	// if values are 1-60 and there are 6 buckets, bucket range is 10
	// if values are 5-10 and there are 2 buckets, bucket range is 2.5
	bucketRange := float64((max - min + 1)) / float64(numberOfBuckets)

	buckets := make([][]int, numberOfBuckets)

	// Sort data into buckets
	for _, v := range data {
		bucketIndex := int(float64((v - min)) / bucketRange)
		if bucketIndex >= numberOfBuckets {
			bucketIndex = numberOfBuckets - 1
		}
		buckets[bucketIndex] = append(buckets[bucketIndex], v)
	}

	i := 0

	// Sort each bucket
	for _, bucket := range buckets {
		InsertionSort(bucket) // Use insertion sort as it's good for small data sets

		// Put sorted bucket elements back into original slice
		for _, v := range bucket {
			data[i] = v
			i++
		}
	}
}

func findMinMax(data []int) (int, int) {
	min := data[0]
	max := data[0]

	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}
