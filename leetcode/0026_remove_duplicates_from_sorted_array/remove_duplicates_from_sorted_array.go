package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return 1
	}

	pointer := 0
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			pointer++
		} else {
			nums[i-pointer] = nums[i]
		}
	}

	nums = nums[:length-pointer]
	return len(nums)
}
