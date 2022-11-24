package quick_sort

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// QuickSort return increment oreder
func QuickSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		QuickSort(nums, left, p-1)
		QuickSort(nums, p+1, right)
	}
}
