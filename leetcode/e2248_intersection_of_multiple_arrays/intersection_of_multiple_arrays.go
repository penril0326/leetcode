package intersectionofmultiplearrays

func intersection(nums [][]int) []int {
	max := 0
	hash := make(map[int]int)
	for _, arr := range nums {
		for _, n := range arr {
			hash[n]++

			if n > max {
				max = n
			}
		}
	}

	ans := []int{}
	for i := 1; i <= max; i++ {
		if hash[i] == len(nums) {
			ans = append(ans, i)
		}
	}

	return ans
}
