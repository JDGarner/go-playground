package main

import (
	"fmt"

	"github.com/JDGarner/go-playground/algorithms/search"
	"github.com/JDGarner/go-playground/algorithms/sorting"
	"github.com/JDGarner/go-playground/concurrency"
)

func main() {
	// ConcurrencyExamples()
	// SortingExamples()
	// SearchExamples()
}

func ConcurrencyExamples() {
	// ***************************************
	// Fan In Pattern
	// ***************************************
	concurrency.FanInExample()

	// ***************************************
	// Generator Pattern
	// ***************************************
	// concurrency.TickerGeneratorExample()
	// concurrency.FibGeneratorExample()
	// concurrency.DoubleGeneratorExample()

	// ***************************************
	// Misc
	// ***************************************
	// concurrency.TickerWithDoneChannel()
	// concurrency.CancellableExample()
	// concurrency.FirstResponseExample()
	// concurrency.AllResponsesExample()
}

func SortingExamples() {
	toBucketSort := getDataToSort()
	sorting.BucketSort(toBucketSort)
	fmt.Println("bucket sorted: ", toBucketSort)

	toInsertionSort := getDataToSort()
	sorting.InsertionSort(toInsertionSort)
	fmt.Println("insertion sorted: ", toInsertionSort)

	toMergeSort := getDataToSort()
	sorting.MergeSort(toMergeSort)
	fmt.Println("merge sorted: ", toMergeSort)

	toQuickSort := getDataToSort()
	sorting.QuickSort(toQuickSort)
	fmt.Println("quick sorted: ", toQuickSort)
}

func getDataToSort() []int {
	return []int{100, 101, 116, 107, 111, 115, 115, 110, 106, 103, 100, 116, 104}
}

func SearchExamples() {
	toSearch := []int{-10, -5, -3, -2, -1}
	target := -10
	result, found := search.BinarySearch(toSearch, target)
	fmt.Printf("binary search %v for %d: result: %d, found: %v\n", toSearch, target, result, found)
}
