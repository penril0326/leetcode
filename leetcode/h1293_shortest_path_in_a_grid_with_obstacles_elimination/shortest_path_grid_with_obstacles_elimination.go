package shortestpathinagridwithobstacleselimination

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	seen := make([][][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			seen[i][j] = make([]bool, k+1)
		}
	}

	type state struct {
		r      int
		c      int
		steps  int
		remain int
	}

	queue := make([]state, 0)
	queue = append(queue, state{0, 0, 0, k})
	seen[0][0][k] = true
	direction := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		row, col, step, remain := cell.r, cell.c, cell.steps, cell.remain

		if (row == m-1) && (col == n-1) {
			return step
		}

		for _, dir := range direction {
			neighborx, neighbory := row+dir[0], col+dir[1]
			if isInBound(neighborx, neighbory, m, n) {
				if (grid[neighborx][neighbory] == 0) && !seen[neighborx][neighbory][remain] {
					seen[neighborx][neighbory][remain] = true
					queue = append(queue, state{neighborx, neighbory, step + 1, remain})
				} else if (remain > 0) && !seen[neighborx][neighbory][remain-1] {
					seen[neighborx][neighbory][remain-1] = true
					queue = append(queue, state{neighborx, neighbory, step + 1, remain - 1})
				}
			}
		}
	}

	return -1
}

func isInBound(i, j, boundR, boundC int) bool {
	return (i >= 0) && (i < boundR) && (j >= 0) && (j < boundC)
}
