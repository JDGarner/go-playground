package search_test

import (
	"fmt"
	"testing"

	"github.com/JDGarner/go-playground/algorithms/search"
	"github.com/JDGarner/go-playground/algorithms/search/helpers"
	"github.com/JDGarner/go-playground/datastructures/binarysearchtree"
)

func TestBSTSearch(t *testing.T) {
	for label, testcase := range helpers.SearchTestCases {
		t.Run(fmt.Sprintf("search in binary search tree: %s", label), func(t *testing.T) {
			bst := binarysearchtree.NewFromList(testcase.Input)

			found := search.BSTSearch(bst, testcase.Target)
			if found != testcase.Found {
				t.Fatalf("found not equal: got %v, want %v", found, testcase.Found)
			}
		})
	}
}

func TestBSTSearchNonRecursive(t *testing.T) {
	for label, testcase := range helpers.SearchTestCases {
		t.Run(fmt.Sprintf("search in binary search tree (non recursive): %s", label), func(t *testing.T) {
			bst := binarysearchtree.NewFromList(testcase.Input)

			found := search.BSTSearchNonRecursive(bst, testcase.Target)
			if found != testcase.Found {
				t.Fatalf("found not equal: got %v, want %v", found, testcase.Found)
			}
		})
	}
}
