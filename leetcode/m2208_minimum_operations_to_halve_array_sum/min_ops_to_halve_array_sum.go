package minimumoperationstohalvearraysum

import "container/heap"

// Time: O(nlogn)
// Space: O(n)
func halveArray(nums []int) int {
	sum := 0.0
	maxHeap := &maxFloatHeap{}
	for _, num := range nums {
		sum += float64(num)
		*maxHeap = append(*maxHeap, float64(num))
	}

	opsCount := 0
	heap.Init(maxHeap)

	final := sum / 2.0
	for sum > final {
		x := heap.Pop(maxHeap).(float64) / 2.0
		sum -= x

		heap.Push(maxHeap, x)
		opsCount++
	}

	return opsCount
}

type maxFloatHeap []float64

func (f maxFloatHeap) Len() int {
	return len(f)
}

func (f maxFloatHeap) Less(i, j int) bool {
	return f[i] >= f[j]
}

func (f maxFloatHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *maxFloatHeap) Push(x interface{}) {
	*f = append(*f, x.(float64))
}

func (f *maxFloatHeap) Pop() interface{} {
	n := f.Len()
	end := (*f)[n-1]
	*f = (*f)[:n-1]
	return end
}
