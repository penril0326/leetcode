package numberofconnectedcomponentsinanundirectedgraph

// Time: O(E+V), E is total of edges, V is total of vertices
// Space: O(E+V)
func countComponents(n int, edges [][]int) int {
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
	components := 0
	for cur := 0; cur < n; cur++ {
		if !visited[cur] {
			components++
			dfs(cur, graph, visited)
		}
	}

	return components
}

func dfs(node int, graph map[int][]int, visited []bool) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}
