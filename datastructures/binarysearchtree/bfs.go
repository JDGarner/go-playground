package binarysearchtree

import "github.com/JDGarner/go-playground/datastructures/queue"

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
