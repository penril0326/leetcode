package search2dmatrix

// Time: O(logM*N)
// Space: O(1)
func searchMatrixOptimize(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)*len(matrix[0])-1
	for l <= r {
		mid := (l + r) >> 1
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

// Time: O(logM*logN)
// Space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	top, down := 0, len(matrix)-1
	for top <= down {
		midRow := (top + down) >> 1
		rowL, rowR := 0, len(matrix[midRow])-1
		if (target >= matrix[midRow][rowL]) && (target <= matrix[midRow][rowR]) {
			// do binary search current row
			for rowL <= rowR {
				mid := (rowL + rowR) >> 1
				if target == matrix[midRow][mid] {
					return true
				} else if target > matrix[midRow][mid] {
					rowL = mid + 1
				} else {
					rowR = mid - 1
				}
			}

			return false
		} else if target < matrix[midRow][rowL] {
			down = midRow - 1
		} else {
			top = midRow + 1
		}
	}

	return false
}
