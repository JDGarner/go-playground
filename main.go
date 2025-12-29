package main

import (
	"github.com/JDGarner/go-playground/algorithms"
	"github.com/JDGarner/go-playground/algorithms/graph/waterflow"
	"github.com/JDGarner/go-playground/concurrency"
	"github.com/JDGarner/go-playground/datastructures/binaryheap"
	"github.com/JDGarner/go-playground/datastructures/graph"
	"github.com/JDGarner/go-playground/datastructures/hashmap"
	"github.com/JDGarner/go-playground/datastructures/linkedlist"
)

func main() {
	AlgorithmExamples()
	// DataStructureExamples()
	// ConcurrencyExamples()
}

func ConcurrencyExamples() {
	// ***************************************
	// Fan In Pattern
	// ***************************************
	// concurrency.FanInExample()
	// concurrency.FanInLoggerExample()

	// ***************************************
	// Fan Out Pattern
	// ***************************************
	concurrency.FanOutExample()
	// concurrency.FanOutBroadcastExample()
	// concurrency.FanOutJobsExample()

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
	// SearchExamples()
	// BinarySearchTreeExamples()
	// rottenfruit.RottenFruitExample()
	// dynamic.CountPathsExample()
	// courseschedule.CanFinishExample()
	// errgroup.Example()
	// islands.Example()
	waterflow.Example()
	// array.ProductExceptSelfExample()
	// slidingwindow.CharacterReplacementExample()
	// array.RotatedSortedExample()
	// trees.KthSmallestNode()
	// backtracking.WordSearchExample()
	// backtracking.WordSearch2Example()
	// backtracking.CombinationSumExample()
	// regions.SurroundedRegionsExample()
	// numberboard.SnakesAndLaddersExample()
	// array.SpiralMatrixExample()
}

func SortingExamples() {
	// algorithms.BucketSortExample()
	// algorithms.InsertionSortExample()
	// algorithms.MergeSortExample()
	// algorithms.QuickSortExample()
}

func SearchExamples() {
	// algorithms.BinarySearchExample()
	// algorithms.BSTSearchExample()
}

func DataStructureExamples() {
	// BinarySearchTreeExamples()
	// LinkedListExamples()
	// BinaryHeapExamples()
	// HashMapExamples()
	GraphExamples()
}

func BinarySearchTreeExamples() {
	// algorithms.BSTInsertAndRemovalExample()
	// algorithms.BSTDFSTraversalExample()
	// algorithms.BSTBFSTraversalExample()
	algorithms.BacktrackingExample()
}

func BinaryHeapExamples() {
	// binaryheap.PushAndPopExample()
	binaryheap.HeapifyExample()
}

func HashMapExamples() {
	hashmap.Example()
}

func LinkedListExamples() {
	linkedlist.Example()
}

func GraphExamples() {
	// graph.MatrixDFSExample()
	// graph.MatrixBFSExample()
	graph.AdjacencyListExample()
}
