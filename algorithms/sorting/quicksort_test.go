package sorting_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/JDGarner/go-playground/algorithms/sorting"
	"github.com/JDGarner/go-playground/algorithms/sorting/helpers"
)

func TestQuickSort(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("quick sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			actual := sorting.QuickSort(inputCopy)
			if !slices.Equal(actual, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", actual, testcase.Expected)
			}
		})
	}
}

func TestQuickSortInPlace(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("quick sort in place: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			sorting.QuickSortInPlace(inputCopy)
			if !slices.Equal(inputCopy, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", inputCopy, testcase.Expected)
			}
		})
	}
}
