package numberofprovinces

// Time: O(N^2)
// Space: O(N)
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	provinces := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			dfs(i, isConnected, visited)
		}
	}

	return provinces
}

func dfs(city int, isConnected [][]int, visited []bool) {
	visited[city] = true

	for i, j := range isConnected[city] {
		if !visited[i] && j == 1 {
			dfs(i, isConnected, visited)
		}
	}
}
