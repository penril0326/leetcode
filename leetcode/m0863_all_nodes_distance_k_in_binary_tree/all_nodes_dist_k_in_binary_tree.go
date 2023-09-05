package allnodesdistancekinbinarytree

import "practice/data_structure/node"

// Time: O(N), N is total number of nodes
// Space: O(N)
func distanceK(root *node.TreeNode, target *node.TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	graph := map[int]*node.TreeNode{}
	dfs(root, &graph)

	queue := make([]*node.TreeNode, 0)
	queue = append(queue, target)
	dist := k
	result := make([]int, 0)
	visited := make(map[int]bool)
	visited[target.Val] = true
	for len(queue) > 0 {
		depth := len(queue)
		dist--
		for i := 0; i < depth; i++ {
			cur := queue[0]
			queue = queue[1:]

			if (cur.Left != nil) && !visited[cur.Left.Val] {
				visited[cur.Left.Val] = true
				if dist == 0 {
					result = append(result, cur.Left.Val)
				} else {
					queue = append(queue, cur.Left)
				}
			}

			if (cur.Right != nil) && !visited[cur.Right.Val] {
				visited[cur.Right.Val] = true
				if dist == 0 {
					result = append(result, cur.Right.Val)
				} else {
					queue = append(queue, cur.Right)
				}
			}

			if parent, isExist := graph[cur.Val]; isExist && !visited[parent.Val] {
				visited[parent.Val] = true
				if dist == 0 {
					result = append(result, parent.Val)
				} else {
					queue = append(queue, parent)
				}
			}
		}

		if dist == 0 {
			return result
		}
	}

	return []int{}
}

func dfs(node *node.TreeNode, graph *map[int]*node.TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		(*graph)[node.Left.Val] = node
		dfs(node.Left, graph)
	}

	if node.Right != nil {
		(*graph)[node.Right.Val] = node
		dfs(node.Right, graph)
	}
}
