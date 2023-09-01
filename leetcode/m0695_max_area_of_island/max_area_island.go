package maxareaofisland

import "practice/utility/math"

// Time: O(R*C), R is total of rows, C is total of columns
// Space: O(R*C)
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			maxArea = math.Max(maxArea, dfs(grid, i, j))
		}
	}

	return maxArea
}

func dfs(grid [][]int, i int, j int) int {
	if (i < 0) || (i >= len(grid)) || (j < 0) || (j >= len(grid[0])) || (grid[i][j] == 0) {
		return 0
	}

	grid[i][j] = 0

	return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
}
