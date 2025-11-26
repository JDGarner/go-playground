package binarysearchtree

// (This would normally be done on a regular tree rather than binary search tree)
func (b *BST) HasPathWithout(target int) bool {
	if b == nil {
		return true
	}

	if b.Value == target {
		return false
	}

	if b.IsLeafNode() {
		return true
	}

	if b.Left != nil && b.Left.HasPathWithout(target) {
		return true
	}
	if b.Right != nil && b.Right.HasPathWithout(target) {
		return true
	}

	return false
}
