package removeelment

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}

	nums = nums[:slow]
	return len(nums)
}
