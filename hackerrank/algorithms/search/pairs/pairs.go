package pairs

import "sort"

func pairs(k int32, arr []int32) int32 {
	// 先排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	first, second := 0, 1
	count := 0
	for second < len(arr) {
		interval := arr[second] - arr[first]

		// 區間 < k表示second指針要往後移
		if interval < k {
			second++
		} else {
			if interval == k {
				count++
			}

			first++
			second = first + 1
		}
	}

	return int32(count)
}
