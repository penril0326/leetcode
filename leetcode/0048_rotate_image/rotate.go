package rotate

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			matrix[i][j], matrix[length-1-j][length-1-i] = matrix[length-1-j][length-1-i], matrix[i][j]
		}
	}

	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
}
