package validatebst

import "practice/data_structure/node"

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
