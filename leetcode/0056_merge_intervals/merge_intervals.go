package merge_intervals

import "sort"

type intervalSturct struct {
	start int
	end   int
}

type intervalList []intervalSturct

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	list := intervalList{}
	for _, interval := range intervals {
		temp := intervalSturct{
			start: interval[0],
			end:   interval[1],
		}

		list = append(list, temp)
	}

	sort.Sort(list)

	result := [][]int{}
	result = append(result, []int{list[0].start, list[0].end})
	for i := 1; i < list.Len(); i++ {
		if result[len(result)-1][1] < list[i].start {
			result = append(result, []int{
				list[i].start,
				list[i].end,
			})
		} else {
			result[len(result)-1][1] = max(result[len(result)-1][1], list[i].end)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (list intervalList) Len() int {
	return len(list)
}

func (list intervalList) Less(i, j int) bool {
	return list[i].start < list[j].start
}

func (list intervalList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
