package subsets

// Time: O(2^n*n)
// Space: O(n)
func subsets(nums []int) [][]int {
	ans := [][]int{}
	curr := []int{}
	backtracking(curr, nums, 0, &ans)
	return ans
}

func backtracking(curr []int, nums []int, start int, ans *[][]int) {
	new := make([]int, len(curr))
	copy(new, curr)
	*ans = append(*ans, new)

	if start == len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtracking(curr, nums, i+1, ans)
		curr = curr[:len(curr)-1]
	}
}
