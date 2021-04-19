package median

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if 0 == m {
		return getMedian(nums2)
	} else if 0 == n {
		return getMedian(nums1)
	}

	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	halfLen := (m + n + 1) / 2
	min1, max1 := 0, m
	for min1 <= max1 {
		i := (min1 + max1) / 2
		j := halfLen - i

		if (i < max1) && (nums2[j-1] > nums1[i]) {
			// i太小(nums1提供太少元素)
			min1 = i + 1
		} else if (i > min1) && (nums1[i-1] > nums2[j]) {
			// i太大(nums1提供太多元素)
			max1 = i - 1
		} else {
			// i，j都正確了
			maxLeft := 0
			if 0 == i {
				// i=0表示nums1一個元素都沒提供
				maxLeft = nums2[j-1]
			} else if 0 == j {
				// j=0表示nums2一個元素都沒有提供
				maxLeft = nums1[i-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			if 1 == (m+n)%2 {
				return float64(maxLeft)
			}

			minRight := 0
			if m == i {
				minRight = nums2[j]
			} else if n == j {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return (float64(maxLeft) + float64(minRight)) / 2.0
		}
	}

	return 0.0
}

func getMedian(list []int) float64 {
	len := len(list)
	index := len / 2
	median := float64(list[index])

	if 0 == len%2 {
		median = float64((list[index] + list[index-1])) / 2.0
	}

	return median
}
