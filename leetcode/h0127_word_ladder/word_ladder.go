package h0127wordladder

// Time: O(M*N^2), M is word length, N is wordList size
// Space: O(M^2*N^2) (?)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := map[string][]string{}
	graph[beginWord] = getNeighbors(beginWord, wordList)

	if len(graph[beginWord]) == 0 {
		return 0
	}

	for i := 0; i < len(wordList); i++ {
		if beginWord == wordList[i] {
			continue
		}

		graph[wordList[i]] = getNeighbors(wordList[i], wordList)
	}

	type state struct {
		s    string
		step int
	}

	queue := make([]state, 0)
	queue = append(queue, state{beginWord, 1})
	visited := map[string]bool{}
	visited[beginWord] = true
	for len(queue) > 0 {
		curr, step := queue[0].s, queue[0].step
		queue = queue[1:]
		if curr == endWord {
			return step
		}

		for _, neighbor := range graph[curr] {
			if !visited[neighbor] {
				queue = append(queue, state{neighbor, step + 1})
				visited[neighbor] = true
			}
		}
	}

	return 0
}

func getNeighbors(s string, wordList []string) []string {
	neighbors := []string{}
	for i := 0; i < len(wordList); i++ {
		if s == wordList[i] {
			continue
		}

		if isNeighbor(s, wordList[i]) {
			neighbors = append(neighbors, wordList[i])
		}
	}

	return neighbors
}

func isNeighbor(s1, s2 string) bool {
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
		}
	}

	return diffCount == 1
}
