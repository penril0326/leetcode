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
