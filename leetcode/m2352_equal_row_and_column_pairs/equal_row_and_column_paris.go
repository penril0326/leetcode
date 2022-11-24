package equalrowandcolumnpairs

import (
	"strconv"
	"strings"
)

// Time complexity: O(n^2)
// Space complexity: O(n^2)
func equalPairs(grid [][]int) int {
	rowHash := make(map[string]int)
	for _, row := range grid {
		key := convertArrayToString(row)
		rowHash[key]++
	}

	colHash := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		for j := i; j < len(grid); j++ {
			if i == j {
				continue
			}

			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}

		key := convertArrayToString(grid[i])
		colHash[key]++
	}

	sum := 0
	for key, v := range rowHash {
		if colHash[key] > 0 {
			sum += colHash[key] * v
		}
	}

	return sum
}

// string builder can concatenate string with O(1)
func convertArrayToString(arr []int) string {
	var sb strings.Builder
	for _, v := range arr {
		s := strconv.Itoa(v)
		sb.WriteString(s)
		sb.WriteString(",")
	}

	return sb.String()
}
