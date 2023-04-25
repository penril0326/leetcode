package binarytreerightsideview

import (
	"container/list"
	"practice/data_structure/node"
)

// Time: O(N)
// Space: O(D), which D is tree diameter.
// For instance, a completely binary tree [1, 2, 3, 4, 5, 6, 7] has a diameter 4, the longest distance is from node 4 to node 7.
func rightSideView(root *node.TreeNode) []int {
	// BFS
	if root == nil {
		return []int{}
	}

	l := list.New()
	l.PushBack(root)

	result := []int{}
	for l.Len() != 0 {
		currentLevel := l.Len()
		front := l.Front().Value.(*node.TreeNode)
		result = append(result, front.Val)

		for i := 0; i < currentLevel; i++ {
			node := l.Remove(l.Front()).(*node.TreeNode)
			if node.Right != nil {
				l.PushBack(node.Right)
			}

			if node.Left != nil {
				l.PushBack(node.Left)
			}
		}
	}

	return result
}

// Time: O(N)
// Space: O(D), which D is tree diameter.
func rightSideViewDFS(root *node.TreeNode) []int {
	// DFS
	if root == nil {
		return []int{}
	}

	dict := map[int]int{}
	dfs(root, 0, &dict)
	result := make([]int, len(dict))
	for k, v := range dict {
		result[k] = v
	}

	return result
}

func dfs(n *node.TreeNode, depth int, m *map[int]int) {
	if n == nil {
		return
	}

	(*m)[depth] = n.Val

	dfs(n.Left, depth+1, m)
	dfs(n.Right, depth+1, m)
}
