package squaresofasortedarray

// Time complexity: O(N)
// Space complexity: O(N), if take output into account
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		rightSqur, leftSqur := squr(nums[right]), squr(nums[left])
		if leftSqur < rightSqur {
			result[i] = rightSqur
			right--
		} else {
			result[i] = leftSqur
			left++
		}
	}

	return result
}

func squr(x int) int {
	return x * x
}
