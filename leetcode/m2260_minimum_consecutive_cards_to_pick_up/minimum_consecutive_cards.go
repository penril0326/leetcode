package minimumconsecutivecardstopickup

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Hash map + sliding window
// Time complexity: O(N)
// Space complexity: O(N)
func minimumCardPickup(cards []int) int {
	left := 0
	ans := len(cards)
	count := make(map[int]int)
	for right := 0; right < len(cards); right++ {
		count[cards[right]]++
		if len(count) < (right - left + 1) {
			ans = min(ans, right-left+1)

			for left < right {
				count[cards[left]]--
				if count[cards[left]] == 0 {
					delete(count, cards[left])
				}

				left++
				if len(count) == right-left+1 {
					break
				}

				ans = min(ans, right-left+1)
			}
		}
	}

	if len(count) == len(cards) {
		ans = -1
	}

	return ans
}

// Only hash map
// Time complexity: O(N)
// Space complexity: O(N)
func minimumCardPickup2(cards []int) int {
	ans := math.MaxInt
	idxMap := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		if idx, exist := idxMap[cards[i]]; exist {
			ans = min(ans, i-idx+1)
		}

		idxMap[cards[i]] = i
	}

	if ans == math.MaxInt {
		ans = -1
	}

	return ans
}
