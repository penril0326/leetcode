package validatestacksequences

import datastructure "practice/data_structure"

// Time complexity: O(n)
// Space complexity: O(n)
func validateStackSequences(pushed []int, popped []int) bool {
	s := datastructure.NewStack()
	j := 0
	for _, i := range pushed {
		s.Push(i)
		for !s.IsEmpty() && (s.Top() == popped[j]) {
			s.Pop()
			j++
		}
	}

	return s.IsEmpty()
}
