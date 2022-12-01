package jewelsandstones

// Time complexity: O(J+S), which J is length of jewels, S is length of stones
// Space complexity: O(1)
func numJewelsInStones(jewels string, stones string) int {
	// 'A' ~ 'Z' -> 65 ~ 90
	// 'a' ~ 'z' -> 97 ~ 122
	stoneCnt := [58]int{}
	for _, r := range stones {
		stoneCnt[byte(r)-'A']++
	}

	total := 0
	for _, r := range jewels {
		total += stoneCnt[byte(r)-'A']
	}

	return total
}
