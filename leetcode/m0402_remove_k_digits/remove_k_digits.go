package removekdigits

import "strings"

// Time: O(n)
// Space: O(n)
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	st := NewStack()
	for i := 0; i < len(num); i++ {
		for !st.IsEmpty() && k > 0 {
			// check top
			top := st.Top()
			if (top - '0') > (num[i] - '0') {
				st.Pop()
				k--
			} else {
				break
			}
		}

		st.Push(num[i])
	}

	for k > 0 {
		st.Pop()
		k--
	}

	return st.ToString()
}

type myStack struct {
	bs []byte
}

func NewStack() myStack {
	return myStack{
		bs: make([]byte, 0),
	}
}

func (s myStack) IsEmpty() bool {
	return len(s.bs) == 0
}

func (s myStack) Top() byte {
	return s.bs[len(s.bs)-1]
}

func (s *myStack) Push(b byte) {
	s.bs = append(s.bs, b)
}

func (s *myStack) Pop() byte {
	top := s.Top()
	s.bs = s.bs[:len(s.bs)-1]
	return top
}

func (s myStack) ToString() string {
	var sb strings.Builder
	nonZeroHeadIdx := -1
	for idx, b := range s.bs {
		if b != '0' {
			nonZeroHeadIdx = idx
			break
		}
	}

	if nonZeroHeadIdx == -1 {
		return "0"
	}

	for i := nonZeroHeadIdx; i < len(s.bs); i++ {
		sb.WriteByte(s.bs[i])
	}

	return sb.String()
}
