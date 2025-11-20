package main

import (
	"github.com/JDGarner/go-playground/algorithms"
	"github.com/JDGarner/go-playground/concurrency"
)

func main() {
	AlgorithmExamples()
	// ConcurrencyExamples()
}

func ConcurrencyExamples() {
	// ***************************************
	// Fan In Pattern
	// ***************************************
	// concurrency.FanInExample()
	concurrency.FanInLoggerExample()

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

func AlgorithmExamples() {
	// SortingExamples()
	SearchExamples()
}

func SortingExamples() {
	algorithms.BucketSortExample()
	algorithms.InsertionSortExample()
	algorithms.MergeSortExample()
	algorithms.QuickSortExample()
}

func SearchExamples() {
	// algorithms.BinarySearchExample()
	// algorithms.BSTSearchExample()
	// algorithms.BSTInsertAndRemovalExample()
	// algorithms.BSTDFSTraversalExample()
	algorithms.BSTBFSTraversalExample()
}
