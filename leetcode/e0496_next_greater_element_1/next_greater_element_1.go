package nextgreaterelement1

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexTable := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				indexTable[nums1[i]] = j
				break
			}
		}
	}

	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		index := indexTable[nums1[i]]
		nexGreater := -1
		for j := index + 1; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				nexGreater = nums2[j]
				break
			}
		}

		result = append(result, nexGreater)
	}

	return result
}
