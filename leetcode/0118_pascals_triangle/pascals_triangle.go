package pascalstriangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result = append(result, row)
	}

	return result
}
