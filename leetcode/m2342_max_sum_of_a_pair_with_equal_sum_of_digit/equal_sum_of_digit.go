package maxsumofapairwithequalsumofdigit

import (
	"practice/algorithm/quick_sort"
	"practice/utility/digit"
	"practice/utility/math"
)

// hash + sort
// Time complexity: O(n*logn)
// Space complexity: O(n)
func maximumSum(nums []int) int {
	hash := make(map[int][]int)
	for _, n := range nums {
		digitSum := digit.SumOfDigit(n) // worst case is O(9)
		list := make([]int, 0)
		if v, exist := hash[digitSum]; exist {
			list = v
		}

		list = append(list, n)
		hash[digitSum] = list
	}

	ans := -1
	for _, list := range hash {
		if len(list) < 2 {
			continue
		}

		quick_sort.QuickSort(list, 0, len(list)-1) // O(n*logn)
		ans = math.Max(ans, list[len(list)-1]+list[len(list)-2])
	}

	return ans
}
