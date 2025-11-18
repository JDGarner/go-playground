package binarysearchtree

import (
	"fmt"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func New(value int) *BST {
	return &BST{
		Value: value,
	}
}

func (b *BST) SetLeft(v int) error {
	if v == b.Value {
		return nil // Ignore duplicates
	}

	if v > b.Value {
		return fmt.Errorf("left value must be less than its parent value (%d)", b.Value)
	}

	b.Left = New(v)
	return nil
}

func (b *BST) SetRight(v int) error {
	if v == b.Value {
		return nil // Ignore duplicates
	}

	if v < b.Value {
		return fmt.Errorf("right value must be more than its parent value (%d)", b.Value)
	}

	b.Right = New(v)
	return nil
}

func (b *BST) IsLeafNode() bool {
	return b.Left == nil && b.Right == nil
}

func (b *BST) Insert(value int) *BST {
	if b == nil {
		return New(value)
	}

	if value < b.Value {
		b.Left = b.Left.Insert(value)
	} else {
		b.Right = b.Right.Insert(value)
	}

	return b
}

func (b *BST) FindMin() int {
	if b.Left == nil {
		return b.Value
	}

	return b.Left.FindMin()
}

func (b *BST) FindMax() int {
	if b.Right == nil {
		return b.Value
	}

	return b.Right.FindMax()
}

func (b *BST) IsValid() bool {
	if b.Left != nil && b.Value < b.Left.Value {
		return false
	}

	if b.Right != nil && b.Value > b.Right.Value {
		return false
	}

	leftValid := true
	if b.Left != nil {
		leftValid = b.Left.IsValid()
	}

	rightValid := true
	if b.Right != nil {
		rightValid = b.Right.IsValid()
	}

	return rightValid && leftValid
}

// TODO: redo:
func (b *BST) Remove(value int) {
	node := b.Find(value)
	if node == nil {
		return
	}

	parent := b.FindParentNode(value, b)
	if parent == nil {
		return
	}

	if node.Value == parent.Value {
		// TODO: how to handle root?
	}

	// Node is on parent's Right
	if value > parent.Value {
		if node.Right != nil {
			replacement := node.Right
			parent.Right = replacement
			replacement.Left = node.Left
			node = nil
			return
		}

		if node.Left != nil {
			replacement := node.Left
			parent.Right = replacement
			replacement.Right = node.Right
			node = nil
			return
		}

		// If we are here, it must be a leaf node
		parent.Right = nil
		node = nil
		return
	}

	// Node is on parent's Left

	if node.Right != nil {
		replacement := node.Right
		parent.Left = replacement
		replacement.Left = node.Left
		node = nil
		return
	}

	if node.Left != nil {
		replacement := node.Left
		parent.Left = replacement
		replacement.Right = node.Right
		node = nil
		return
	}

	// If we are here, it must be a leaf node
	parent.Left = nil
}

func (b *BST) Find(value int) *BST {
	if b == nil {
		return nil
	}

	if value == b.Value {
		return b
	}

	if value < b.Value {
		return b.Left.Find(value)
	}

	return b.Right.Find(value)
}

func (b *BST) FindParentNode(value int, parent *BST) *BST {
	if b == nil {
		return parent
	}

	if value == b.Value {
		return parent
	}

	if value < b.Value {
		return b.Left.FindParentNode(value, b)
	}

	return b.Right.FindParentNode(value, b)
}

func NewFromList(data []int) *BST {
	if len(data) == 0 {
		return nil
	}

	return buildTree(data, 0, len(data)-1)
}

func buildTree(data []int, start, end int) *BST {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	bst := New(data[mid])

	// Build left tree
	bst.Left = buildTree(data, start, mid-1)
	// Build right tree
	bst.Right = buildTree(data, mid+1, end)

	return bst
}
