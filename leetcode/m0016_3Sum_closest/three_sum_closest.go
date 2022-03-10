package threesumclosest

import (
	"math"
	"sort"
)

// hint: same logic with 3Sum(leetcode 15)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	minDist, length, result := math.MaxInt64, len(nums), 0
	for i := 0; i < length-2; i++ {
		head, tail := i+1, length-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			dist := abs(target - sum)
			if dist < minDist {
				minDist = dist
				result = sum
			}

			if sum < target {
				head++
			} else {
				tail--
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
