package convertsortedarraytobst

import "practice/leetcode"

// binary search + recursive
func sortedArrayToBST(nums []int) *leetcode.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &leetcode.TreeNode{}
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
