package minimumabsolutedifferenceinbst

import (
	"math"
	"practice/data_structure/node"
)

// Time: O(N)
// Space: O(N), N for call stack, another N for sorted list, so total is 2N.
func getMinimumDifference(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	list := []int{}
	dfs(root, &list)

	min := math.MaxInt
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		if diff < min {
			min = diff
		}
	}

	return min
}

func dfs(node *node.TreeNode, l *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, l)
	*l = append(*l, node.Val)
	dfs(node.Right, l)
}

// Time: O(N)
// Space: O(N), call stack for N node
func getMinimumDifference2(root *node.TreeNode) int {
	// DFS without a list
	if root == nil {
		return 0
	}

	prev, min := -1, math.MaxInt
	dfs2(root, &prev, &min)

	return min
}

func dfs2(node *node.TreeNode, prev, min *int) {
	if node == nil {
		return
	}

	dfs2(node.Left, prev, min)

	if *prev != -1 {
		diff := abs(*prev, node.Val)
		if diff < *min {
			*min = diff
		}
	}

	*prev = node.Val
	dfs2(node.Right, prev, min)
}

func abs(a, b int) int {
	res := a - b
	if res > 0 {
		return res
	}

	return -res
}
