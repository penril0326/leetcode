package numberofstudentsunabletoeatlunch

// Time complexity: O(n^2)
// Space complexity: O(n)
func countStudents(students []int, sandwiches []int) int {
	for len(sandwiches) > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			find := false
			for i := 0; i < len(students); i++ {
				if students[i] == sandwiches[0] {
					find = true
					break
				}
			}

			if !find {
				return len(students)
			}

			top := students[0]
			students = append(students[1:], top)
		}
	}

	return 0
}
