package onlinestockspan

import "practice/data_structure/stack"

// Array version
// Time complexity: O(m), which m < n and n is total counts of next() call
// Space complexity: O(n)
type StockSpannerArray struct {
	data []int
}

func ConstructorArray() StockSpannerArray {
	return StockSpannerArray{
		data: make([]int, 0),
	}
}

func (this *StockSpannerArray) Next(price int) int {
	count := 1
	for i := len(this.data) - 1; i >= 0; i-- {
		if this.data[i] <= price {
			count++
		} else {
			break
		}
	}

	this.data = append(this.data, price)

	return count
}

// Monotonic stack version
// Time complexity: O(1) (Amortised)
// Space complexity: O(n)

type StockSpanner struct {
	data stack.MyStack
}

type maximumPrice struct {
	price       int
	consecutive int
}

func Constructor() StockSpanner {
	return StockSpanner{
		data: stack.NewStack(),
	}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for !this.data.IsEmpty() && (this.data.Top().(maximumPrice).price <= price) {
		consDay := this.data.Pop().(maximumPrice).consecutive
		count += consDay
	}

	this.data.Push(maximumPrice{
		price:       price,
		consecutive: count,
	})

	return count
}
