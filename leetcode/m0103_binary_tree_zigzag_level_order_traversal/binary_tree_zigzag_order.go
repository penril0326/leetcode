package binarytreezigzaglevelordertraversal

import "practice/data_structure/node"

// Time: O(N)
// Space: O(N)
func zigzagLevelOrder(root *node.TreeNode) [][]int {
	// BFS
	if root == nil {
		return [][]int{}
	}

	queue := []*node.TreeNode{root}
	ans := [][]int{}
	isOrderLeft := true
	for len(queue) > 0 {
		currentLevel := len(queue)

		tmp := make([]int, currentLevel)
		for i := 0; i < currentLevel; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if isOrderLeft {
				tmp[i] = node.Val
			} else {
				tmp[currentLevel-i-1] = node.Val
			}
		}

		ans = append(ans, tmp)
		isOrderLeft = !isOrderLeft
	}

	return ans
}

// Time: O(N)
// Space: O(N)
func zigzagLevelOrderDFS(root *node.TreeNode) [][]int {
	// DFS
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}
	dfs(root, 0, &ans)
	return ans
}

func dfs(n *node.TreeNode, depth int, ans *[][]int) {
	if n == nil {
		return
	}

	// if current level is nil
	if depth >= len(*ans) {
		*ans = append(*ans, []int{})
	}

	// handle current node
	if depth%2 == 0 {
		(*ans)[depth] = append((*ans)[depth], n.Val)
	} else {
		(*ans)[depth] = append([]int{n.Val}, (*ans)[depth]...)
	}

	// traval from left to right
	dfs(n.Left, depth+1, ans)
	dfs(n.Right, depth+1, ans)
}
