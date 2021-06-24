package max_min

import "sort"

func maxMin(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	start, end := 0, int(k-1)
	min := arr[end] - arr[start]
	start++
	end++
	for start <= len(arr)-int(k) {
		newInterval := arr[end] - arr[start]
		if newInterval < min {
			min = newInterval
		}

		start++
		end++
	}

	return min
}
