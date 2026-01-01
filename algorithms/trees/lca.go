package trees

// Note:
// Not for a BST just a regular binary tree
// -----------------------------------------------

// if p and q are not the left and right of the root,
// that means the LCA is either the left of the root or the right of the root.

// call lowestCommonAncestor() with the new root i.e., root.Left
// and hold the returned node left. Do the same for root.Right and
// hold the returned node right.

//      5
// 3        1

// p = 1
// q = 3

// left := lowestCommonAncestor(3, p, q)
//         = 3
// right := lowestCommonAncestor(1, p, q)
//         = 1

// if they both dont equal nil then the parent is their root
// => 5

//      5
// 3        6
//              1
// if one of them is nil, return the other (in this case )


// Example 2: One Node is Ancestor of the Other
// Find LCA of nodes 5 and 4
//         3
//        / \
//       5   1
//      / \
//     6   2
//    / \
//   7   4
// Call 1: LCA(3, 5, 4)
// In this example when we recurse left, we find 5 immediately and return it
// left = 5
// we don't need to bother looking at the rest of the left tree because if it's not
// in the right subtree then 5 is the LCA
// we then recurse right which ends up returning nil
// so then in Call 1 we return left (5)

// Note:
// If both subtrees return non-nil → current node is the LCA (split point)
// If only one subtree returns non-nil → that node is either:
// - The LCA (one is ancestor of the other), or
// - Just passing the signal upward

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if we found p or q, return this node! (or if it's an empty node return nil)
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// if both are non nil then the root must be the LCA because it is directly above them both
	if left != nil && right != nil {
		return root
	}
	// if we found p or q in the right subtree, return that
	if right != nil {
		return right
	}

	// else return left (this may be nil also if we didn't find p or q in either subtree)
	return left
}

// DFS for p and q separately and return the path for each
// find where the two paths diverge
func lowestCommonAncestorFirstImpl(root, p, q *TreeNode) *TreeNode {
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
