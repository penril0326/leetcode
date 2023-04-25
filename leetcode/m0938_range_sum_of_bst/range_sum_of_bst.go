package rangesumofbst

import "practice/data_structure/node"

// Time: O(N)
// Space: O(N)
func rangeSumBST(root *node.TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}

	if low < root.Val {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
