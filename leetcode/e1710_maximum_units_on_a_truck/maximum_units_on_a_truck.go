package maximumunitsonatruck

import "container/heap"

// Time: O(n*logn)
// Space: O(n)
func maximumUnits(boxTypes [][]int, truckSize int) int {
	q := &pq{}
	*q = append(*q, boxTypes...)
	heap.Init(q)

	total := 0
	for truckSize > 0 && q.Len() > 0 {
		maxBox := heap.Pop(q).([]int)
		if maxBox[0] > truckSize {
			total += truckSize * maxBox[1]
			break
		}

		total += maxBox[0] * maxBox[1]
		truckSize -= maxBox[0]
	}

	return total
}

type pq [][]int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i][1] > q[j][1]
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.([]int))
}

func (q *pq) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}

	lastIdx := q.Len() - 1
	root := (*q)[lastIdx]
	*q = (*q)[:lastIdx]

	return root
}
