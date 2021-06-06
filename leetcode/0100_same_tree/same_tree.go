package same_tree

import "practice/leetcode"

func isSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	if (p == nil) && (q == nil) {
		return true
	} else if ((p != nil) && (q == nil)) ||
		((p == nil) && (q != nil)) {
		return false
	}

	if p.Val != q.Val {
		return false
	} else if isSameTree(p.Left, q.Left) == true {
		return isSameTree(p.Right, q.Right)
	}

	return false
}
