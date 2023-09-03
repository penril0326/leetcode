package reachablenodeswithrestrictions

// Time: O(N)
// Space: O(N)
func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if _, isExist := graph[x]; !isExist {
			graph[x] = make([]int, 0)
		}
		graph[x] = append(graph[x], y)

		if _, isExist := graph[y]; !isExist {
			graph[y] = make([]int, 0)
		}
		graph[y] = append(graph[y], x)
	}

	visited := make([]bool, n)
	for _, resNode := range restricted {
		visited[resNode] = true
	}

	return dfs(0, graph, visited)
}

func dfs(node int, graph map[int][]int, visited []bool) int {
	if visited[node] {
		return 0
	}

	visited[node] = true
	reachCount := 1
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			reachCount += dfs(neighbor, graph, visited)
		}
	}

	return reachCount
}
