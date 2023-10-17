package kokoeatingbananas

import (
	"math"
	myMath "practice/utility/math"
)

// Time: O(n*logk), n is size of piles, k is the max amount of bananas
// Space: O(1)
func minEatingSpeed(piles []int, h int) int {
	minSpeed, maxSpeed := 1, math.MinInt
	for _, bananas := range piles {
		maxSpeed = myMath.Max(maxSpeed, bananas)
	}

	if len(piles) == h {
		return maxSpeed
	}

	for minSpeed <= maxSpeed {
		mid := minSpeed + (maxSpeed-minSpeed)>>1
		totalHours := 0
		for _, bananas := range piles {
			totalHours += ceil(bananas, mid)
		}

		if totalHours <= h {
			maxSpeed = mid - 1
		} else {
			minSpeed = mid + 1
		}
	}

	return minSpeed
}

func ceil(x, y int) int {
	if x < y {
		return 1
	}

	sum := x / y
	if x%y > 0 {
		sum++
	}

	return sum
}
