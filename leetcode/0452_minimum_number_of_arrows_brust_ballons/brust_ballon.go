package brust_ballon

import "sort"

type myPoint struct {
	start int
	end   int
}

type myPoints []myPoint

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	pointList := myPoints{}
	for _, point := range points {
		p := myPoint{
			start: point[0],
			end:   point[1],
		}

		pointList = append(pointList, p)
	}

	sort.Sort(pointList)

	count := 1
	end := pointList[0].end
	for i := 1; i < len(pointList); i++ {
		if pointList[i].start <= end {
			end = min(end, pointList[i].end)
		} else {
			count++
			end = pointList[i].end
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (mp myPoints) Len() int {
	return len(mp)
}

func (mp myPoints) Less(i, j int) bool {
	if mp[i].start < mp[j].start {
		return true
	} else if mp[i].start == mp[j].start {
		if mp[i].end < mp[j].end {
			return true
		}
	}

	return false
}

func (mp myPoints) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}
