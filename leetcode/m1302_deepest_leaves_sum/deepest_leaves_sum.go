package deepestleavessum

import (
	"container/list"
	"practice/data_structure/node"
)

// Time: O(N)
// Space: O(D), which D is the diameter of the tree
// BFS
func deepestLeavesSum(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	sum := 0
	for queue.Len() > 0 {
		currentLevel := queue.Len()

		tempSum := 0
		for i := 0; i < currentLevel; i++ {
			node := queue.Remove(queue.Front()).(*node.TreeNode)
			tempSum += node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		sum = tempSum
	}

	return sum
}

// Time: O(N)
// Space: O(D), which D is the diameter of the tree
// DFS
func deepestLeavesSumDFS(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	dict := map[int]int{}

	dfs(root, 0, &dict)

	return dict[len(dict)-1]
}

func dfs(n *node.TreeNode, depth int, m *map[int]int) {
	if n == nil {
		return
	}

	(*m)[depth] += n.Val

	dfs(n.Left, depth+1, m)
	dfs(n.Right, depth+1, m)
}
