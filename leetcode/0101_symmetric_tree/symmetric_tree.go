package symmetrictree_test

import "practice/data_structure/node"

func isSymmetric(root *node.TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(t1 *node.TreeNode, t2 *node.TreeNode) bool {
	if (t1 == nil) && (t2 == nil) {
		return true
	}

	if ((t1 == nil) && (t2 != nil)) ||
		(t1 != nil && t2 == nil) ||
		(t1.Val != t2.Val) {
		return false
	}

	if isMirror(t1.Left, t2.Right) == true {
		return isMirror(t1.Right, t2.Left)
	}

	return false
}
