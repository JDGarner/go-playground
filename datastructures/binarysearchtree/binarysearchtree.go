package binarysearchtree

import (
	"fmt"
	"strings"
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
	if v >= b.Value {
		return fmt.Errorf("left value must be less than its parent value (%d)", b.Value)
	}

	b.Left = New(v)
	return nil
}

func (b *BST) SetRight(v int) error {
	if v < b.Value {
		return fmt.Errorf("right value must be more than or equal to its parent value (%d)", b.Value)
	}

	b.Right = New(v)
	return nil
}

func (b *BST) IsLeafNode() bool {
	return b.Left == nil && b.Right == nil
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

// (AI generated) String returns a pretty-printed representation of the binary search tree
func (b *BST) String() string {
	if b == nil {
		return ""
	}

	var sb strings.Builder
	b.printTree(&sb, "", "", true)
	return sb.String()
}

// (AI generated) printTree recursively builds the tree string with proper formatting
func (b *BST) printTree(sb *strings.Builder, prefix, childPrefix string, isRoot bool) {
	if b == nil {
		return
	}

	// Print current node
	if isRoot {
		sb.WriteString(fmt.Sprintf("%d\n", b.Value))
	} else {
		sb.WriteString(fmt.Sprintf("%s%d\n", prefix, b.Value))
	}

	// Determine if we have children
	hasLeft := b.Left != nil
	hasRight := b.Right != nil

	// Print left subtree
	if hasLeft {
		if hasRight {
			// More children coming, use ├──
			b.Left.printTree(sb, childPrefix+"├── ", childPrefix+"│   ", false)
		} else {
			// Last child, use └──
			b.Left.printTree(sb, childPrefix+"└── ", childPrefix+"    ", false)
		}
	}

	// Print right subtree
	if hasRight {
		// Right child is always last
		b.Right.printTree(sb, childPrefix+"└── ", childPrefix+"    ", false)
	}
}
