package snakesandladders

import "practice/utility/math"

// Time: O(n^2)
// Space: O(n^2)
func snakesAndLadders(board [][]int) int {
	n := len(board)
	newGraph := make([]int, n*n+1)
	label := 1
	column := make([]int, n)
	for i := 0; i < n; i++ {
		column[i] = i
	}

	for i := n - 1; i >= 0; i-- {
		for _, j := range column {
			newGraph[label] = board[i][j]
			label++
		}

		for j := 0; j < n/2; j++ {
			column[j], column[n-j-1] = column[n-j-1], column[j]
		}

	}

	type state struct {
		curr int
		step int
	}

	queue := make([]state, 0)
	queue = append(queue, state{1, 0})
	visited := make([]bool, n*n+1)
	for len(queue) > 0 {
		cur, step := queue[0].curr, queue[0].step
		queue = queue[1:]
		if cur == n*n {
			return step
		}

		for next := cur + 1; next <= math.Min(cur+6, n*n); next++ {
			destination := next
			if newGraph[next] != -1 {
				destination = newGraph[next]
			}

			if !visited[destination] {
				queue = append(queue, state{destination, step + 1})
				visited[destination] = true
			}
		}
	}

	return -1
}
