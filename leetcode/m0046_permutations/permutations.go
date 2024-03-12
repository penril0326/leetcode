package permutations

// Time: O(n^2*n!), n is length of nums
func permute(nums []int) [][]int {
	ans := [][]int{}
	backtracking([]int{}, &ans, nums)
	return ans
}

func backtracking(curr []int, ans *[][]int, nums []int) {
	if len(curr) == len(nums) {
		new := make([]int, len(curr))
		copy(new, curr)
		*ans = append(*ans, new)
		return
	}

	for _, n := range nums {
		if !IsContain(curr, n) {
			curr = append(curr, n)
			backtracking(curr, ans, nums)
			curr = curr[:len(curr)-1]
		}
	}
}

func IsContain(s []int, v int) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}

	return false
}
