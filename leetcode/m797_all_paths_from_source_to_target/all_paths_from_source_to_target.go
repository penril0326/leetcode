package allpathsfromsourcetotarget

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	backtracking(graph, []int{0}, 0, len(graph)-1, &ans)
	return ans
}

func backtracking(graph [][]int, curr []int, start, end int, ans *[][]int) {
	if start == end {
		new := make([]int, len(curr))
		copy(new, curr)
		*ans = append(*ans, new)
		return
	}

	for _, node := range graph[start] {
		curr = append(curr, node)
		backtracking(graph, curr, node, end, ans)
		curr = curr[:len(curr)-1]
	}
}
