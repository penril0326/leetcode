package successfulpairsofspellsandpotiions

import "sort"

// Time: O((m+n)*logm), n is lenght of spells; m is lenght of potions
// Space: O(logm) or O(m), depend on sorting algorithm and language
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	ans := make([]int, len(spells))
	for idx, spell := range spells {
		l, r := 0, len(potions)-1
		for l <= r {
			mid := (l + r) >> 1
			target := float64(success) / float64(spell)
			if target > float64(potions[mid]) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		ans[idx] = len(potions) - l
	}

	return ans
}
