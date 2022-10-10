package minstack

type impvTwoStack struct {
	main []int
	min  [][2]int
}

func ConstructorImpvTwoStack() impvTwoStack {
	return impvTwoStack{
		main: make([]int, 0),
		min:  make([][2]int, 0),
	}
}

func (s *impvTwoStack) Push(val int) {
	if s != nil {
		s.main = append(s.main, val)
		if (len(s.min) == 0) || (val < s.min[len(s.min)-1][0]) {
			s.min = append(s.min, [2]int{val, 1})
		} else if val == s.min[len(s.min)-1][0] {
			s.min[len(s.min)-1][1]++
		}
	}
}

func (s *impvTwoStack) Pop() {
	if (s != nil) && (len(s.main) > 0) && (len(s.min) > 0) {
		if s.main[len(s.main)-1] == s.min[len(s.min)-1][0] {
			if s.min[len(s.min)-1][1] > 1 {
				s.min[len(s.min)-1][1]--
			} else {
				s.min = s.min[:len(s.min)-1]
			}
		}

		s.main = s.main[:len(s.main)-1]
	}
}

func (s *impvTwoStack) Top() int {
	if (s != nil) && (len(s.main) > 0) {
		return s.main[len(s.main)-1]
	}

	return -1
}

func (s *impvTwoStack) GetMin() int {
	if (s != nil) && (len(s.min) > 0) {
		return s.min[len(s.min)-1][0]
	}

	return -1
}
