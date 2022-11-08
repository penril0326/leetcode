package runningsumof1darray

// Time complexity: O(N)
// Space complexity: O(N) if we consider the array "result".
func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}

	return result
}
