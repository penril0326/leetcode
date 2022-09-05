package nextgreaterelement1

import "fmt"

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

	fmt.Println(indexTable)

	result := make([]int,0)
	for i := 0; i < len(nums1); i++ {
		index := indexTable[nums1[i]]
		nexGreater := -1

		if index == len(nums2)-1 {
			result = append(result, nexGreater)
			continue
		}

		for j := index; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				nexGreater = nums2[]
			}
		}
	}

	return []int{}
}
