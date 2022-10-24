package dailytemperatures

// Time complexity: O(n^2)
// Space complexity: O(1)
func dailyTemperatures(temperatures []int) []int {
	currentDay := 0
	result := make([]int, len(temperatures))
	for currentDay < len(temperatures) {
		warmerDay := currentDay + 1
		count := 0
		for warmerDay < len(temperatures) {
			if temperatures[currentDay] < temperatures[warmerDay] {
				result[currentDay] = count + 1
				break
			}

			warmerDay++
			count++
			if warmerDay == len(temperatures) {
				result[currentDay] = 0
			}
		}

		currentDay++
	}

	return result
}

// -----------------------------
