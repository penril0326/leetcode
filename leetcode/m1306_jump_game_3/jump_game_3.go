package jumpgame3

// Time: O(n), n is length of arr
// Space: O(n)
func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))

	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]
		if arr[index] == 0 {
			return true
		}

		neighbors := []int{index + arr[index], index - arr[index]}
		for _, neighbor := range neighbors {
			if isValidNeighbor(neighbor, len(arr)) && !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false
}

func isValidNeighbor(idx, lenght int) bool {
	return (idx >= 0) && (idx < lenght)
}
