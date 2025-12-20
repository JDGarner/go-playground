package trees

import (
	"github.com/JDGarner/go-playground/datastructures/binarysearchtree"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary search tree, and an integer k,
// return the kth smallest value (1-indexed) in the tree.

//     4
//   3   5
// 2
// k=2, output = 3
// k=4, output = 5

func kthSmallestFinal(root *binarysearchtree.BST, k int) int {
	kSmallestNodes := inOrderTraversal(root, []*binarysearchtree.BST{}, k)

	return kSmallestNodes[k-1].Value
}

func inOrderTraversal(node *binarysearchtree.BST, acc []*binarysearchtree.BST, k int) []*binarysearchtree.BST {
	if node == nil {
		return acc
	}

	acc = inOrderTraversal(node.Left, acc, k)
	acc = append(acc, node)
	if len(acc) == k {
		return acc
	}
	acc = inOrderTraversal(node.Right, acc, k)

	return acc
}

func kthSmallestFirstImpl(root *TreeNode, k int) int {
	var inOrderTraversal func(node *TreeNode)
	var nodes []*TreeNode

	inOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		if len(nodes) == k {
			return
		}

		inOrderTraversal(node.Left)
		nodes = append(nodes, node)
		inOrderTraversal(node.Right)
	}

	inOrderTraversal(root)

	return nodes[k-1].Val
}

func kthSmallestWithTN(root *TreeNode, k int) int {
	kSmallestNodes := inOrderTraversalTN(root, []*TreeNode{}, k)

	return kSmallestNodes[k-1].Val
}

func inOrderTraversalTN(node *TreeNode, acc []*TreeNode, k int) []*TreeNode {
	if node == nil || len(acc) == k {
		return acc
	}

	acc = inOrderTraversalTN(node.Left, acc, k)
	if len(acc) < k {
		acc = append(acc, node)
	}
	acc = inOrderTraversalTN(node.Right, acc, k)

	return acc
}
