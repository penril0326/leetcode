package e0112pathsum

import "practice/data_structure/node"

// Time: O(N)
// Space: O(N) for worst case, O(logN) for best case
func hasPathSum(root *node.TreeNode, targetSum int) bool {
	return pathSum(root, 0, targetSum)
}

func pathSum(node *node.TreeNode, curr, targetSum int) bool {
	if node == nil {
		return false
	}

	if (node.Left == nil) && (node.Right == nil) {
		return (node.Val+curr == targetSum)
	}

	curr += node.Val

	return pathSum(node.Left, curr, targetSum) || pathSum(node.Right, curr, targetSum)
}
