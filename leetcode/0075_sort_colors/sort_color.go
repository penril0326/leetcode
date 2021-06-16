package sort_color

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			i--
			right--
		}
	}
}
