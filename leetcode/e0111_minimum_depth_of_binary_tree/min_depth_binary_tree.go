package minimumdepthofbinarytree

import "practice/data_structure/node"

// DFS
// Time: O(N)
// Space: O(N)
func minDepth(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// BFS
// Time: O(N)
// Space: O(N)
func minDepth2(root *node.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*node.TreeNode{}
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			cur := queue[0]
			if (cur.Left == nil) && (cur.Right == nil) {
				return depth
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			size--
			queue = queue[1:]
		}
		depth++
	}

	return depth
}
