package movingaveragefromdatastream

// Time: O(1)
// Space: O(N)
type MovingAverage struct {
	q    []int
	sum  int
	size int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		q:    make([]int, 0),
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum += val
	this.q = append(this.q, val)
	if len(this.q) > this.size {
		front := this.q[0]
		this.q = this.q[1:]
		this.sum -= front
	}

	return float64(this.sum) / float64(len(this.q))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
