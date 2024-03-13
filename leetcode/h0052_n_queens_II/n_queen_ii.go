package nqueensii

// Time: approximate to O(n!)
// Space: O(n)
func totalNQueens(n int) int {
	ans := 0
	cols, diagonals, antiDiagonals := make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{})

	backtracking(0, n, cols, diagonals, antiDiagonals, &ans)

	return ans
}

func backtracking(i, n int, cols, diagonals, antiDiagonals map[int]struct{}, ans *int) {
	if i == n {
		(*ans)++
		return
	}

	for j := 0; j < n; j++ {
		diagonal := i - j
		antiDiagonal := i + j
		_, colOccupied := cols[j]
		_, diagOccupied := diagonals[diagonal]
		_, antiDiagOccupied := antiDiagonals[antiDiagonal]
		if colOccupied || diagOccupied || antiDiagOccupied {
			continue
		}

		cols[j] = struct{}{}
		diagonals[diagonal] = struct{}{}
		antiDiagonals[antiDiagonal] = struct{}{}
		backtracking(i+1, n, cols, diagonals, antiDiagonals, ans)
		delete(cols, j)
		delete(diagonals, diagonal)
		delete(antiDiagonals, antiDiagonal)
	}
}
