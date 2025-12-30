package trees

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(current, target *TreeNode, path *[]*TreeNode) bool

	dfs = func(current, target *TreeNode, path *[]*TreeNode) bool {
		if current == nil {
			return false
		}

		*path = append(*path, current)

		if current.Val == target.Val {
			return true
		}

		// explore left and right
		if dfs(current.Left, target, path) {
			return true
		}
		if dfs(current.Right, target, path) {
			return true
		}

		// if it wasn't found in either, remove current from path
		*path = (*path)[:len(*path)-1]

		return false
	}

	pPath := []*TreeNode{}
	qPath := []*TreeNode{}

	dfs(root, p, &pPath)
	dfs(root, q, &qPath)

	result := &TreeNode{}

	// p = [1, 2]
	// q = [1]

	// iterate through pPath and qPath until they diverge, return that node
	for i, v := range pPath {
		if i == len(qPath) {
			break
		}

		if qPath[i] != v {
			break
		}
		result = v
	}

	return result
}
