package longestsubsequencewithlimitedsum

import "sort"

// Time: O((m+n)*logn), which m is length of queries, n is length of prefixSum quilavent to length of nums
// Space: O(n)
func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(ans); i++ {
		ans[i] = findSubsequence(prefixSum, queries[i])
	}

	return ans
}

func findSubsequence(prefixSum []int, query int) int {
	l, r := 0, len(prefixSum)-1
	for l <= r {
		mid := (l + r) >> 1
		if query == prefixSum[mid] {
			return mid
		} else if query < prefixSum[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r
}
