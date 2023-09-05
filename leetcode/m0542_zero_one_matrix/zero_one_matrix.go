package zeroonematrix

// Time: O(m*n)
// Space: O(m*n)
func updateMatrix(mat [][]int) [][]int {
	ans := make([][]int, len(mat))
	seen := make([][]bool, len(mat))
	for i := 0; i < len(mat); i++ {
		ans[i] = make([]int, len(mat[i]))
		seen[i] = make([]bool, len(mat[i]))
	}

	type coordinate struct {
		r int
		c int
	}

	queue := make([]coordinate, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				seen[i][j] = true
				queue = append(queue, coordinate{
					r: i,
					c: j,
				})
			}
		}
	}

	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		x, y := queue[0].r, queue[0].c
		queue = queue[1:]
		for _, dir := range direction {
			nextR, nextC := x+dir[0], y+dir[1]
			if isInBound(nextR, nextC, len(mat), len(mat[0])) && !seen[nextR][nextC] {
				seen[nextR][nextC] = true
				queue = append(queue, coordinate{
					r: nextR,
					c: nextC,
				})
				ans[nextR][nextC] = ans[x][y] + 1
			}
		}
	}

	return ans
}

func isInBound(i, j int, lengthR, lengthC int) bool {
	return (i >= 0) && (i < lengthR) && (j >= 0) && (j < lengthC)
}
