package convertsortedarraytobst

import "practice/data_structure/node"

// binary search + recursive
func sortedArrayToBST(nums []int) *node.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &node.TreeNode{}
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
