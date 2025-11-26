package binarysearchtree

func (b *BST) DFSTraversal(f func(value int)) {
	if b == nil {
		return
	}

	b.Left.DFSTraversal(f)
	f(b.Value)
	b.Right.DFSTraversal(f)
}
