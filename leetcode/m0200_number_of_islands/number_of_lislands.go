package numberofislands

// Time: O(M*N)
// Space: O(M*N)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, r, c int) {
	nr := len(grid)
	nc := len(grid[0])

	if (r < 0) || (r >= nr) || (c < 0) || (c >= nc) || (grid[r][c] == '0') {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
