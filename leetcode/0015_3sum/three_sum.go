package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	length := len(nums)
	if (0 == length) || (0 < nums[0]) || (0 > nums[length-1]) {
		return [][]int{}
	}

	result := [][]int{}
	for i := 0; i < length-2; i++ {
		if 0 < nums[i] {
			break
		}

		if (0 < i) && (nums[i] == nums[i-1]) {
			continue
		}

		target, head, tail := nums[i], i+1, length-1
		for head < tail {
			sum := target + nums[head] + nums[tail]
			if 0 == sum {
				result = append(result, []int{target, nums[head], nums[tail]})
				for (head < tail) && (nums[head] == nums[head+1]) {
					head++
				}

				for (head < tail) && (nums[tail] == nums[tail-1]) {
					tail--
				}

				head++
				tail--
			} else if 0 > sum {
				head++
			} else {
				tail--
			}
		}
	}

	return result
}
