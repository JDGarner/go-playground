package search

import (
	"github.com/JDGarner/go-playground/datastructures/binarysearchtree"
)

func BSTSearch(bst *binarysearchtree.BST, target int) (found bool) {
	if bst == nil {
		return false
	}

	if bst.Value == target {
		return true
	}

	if target < bst.Value {
		return BSTSearch(bst.Left, target)
	}

	return BSTSearch(bst.Right, target)
}

func BSTSearchNonRecursive(bst *binarysearchtree.BST, target int) (found bool) {
	if bst == nil {
		return false
	}

	if bst.Value == target {
		return true
	}

	currentNode := bst

	for currentNode != nil && !currentNode.IsLeafNode() {
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		if currentNode != nil && currentNode.Value == target {
			return true
		}
	}

	return false
}
