package m1448countgoodnodesinbinarytree

import (
	"math"
	"practice/data_structure/node"
)

// Time: O(n)
// Space: O(n)
func goodNodes(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	dfs(root, math.MinInt32, &count)

	return count
}

func dfs(node *node.TreeNode, max int, count *int) {
	if node == nil {
		return
	}

	if node.Val >= max {
		*count++
		max = node.Val
	}

	dfs(node.Left, max, count)
	dfs(node.Right, max, count)
}
