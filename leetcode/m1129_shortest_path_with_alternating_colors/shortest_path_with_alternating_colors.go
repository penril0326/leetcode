package shortestpathwithalternatingcolors

import (
	"math"
	mymath "practice/utility/math"
)

// Time: O(n+e), n is total of nodes and e is total of edges
// Space: O(n+e)
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	const (
		RED  = 0
		BLUE = 1
	)

	graph := make(map[int]map[int][]int)
	graph[RED] = make(map[int][]int)
	graph[BLUE] = make(map[int][]int)

	for _, edge := range redEdges {
		from, to := edge[0], edge[1]
		graph[RED][from] = append(graph[RED][from], to)
	}

	for _, edge := range blueEdges {
		from, to := edge[0], edge[1]
		graph[BLUE][from] = append(graph[BLUE][from], to)
	}

	type state struct {
		node     int
		lastEdge int
		step     int
	}

	queue := make([]state, 0)
	queue = append(queue, state{0, RED, 0}, state{0, BLUE, 0})

	visited := make([][2]bool, n)
	visited[0][RED] = true
	visited[0][BLUE] = true

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = math.MaxInt
	}

	for len(queue) > 0 {
		node, edgeColor, step := queue[0].node, queue[0].lastEdge, queue[0].step
		queue = queue[1:]
		ans[node] = mymath.Min(ans[node], step)

		for _, neighbor := range graph[edgeColor][node] {
			if !visited[neighbor][1-edgeColor] {
				visited[neighbor][1-edgeColor] = true
				queue = append(queue, state{neighbor, 1 - edgeColor, step + 1})
			}
		}
	}

	for i := 0; i < n; i++ {
		if ans[i] == math.MaxInt {
			ans[i] = -1
		}
	}

	return ans
}
