package maximum_depth

import "practice/leetcode"

func maxDepth(root *leetcode.TreeNode) int {
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
