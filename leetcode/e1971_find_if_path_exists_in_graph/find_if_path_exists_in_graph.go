package findifpathexistsingraph

// Time: O(n+m), n is total of verticss, m is total of edges
// Space: O(n+m)
func validPath(n int, edges [][]int, source int, destination int) bool {
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
	dfs(source, graph, visited)

	return visited[destination]
}

func dfs(source int, graph map[int][]int, visited []bool) {
	visited[source] = true
	for _, neighbor := range graph[source] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}
