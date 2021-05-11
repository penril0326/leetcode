package container_water

func maxArea(height []int) int {
	head, tail := 0, len(height)-1
	maxArea := 0
	for head < tail {
		h := min(height[head], height[tail])
		w := tail - head
		area := h * w

		maxArea = max(maxArea, area)

		for (head < tail) && (h >= height[head]) {
			head++
		}

		for (head < tail) && (h >= height[tail]) {
			tail--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
