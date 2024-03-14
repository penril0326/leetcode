package wordsearch

// Time: O(m*n*3^L), where L is length of word
// Space: O(L) if we use set to store the node we seen; O(m*n) if we use 2D boolean array to store whole board state.
func exist(board [][]byte, word string) bool {
	seen := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		seen[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if word[0] == board[i][j] {
				seen[i][j] = true
				if backtracking(board, seen, i, j, 1, word) {
					return true
				}
				seen[i][j] = false
			}
		}
	}

	return false
}

func backtracking(board [][]byte, seen [][]bool, starti, startj, ansLen int, word string) bool {
	if ansLen == len(word) {
		return true
	}

	for _, node := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		x, y := starti+node[0], startj+node[1]
		if valid(x, y, len(board), len(board[starti])) && !seen[x][y] {
			if board[x][y] == word[ansLen] {
				seen[x][y] = true
				if backtracking(board, seen, x, y, ansLen+1, word) {
					return true
				}
				seen[x][y] = false
			}
		}
	}

	return false
}

func valid(i, j, maxi, maxj int) bool {
	return i >= 0 && j >= 0 && i < maxi && j < maxj
}
