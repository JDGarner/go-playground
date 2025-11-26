package binarysearchtree

// (This would normally be done on a regular tree rather than binary search tree)
func (b *BST) HasPathWithout(target int) bool {
	if b == nil || b.Value == target {
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

// This function returns the first path found that does not contain the target
// (if no path, returns empty slice and false)
func (b *BST) FindPathWithout(target int) ([]int, bool) {
	if b == nil || b.Value == target {
		return []int{}, false
	}

	if b.IsLeafNode() {
		return []int{b.Value}, true
	}

	if b.Left != nil {
		path, ok := b.Left.FindPathWithout(target)
		if ok {
			return append([]int{b.Value}, path...), true
		}
	}
	if b.Right != nil {
		path, ok := b.Right.FindPathWithout(target)
		if ok {
			return append([]int{b.Value}, path...), true
		}
	}

	return []int{}, false
}
