package maximumdiffbetweennodeandancestor

import (
	"practice/data_structure/node"
	"practice/utility/math"
)

// Time: O(N)
// Space: O(N)
func maxAncestorDiff(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, root.Val, root.Val)
}

func dfs(node *node.TreeNode, curMax, curMin int) int {
	if node == nil {
		return curMax - curMin
	}

	curMax = math.Max(node.Val, curMax)
	curMin = math.Min(node.Val, curMin)
	left := dfs(node.Left, curMax, curMin)
	right := dfs(node.Right, curMax, curMin)

	return math.Max(left, right)
}
