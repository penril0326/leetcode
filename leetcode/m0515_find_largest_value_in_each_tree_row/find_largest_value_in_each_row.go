package findlargestvalueineachtreerow

import (
	"container/list"
	"practice/data_structure/node"
)

// Time: O(N)
// Space: O(D), which D is the diameter of the tree
func largestValues(root *node.TreeNode) []int {
	// BFS solution
	if root == nil {
		return []int{}
	}

	l := list.New()
	l.PushBack(root)
	ans := []int{}
	for l.Len() != 0 {
		currentLevel := l.Len()
		max := l.Front().Value.(*node.TreeNode).Val

		for i := 0; i < currentLevel; i++ {
			node := l.Remove(l.Front()).(*node.TreeNode)
			if max < node.Val {
				max = node.Val
			}

			if node.Left != nil {
				l.PushBack(node.Left)
			}

			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}

		ans = append(ans, max)
	}

	return ans
}
