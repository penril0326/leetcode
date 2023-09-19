package openthelock

import "strings"

// Time: O(N^2*A^N+d), N is number of slots, in this case is 4 slots;
// A is number of digits in alphabet, in this case is 0 ~ 9; d is length of deadends
// Space: O(A^N)
func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
	}

	if visited["0000"] {
		return -1
	}

	type state struct {
		curr string
		step int
	}

	queue := make([]state, 0)
	queue = append(queue, state{"0000", 0})

	for len(queue) > 0 {
		curr, step := queue[0].curr, queue[0].step
		queue = queue[1:]
		if curr == target {
			return step
		}

		for _, neighbor := range getNeighbors(curr) {
			if !visited[neighbor] {
				queue = append(queue, state{neighbor, step + 1})
				visited[neighbor] = true
			}
		}
	}

	return -1
}

func getNeighbors(s string) []string {
	offsets := []int{-1, 1}
	neighbors := []string{}
	for _, offset := range offsets {
		for idx, r := range s {
			var sb strings.Builder
			i := int(r - '0')
			x := (i + offset + 10) % 10
			sb.WriteString(s[:idx])
			sb.WriteRune(rune(x + '0'))
			sb.WriteString(s[idx+1:])
			neighbors = append(neighbors, sb.String())
		}
	}

	return neighbors
}
