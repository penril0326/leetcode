package combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	current, result := make([]int, 0), make([][]int, 0)
	dfs(candidates, target, 0, current, &result)

	return result
}

func dfs(candidates []int, target, start int, current []int, result *[][]int) {
	if 0 >= target {
		if 0 == target {
			temp := make([]int, len(current))
			copy(temp, current)
			*result = append(*result, temp)
		}
	}

	for i := start; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}

		current = append(current, candidates[i])
		dfs(candidates, target-candidates[i], i, current, result)
		current = current[:len(current)-1]
	}
}
