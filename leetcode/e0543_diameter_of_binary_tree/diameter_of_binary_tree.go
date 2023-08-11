package diameterofbinarytree

import (
	"practice/data_structure/node"
	"practice/utility/math"
)

// Time: O(N)
// Space: O(N)
func diameterOfBinaryTree(root *node.TreeNode) int {
	dia := 0
	dfs(root, &dia)

	return dia
}

func dfs(node *node.TreeNode, maxDia *int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, maxDia)
	right := dfs(node.Right, maxDia)

	*maxDia = math.Max(*maxDia, left+right)
	return math.Max(left, right) + 1
}
