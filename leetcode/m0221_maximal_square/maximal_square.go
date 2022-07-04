package maximal_square

// brute force
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}

	max := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				l := 1
				isFind := true
				for (i+l < rows) && (j+l < cols) && (isFind == true) {
					for k := j; k <= j+l; k++ {
						if matrix[i+l][k] == '0' {
							isFind = false
							break
						}
					}

					for k := i; k <= i+l; k++ {
						if matrix[k][j+l] == '0' {
							isFind = false
							break
						}
					}

					if isFind == true {
						l++
					}
				}
				if max < l {
					max = l
				}
			}
		}
	}

	return max * max
}

// DP
func maximalSquareDP(matrix [][]byte) int {
	row, column := len(matrix), 0
	if row > 0 {
		column = len(matrix[0])
	}

	dp := make([][]int, 0)
	for i := 0; i < row; i++ {
		temp := make([]int, column)
		dp = append(dp, temp)
	}

	max := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == '1' {
				if (i == 0) || (j == 0) {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}

				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}

		return b
	} else {
		if a > c {
			return c
		}

		return a
	}
}
