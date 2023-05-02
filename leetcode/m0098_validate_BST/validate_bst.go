package validatebst

import (
	"math"
	"practice/data_structure/node"
)

// Time: O(N)
// Space: O(2N) = O(N)
// Inorder traversal approach
func isValidBST(root *node.TreeNode) bool {
	if root == nil {
		return false
	}

	list := make([]int, 0)
	inOrder(root, &list)

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}

	return true
}

func inOrder(node *node.TreeNode, list *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, list)
	*list = append(*list, node.Val)
	inOrder(node.Right, list)
}

// Time: O(N)
// Space: O(N)
// DFS range check
func isValidBST2(root *node.TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *node.TreeNode, low, high int) bool {
	if node == nil {
		return true
	}

	if (node.Val <= low) || (node.Val >= high) {
		return false
	}

	isLeftBST := dfs(node.Left, low, node.Val)
	isRightBST := dfs(node.Right, node.Val, high)

	return isLeftBST && isRightBST
}
