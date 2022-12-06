package longestsubstringwithoutrepeatingchars

import "practice/utility/math"

// Time complexity: O(N)
// Space complexity: O(N)
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	ans := 0
	count := make(map[byte]int)
	for right < len(s) {
		count[s[right]]++
		for left < right && count[s[right]] > 1 {
			count[s[left]]--
			left++
		}

		ans = math.Max(ans, right-left+1)

		right++
	}

	return ans
}
