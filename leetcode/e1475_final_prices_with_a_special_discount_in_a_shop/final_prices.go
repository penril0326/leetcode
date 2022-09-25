package finalpriceswithaspecialdiscountinashop

// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []int
}

func constructor() myStack {
	return myStack{
		stack: make([]int, 0),
	}
}

func (s myStack) top() int {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return 0
}

func (s *myStack) push(v int) {
	if (s != nil) && (s.stack != nil) {
		s.stack = append(s.stack, v)
	}
}

func (s *myStack) pop() int {
	if len(s.stack) > 0 {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return 0
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func finalPrices(prices []int) []int {
	stack := constructor()
	for i := 0; i < len(prices); i++ {
		for !stack.isEmpty() {
			topIndex := stack.top()
			if prices[i] > prices[topIndex] {
				break
			}

			prices[topIndex] -= prices[i]
			stack.pop()
		}

		stack.push(i)
	}

	return prices
}

// -----------------------------
// Time complexity: O(n^2)
// Space complexity: O(1)
func finalPricesBruteForce(prices []int) []int {
	if len(prices) <= 1 {
		return prices
	}

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}

	return prices
}
