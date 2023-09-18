package nearestexitfromentranceinmaze

// Time: O(m*n)
// Space: O(m+n)
func nearestExit(maze [][]byte, entrance []int) int {
	rLen, cLen := len(maze), len(maze[0])
	visited := make([][]bool, rLen)
	for i := 0; i < rLen; i++ {
		visited[i] = make([]bool, cLen)
	}

	type state struct {
		r    int
		c    int
		step int
	}

	queue := []state{}
	queue = append(queue, state{entrance[0], entrance[1], 0})
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		r, c, step := queue[0].r, queue[0].c, queue[0].step
		queue = queue[1:]
		if isOnBorder(r, c, rLen, cLen) {
			if !isOnEntrance(r, c, entrance[0], entrance[1]) {
				return step
			}
		}
		for _, dir := range direction {
			neighborR, neighborC := r+dir[0], c+dir[1]
			if inBound(neighborR, neighborC, rLen, cLen) {
				if maze[neighborR][neighborC] == '.' && !visited[neighborR][neighborC] {
					visited[neighborR][neighborC] = true
					queue = append(queue, state{neighborR, neighborC, step + 1})
				}
			}
		}
	}

	return -1
}

func inBound(x, y, lr, lc int) bool {
	return (x >= 0) && (x < lr) && (y >= 0) && (y < lc)
}

func isOnBorder(x, y, lr, lc int) bool {
	return (x == 0) || (y == 0) || (x == lr-1) || (y == lc-1)
}

func isOnEntrance(x, y, entranceX, entranceY int) bool {
	return (x == entranceX) && (y == entranceY)
}
