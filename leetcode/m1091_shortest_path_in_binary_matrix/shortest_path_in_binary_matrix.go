package shortestpathinbinarymatrix

// Time: O(N^2)
// Space: O(N), additional space is queue, and we only store next depth nodes.
func shortestPathBinaryMatrix(grid [][]int) int {
	N := len(grid)
	if (grid[0][0] == 1) || (grid[N-1][N-1] == 1) {
		return -1
	}

	queue := make([]node, 0)
	queue = append(queue, node{
		i: 0,
		j: 0,
	})
	path := 0

	for len(queue) > 0 {
		level := len(queue)
		path++
		for i := 0; i < level; i++ {
			n := queue[0]
			queue = queue[1:]

			if (n.i == N-1) && (n.j == N-1) {
				return path
			}

			addNode(n.i-1, n.j-1, grid, &queue)
			addNode(n.i, n.j-1, grid, &queue)
			addNode(n.i+1, n.j-1, grid, &queue)
			addNode(n.i-1, n.j, grid, &queue)
			addNode(n.i+1, n.j, grid, &queue)
			addNode(n.i-1, n.j+1, grid, &queue)
			addNode(n.i, n.j+1, grid, &queue)
			addNode(n.i+1, n.j+1, grid, &queue)
		}
	}

	return -1
}

type node struct {
	i int
	j int
}

func addNode(i, j int, grid [][]int, queue *[]node) {
	if i < len(grid) && j < len(grid[0]) && i >= 0 && j >= 0 && grid[i][j] == 0 {
		(*queue) = append(*queue, node{
			i: i,
			j: j,
		})

		grid[i][j] = 1
	}
}
