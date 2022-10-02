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

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []int
}

func constructorStack() myStack {
	return myStack{
		stack: make([]int, 0),
	}
}

func (s myStack) top() int {
	if !s.isEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return -1
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *myStack) push(v int) {
	if s != nil {
		s.stack = append(s.stack, v)
	}
}

func (s *myStack) pop() int {
	if (s != nil) && !s.isEmpty() {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return -1
}

type myQueue struct {
	que  []int
	leng int
}

func constructorQue() myQueue {
	return myQueue{
		que: make([]int, 0),
	}
}

func (q myQueue) front() int {
	if !q.isEmpty() {
		return q.que[0]
	}

	return -1
}

func (q myQueue) isEmpty() bool {
	return q.leng == 0
}

func (q *myQueue) pushBack(v int) {
	if q != nil {
		q.que = append(q.que, v)
		q.leng++
	}
}

func (q *myQueue) removeFront() int {
	if (q != nil) && !q.isEmpty() {
		front := q.front()
		q.que = q.que[1:]
		q.leng--
		return front
	}

	return -1
}

func (q myQueue) length() int {
	return q.leng
}

func countStudentsQueStack(students []int, sandwiches []int) int {
	stus := constructorQue()
	for i := 0; i < len(students); i++ {
		stus.pushBack(students[i])
	}

	sands := constructorStack()
	for i := len(sandwiches) - 1; i >= 0; i-- {
		sands.push(sandwiches[i])
	}

	count := 0
	for !stus.isEmpty() {
		if stus.front() == sands.top() {
			stus.removeFront()
			sands.pop()
			count = 0
		} else {
			count++
			stus.pushBack(stus.front())
			stus.removeFront()
		}

		restStus := stus.length()
		if count >= restStus {
			return restStus
		}
	}

	return 0
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(1)
func countStudentsBest(students []int, sandwiches []int) int {
	preference := [2]int{}
	for _, stu := range students {
		preference[stu]++
	}

	for i := 0; i < len(sandwiches); i++ {

		sandwich := sandwiches[i]
		if preference[sandwich] == 0 {
			return len(sandwiches) - i
		}

		preference[sandwich]--
	}

	return 0
}
