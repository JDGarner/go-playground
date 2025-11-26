package algorithms

import (
	"fmt"

	"github.com/JDGarner/go-playground/algorithms/search"
	"github.com/JDGarner/go-playground/algorithms/sorting"
	"github.com/JDGarner/go-playground/datastructures/binarysearchtree"
)

// ------------------------------------------
// SORTING
// ------------------------------------------
func BucketSortExample() {
	toSort := getDataToSort()
	sorting.BucketSort(toSort)
	fmt.Println("bucket sorted:", toSort)
}

func InsertionSortExample() {
	toSort := getDataToSort()
	sorting.InsertionSort(toSort)
	fmt.Println("insertion sorted:", toSort)
}

func MergeSortExample() {
	toSort := getDataToSort()
	sorting.MergeSort(toSort)
	fmt.Println("merge sorted:", toSort)
}

func QuickSortExample() {
	toSort := getDataToSort()
	sorting.QuickSort(toSort)
	fmt.Println("quick sorted:", toSort)
}

func getDataToSort() []int {
	return []int{100, 101, 116, 107, 111, 115, 115, 110, 106, 103, 100, 116, 104}
}

// ------------------------------------------
// SEARCH
// ------------------------------------------
func BinarySearchExample() {
	toSearch := []int{-10, -5, -3, -2, -1}
	target := -10
	result, found := search.BinarySearch(toSearch, target)
	fmt.Printf("binary search %v for %d: result: %d, found: %v\n", toSearch, target, result, found)
}

func BSTSearchExample() {
	bst := binarysearchtree.New(5)
	bst.SetLeft(3)
	bst.Left.SetLeft(1)
	bst.SetRight(7)
	bst.Right.SetLeft(6)
	bst.Right.SetRight(19)

	fmt.Println(bst.String())

	target := 19
	found := search.BSTSearch(bst, target)
	fmt.Printf("search in binary search tree for %d: found: %v\n", target, found)

	bst = binarysearchtree.NewFromList([]int{0, 1, 6, 7, 18})

	fmt.Println(bst.String())

	target = 18
	found = search.BSTSearch(bst, target)
	fmt.Printf("search in binary search tree for %d: found: %v\n", target, found)
	target = 42
	found = search.BSTSearch(bst, target)
	fmt.Printf("search in binary search tree for %d: found: %v\n", target, found)
}

func BSTInsertAndRemovalExample() {
	bst := binarysearchtree.NewFromList([]int{0, 1, 6, 7, 18})
	fmt.Println("original:")
	fmt.Println(bst.String())
	fmt.Println("min: ", bst.FindMin().Value)
	fmt.Println("max: ", bst.FindMax().Value)

	fmt.Println("-----------------------------")
	fmt.Println("insert 8, 27, 4, 3, 5, 17")
	bst.Insert(8)
	bst.Insert(27)
	bst.Insert(4)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(17)
	bst.Insert(-4)

	fmt.Println("after insertion:")
	fmt.Println(bst.String())

	fmt.Println("min: ", bst.FindMin().Value)
	fmt.Println("max: ", bst.FindMax().Value)

	fmt.Println("-----------------------------")
	fmt.Println("remove 17, 5, 3, 4, 27, 8, -4")
	bst.Remove(17)
	bst.Remove(5)
	bst.Remove(3)
	bst.Remove(4)
	bst.Remove(27)
	bst.Remove(8)
	bst.Remove(-4)

	fmt.Println("after removal:")
	fmt.Println(bst.String())
}

func BSTDFSTraversalExample() {
	bst := binarysearchtree.NewFromList([]int{-4, 0, 1, 3, 4, 5, 6, 7, 8, 17, 18, 27})
	fmt.Println("tree:")
	fmt.Println(bst.String())

	fmt.Println("dfs traversal:")
	bst.DFSTraversal(func(value int) {
		fmt.Println("traversing: ", value)
	})
}

func BSTBFSTraversalExample() {
	bst := binarysearchtree.NewFromList([]int{-4, 0, 1, 3, 4, 5, 6, 7, 8, 17, 18, 27})
	fmt.Println("tree:")
	fmt.Println(bst.String())

	fmt.Println("bfs traversal:")
	bst.BFSTraversal(func(value int) {
		fmt.Println("traversing: ", value)
	})

	fmt.Println("bfs traversal using queue:")
	bst.BFSTraversalWithQueue(func(value int) {
		fmt.Println("traversing: ", value)
	})
}

func BacktrackingExample() {
	bst := binarysearchtree.NewFromList([]int{27, 3, 2, 27, 7, 27, 19, 27, 10})

	fmt.Println(bst)

	without := bst.HasPathWithout(27)
	fmt.Println("has path without 27: ", without)
	without = bst.HasPathWithout(42)
	fmt.Println("has path without 42: ", without)
	without = bst.HasPathWithout(3)
	fmt.Println("has path without 3: ", without)

	path, found := bst.FindPathWithout(27)
	fmt.Println("path without 27: ", path, found)
	path, found = bst.FindPathWithout(42)
	fmt.Println("path without 42: ", path, found)
	path, found = bst.FindPathWithout(3)
	fmt.Println("path without 3: ", path, found)
}
