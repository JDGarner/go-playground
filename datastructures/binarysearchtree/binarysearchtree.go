package binarysearchtree

import (
	"fmt"

	"github.com/JDGarner/go-playground/datastructures/queue"
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

func (b *BST) DFSTraversal(f func(value int)) {
	if b == nil {
		return
	}

	b.Left.DFSTraversal(f)
	f(b.Value)
	b.Right.DFSTraversal(f)
}

func (b *BST) BFSTraversal(f func(value int)) {
	if b == nil {
		return
	}

	f(b.Value)
	children := getImmediateChildren([]*BST{b})

	for len(children) > 0 {
		for _, child := range children {
			f(child.Value)
		}

		children = getImmediateChildren(children)
	}
}

func getImmediateChildren(nodes []*BST) (children []*BST) {
	for _, node := range nodes {
		if node.Left != nil {
			children = append(children, node.Left)
		}
		if node.Right != nil {
			children = append(children, node.Right)
		}
	}

	return children
}

func (b *BST) BFSTraversalWithQueue(f func(value int)) {
	if b == nil {
		return
	}

	q := queue.New[*BST]()
	q.Enqueue(b)

	for q.Len() > 0 {
		node := q.Dequeue()
		f(node.Value)

		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
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

func (b *BST) FindMin() *BST {
	if b.Left == nil {
		return b
	}

	return b.Left.FindMin()
}

func (b *BST) FindMax() *BST {
	if b.Right == nil {
		return b
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

func (b *BST) Remove(value int) *BST {
	if b == nil {
		return nil
	}

	if value < b.Value {
		b.Left = b.Left.Remove(value)
		return b
	} else if value > b.Value {
		b.Right = b.Right.Remove(value)
		return b
	}
	// if we are here we have found the node to remove

	// CASE 1 - Leaf node
	if b.IsLeafNode() {
		return nil // return nil to replace current node with nil (no need to worry about any child nodes with leaf node)
	}

	// CASE 2 - Node has only 1 child
	if b.Left == nil || b.Right == nil {
		// If node has only a right child, then return this to replace the current node
		if b.Right != nil {
			return b.Right
		}

		return b.Left // Otherwise it only has a left child, return that
	}

	// CASE 3 - Node has 2 children
	min := b.Right.FindMin()
	b.Value = min.Value // Replace node's value with the minimum from the right tree

	b.Right = b.Right.Remove(min.Value) // Now remove that value from the tree
	return b
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
