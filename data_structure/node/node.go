package node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type NaryNode struct {
	Val      int
	Children []*NaryNode
}
