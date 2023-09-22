package detonatethemaximumbombs

// Time: O(N^3)
// Space: O(N^2)
func maximumDetonation(bombs [][]int) int {
	graph := make(map[int][]int)

	// Build graph: cost O(N^2)
	for i := 0; i < len(bombs); i++ {
		x1, y1, r1 := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := i + 1; j < len(bombs); j++ {
			x2, y2, r2 := bombs[j][0], bombs[j][1], bombs[j][2]
			dis := findDistance(x1, x2, y1, y2)
			if (r1 * r1) >= dis {
				graph[i] = append(graph[i], j)
			}

			if (r2 * r2) >= dis {
				graph[j] = append(graph[j], i)
			}
		}
	}

	maxComponents := 1

	// Run N times DFS
	for i, _ := range bombs {
		visited := make([]bool, len(bombs))

		// DFS: N nodes and at most we have N^2 edges.
		components := dfs(graph, i, visited)
		if components > maxComponents {
			maxComponents = components
		}
	}

	return maxComponents
}

func findDistance(x1, x2, y1, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func dfs(graph map[int][]int, b int, visited []bool) int {
	if visited[b] {
		return 0
	}

	visited[b] = true

	result := 1
	for _, neighbor := range graph[b] {
		if !visited[neighbor] {
			result += dfs(graph, neighbor, visited)
		}
	}

	return result
}
