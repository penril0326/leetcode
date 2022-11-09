package maximum_depth

import "practice/data_structure/node"

func maxDepth(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
