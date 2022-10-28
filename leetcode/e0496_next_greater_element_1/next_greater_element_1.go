package nextgreaterelement1

import "practice/data_structure/stack"

// Time complexity: O(n^2)
// Space complexity: O(1)
func nextGreaterElementBruteForce(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	for idx1, n1 := range nums1 {
		result[idx1] = -1
		findNums := false
		for _, n2 := range nums2 {
			if n1 == n2 {
				findNums = true
			}

			if findNums {
				if n2 > n1 {
					result[idx1] = n2
					findNums = false
					break
				}
			}
		}
	}

	return result
}

// Monotonic stack
// Time complexity: O(n)
// Space complexity: O(n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	for idx, v := range nums1 {
		hash[v] = idx
	}

	decreasing := stack.NewStack()
	result := make([]int, len(nums1))
	for _, v := range nums2 {
		for !decreasing.IsEmpty() && (decreasing.Top().(int) < v) {
			if index, isExist := hash[decreasing.Top().(int)]; isExist {
				result[index] = v
			}

			decreasing.Pop()
		}

		decreasing.Push(v)
	}

	for !decreasing.IsEmpty() {
		if index, isExist := hash[decreasing.Pop().(int)]; isExist {
			result[index] = -1
		}
	}

	return result
}
