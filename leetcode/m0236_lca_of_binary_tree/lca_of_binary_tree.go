package m0236lcaofbinarytree

import "practice/data_structure/node"

// Time: O(N)
// Space: O(N)
func lowestCommonAncestor(root, p, q *node.TreeNode) *node.TreeNode {
	if root == nil {
		return nil
	}

	if (root == p) || (root == q) {
		return root
	}

	findLeft := lowestCommonAncestor(root.Left, p, q)
	findRight := lowestCommonAncestor(root.Right, p, q)

	if (findLeft != nil) && (findRight != nil) {
		return root
	} else if findLeft != nil {
		return findLeft
	} else {
		return findRight
	}
}
