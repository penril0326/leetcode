package combinations

// Time: O()
func combine(n int, k int) [][]int {
	ans := [][]int{}
	backtracking([]int{}, 1, n, k, &ans)
	return ans
}

func backtracking(curr []int, start, end, k int, ans *[][]int) {
	if len(curr) == k {
		new := make([]int, k)
		copy(new, curr)
		*ans = append(*ans, new)
		return
	}

	for i := start; i <= end; i++ {
		curr = append(curr, i)
		backtracking(curr, i+1, end, k, ans)
		curr = curr[:len(curr)-1]
	}
}
