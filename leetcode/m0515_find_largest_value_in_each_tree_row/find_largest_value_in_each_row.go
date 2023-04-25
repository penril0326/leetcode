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

// Time: O(N)
// Space: O(D), which D is the diameter of the tree
func largestValuesDFS(root *node.TreeNode) []int {
	// DFS solution
	if root == nil {
		return []int{}
	}

	dict := map[int]int{}
	dfs(root, 0, &dict)
	ans := make([]int, len(dict))
	for k, v := range dict {
		ans[k] = v
	}

	return ans
}

func dfs(n *node.TreeNode, depth int, m *map[int]int) {
	if n == nil {
		return
	}

	if value, isExist := (*m)[depth]; !isExist {
		(*m)[depth] = n.Val
	} else {
		if n.Val > value {
			(*m)[depth] = n.Val
		}
	}

	dfs(n.Left, depth+1, m)
	dfs(n.Right, depth+1, m)
}
