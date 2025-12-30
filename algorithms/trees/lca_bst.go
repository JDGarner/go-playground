package trees

// Given a binary search tree (BST) where all node values are unique,
// and two nodes from the tree p and q, return the lowest common ancestor (LCA)
// of the two nodes.

// The lowest common ancestor between two nodes p and q is the lowest node in a tree T
// such that both p and q as descendants.
// The ancestor is allowed to be a descendant of itself.

// Input: root = [5,3,8,1,4,7,9,null,2], p = 3, q = 8
// Output: 5

// Input: root = [5,3,8,1,4,7,9,null,2], p = 3, q = 4
// Output: 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestorBST(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	min, max := getMinMax(p, q)

	return lcaHelper(root, min, max)
}

func lcaHelper(root, min, max *TreeNode) *TreeNode {
	if max.Val >= root.Val && min.Val <= root.Val {
		return root
	}

	if min.Val <= root.Val && max.Val <= root.Val {
		return lcaHelper(root.Left, min, max)
	}

	return lcaHelper(root.Right, min, max)
}

func getMinMax(p *TreeNode, q *TreeNode) (min, max *TreeNode) {
	if p.Val < q.Val {
		return p, q
	}

	return q, p
}
