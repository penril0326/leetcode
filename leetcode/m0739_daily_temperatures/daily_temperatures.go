package dailytemperatures

import datastructure "practice/data_structure"

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
// Time complexity: O(n)
// Space complexity: O(n)
func dailyTemperaturesMonotonicDecreasingStack(temperatures []int) []int {
	result := make([]int, len(temperatures))
	s := datastructure.NewStack()
	for idx, temperature := range temperatures {
		for !s.IsEmpty() && (temperature > temperatures[s.Top().(int)]) {
			top := s.Pop().(int)
			result[top] = idx - top
		}

		s.Push(idx)
	}

	return result
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(1)
func dailyTemperaturesOptimizeArray(temperatures []int) []int {
	result := make([]int, len(temperatures))
	hottest := 0
	for curDay := len(temperatures) - 1; curDay >= 0; curDay-- {
		if temperatures[curDay] >= hottest {
			hottest = temperatures[curDay]
			continue
		}

		day := 1
		for temperatures[curDay] >= temperatures[curDay+day] {
			day += result[curDay+day]
		}
		result[curDay] = day
	}

	return result
}
