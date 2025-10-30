package sorting_test

import (
	"slices"
	"testing"

	"github.com/JDGarner/go-playground/algorithms/sorting"
)

type testCase struct {
	input    []int
	expected []int
}

func TestInsertionSort(t *testing.T) {
	testcases := map[string]testCase{
		"sorting case 1": {
			input:    []int{6, 18, 3, 1, 7},
			expected: []int{1, 3, 6, 7, 18},
		},
		"sorting case 2": {
			input:    []int{1, 1, 19, 2, 1},
			expected: []int{1, 1, 1, 2, 19},
		},
		"sorting case 3": {
			input:    []int{42},
			expected: []int{42},
		},
		"sorting case 4": {
			input:    []int{42, 7},
			expected: []int{7, 42},
		},
	}

	for label, testcase := range testcases {
		t.Run(label, func(t *testing.T) {
			actual := sorting.InsertionSort(testcase.input)
			if !slices.Equal(actual, testcase.expected) {
				t.Fatalf("slices not equal: got %v, want %v", actual, testcase.expected)
			}
		})
	}
}
