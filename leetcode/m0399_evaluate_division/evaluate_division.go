package evaluatedivision

// BFS
// Time: O(N*(E+V)), N is size of queries; E is total of edges; V is total of vertices.
// Space: O(E+V)
func calcEquationBFS(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for idx, n := range equations {
		x, y := n[0], n[1]
		if graph[x] == nil {
			graph[x] = make(map[string]float64)
		}
		graph[x][y] = values[idx]

		if graph[y] == nil {
			graph[y] = make(map[string]float64)
		}
		graph[y][x] = 1.0 / values[idx]
	}

	type state struct {
		node   string
		weight float64
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, target := q[0], q[1]

		// Node not in the graph
		if (graph[start] == nil) || (graph[target] == nil) {
			ans[i] = -1
			continue
		}

		visited := map[string]bool{}
		queue := make([]state, 0)
		queue = append(queue, state{start, 1.0})
		visited[q[0]] = true

		// Standard BFS
		for len(queue) > 0 {
			n, ratio := queue[0].node, queue[0].weight
			queue = queue[1:]
			if n == target {
				ans[i] = ratio
				break
			}

			for neighbor, w := range graph[n] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, state{neighbor, ratio * w})
				}
			}
		}

		// If we finish the BFS and ans[i] still is a default value, which mean we didn't find the path in the given query
		if ans[i] == 0 {
			ans[i] = -1
		}
	}

	return ans
}

// DFS
// Time: O(N*(E+V)), N is size of queries; E is total of edges; V is total of vertices.
// Space: O(E+V)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for idx, n := range equations {
		x, y := n[0], n[1]
		if graph[x] == nil {
			graph[x] = make(map[string]float64)
		}
		graph[x][y] = values[idx]

		if graph[y] == nil {
			graph[y] = make(map[string]float64)
		}
		graph[y][x] = 1.0 / values[idx]
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, target := q[0], q[1]
		visited := map[string]bool{}
		ans[i] = dfs(graph, 1.0, start, target, &visited)
	}

	return ans
}

func dfs(graph map[string]map[string]float64, ratio float64, node string, target string, visited *map[string]bool) float64 {
	if (*visited)[node] {
		return -1.0
	}

	(*visited)[node] = true

	for neighbor, weight := range graph[node] {
		if neighbor == target {
			return ratio * weight
		}

		result := dfs(graph, ratio*weight, neighbor, target, visited)
		if result != -1 {
			return result
		}
	}

	return -1.0
}
